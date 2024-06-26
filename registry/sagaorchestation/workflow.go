package sagaorchestration

import (
	"context"
	"fmt"
	uuid2 "github.com/google/uuid"
	_ "github.com/lib/pq"
	"github.com/ntc-goer/microservice-examples/registry/sagaorchestation/ent"
	"github.com/ntc-goer/microservice-examples/registry/sagaorchestation/ent/sagalogs"
	"github.com/prometheus/common/log"
)

type ProcessStatus string

const (
	SUCCESS       ProcessStatus = "SUCCESS"
	REVERT_FAILED ProcessStatus = "REVERT_FAILED"
	REVERTED      ProcessStatus = "REVERTED"
	SKIPPED       ProcessStatus = "SKIPPED"
	PENDING       ProcessStatus = "PENDING"
	WAITING       ProcessStatus = "WAITING"
	ERROR         ProcessStatus = "ERROR"
)

type Log struct {
	StepName string
	Status   ProcessStatus
	Message  string
}

type WorkflowI interface {
	Start() error
	Revert()
	GetLog() []Log
}

type Workflow[T any] struct {
	ID                uuid2.UUID
	Store             T
	Name              string
	CurrentStep       int
	ResultStatus      ProcessStatus
	Log               []Log
	Steps             []Step[T]
	TrackingDB        *DB
	TrackingDBClient  *ent.Client
	DBTrackingEnabled bool
	RequestID         string
}

type DB struct {
	DriverName string
	Address    string
	Port       string
	DBName     string
	UserName   string
	Password   string
}

type WorkflowConfig[T any] struct {
	Store      T
	TrackingDB *DB
	RequestID  string
}

func NewWorkflow[T any](name string, cfg WorkflowConfig[T]) (*Workflow[T], error) {
	id := uuid2.New()
	wl := &Workflow[T]{
		ID:         id,
		Store:      cfg.Store,
		Name:       name,
		TrackingDB: cfg.TrackingDB,
		RequestID:  cfg.RequestID,
	}
	// Migrate DB
	if cfg.TrackingDB != nil {
		client, err := ent.Open(cfg.TrackingDB.DriverName,
			fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
				cfg.TrackingDB.Address,
				cfg.TrackingDB.Port,
				cfg.TrackingDB.UserName,
				cfg.TrackingDB.Password,
				cfg.TrackingDB.DBName,
			))
		if err != nil {
			return nil, err
		}
		wl.TrackingDBClient = client
		err = wl.TrackingDBClient.Schema.Create(context.Background())
		if err != nil {
			log.Errorf("Migration to DB Fail %s", err)
		}
		wl.TrackingDBClient = client
		wl.DBTrackingEnabled = true
	}
	return wl, nil
}

func (wf *Workflow[T]) initProcess() {
	if wf.RequestID == "" {
		wf.RequestID = uuid2.New().String()
	}
	wf.Log = make([]Log, len(wf.Steps))
	for i := range wf.Log {
		wf.Log[i] = Log{
			StepName: wf.Steps[i].GetName(),
			Status:   WAITING,
		}
	}
}

func (wf *Workflow[T]) Register(step Step[T]) *Workflow[T] {
	wf.Steps = append(wf.Steps, step)
	return wf
}

func (wf *Workflow[T]) RegisterSteps(steps []Step[T]) *Workflow[T] {
	wf.Steps = steps
	return wf
}

func (wf *Workflow[T]) UpdateTrackingStep() {
	if !wf.DBTrackingEnabled {
		return
	}
	ctx := context.Background()
	curStep := wf.Steps[wf.CurrentStep]
	stepName := curStep.GetName()
	stepStatus := wf.Log[wf.CurrentStep].Status
	stepMsg := wf.Log[wf.CurrentStep].Message

	stepLog, err := wf.TrackingDBClient.SagaLogs.Query().Where(
		sagalogs.WorkflowID(wf.ID.String()),
		sagalogs.StepName(stepName),
		sagalogs.RequestID(wf.RequestID),
	).Only(ctx)
	if err != nil && !ent.IsNotFound(err) {
		if ent.IsNotSingular(err) {
			log.Errorf("More than sigle step log found %s", err)
			return
		}
		log.Errorf("Error when checking step log %s", err)
		return
	}
	if ent.IsNotFound(err) {
		// Create Step Log
		_, err := wf.TrackingDBClient.SagaLogs.Create().
			SetWorkflowID(wf.ID.String()).
			SetStepName(stepName).
			SetRequestID(wf.RequestID).
			SetStepOrder(wf.CurrentStep + 1).
			SetStatus(string(stepStatus)).
			SetWorkflowName(wf.Name).
			SetMessage(stepMsg).
			Save(ctx)
		if err != nil {
			log.Errorf("Create tracking log fail %s", err)
			return
		}
	} else {
		// Update Step Log
		_, err := stepLog.Update().SetStatus(string(stepStatus)).SetMessage(stepMsg).Save(ctx)
		if err != nil {
			log.Errorf("Update tracking log fail %s", err)
			return
		}
	}

}

func (wf *Workflow[T]) CreateTrackingStep() {
	if !wf.DBTrackingEnabled {
		return
	}
	ctx := context.Background()
	curStep := wf.Steps[wf.CurrentStep]
	stepName := curStep.GetName()
	stepStatus := wf.Log[wf.CurrentStep].Status
	stepMsg := wf.Log[wf.CurrentStep].Message

	// Create Step Log
	_, err := wf.TrackingDBClient.SagaLogs.Create().
		SetWorkflowID(wf.ID.String()).
		SetStepName(stepName).
		SetRequestID(wf.RequestID).
		SetStepOrder(wf.CurrentStep + 1).
		SetStatus(string(stepStatus)).
		SetWorkflowName(wf.Name).
		SetMessage(stepMsg).
		Save(ctx)
	if err != nil {
		log.Errorf("Create tracking log fail %s", err)
		return
	}
}

func (wf *Workflow[T]) Start() error {
	// Init Process Log
	wf.initProcess()
	for index, step := range wf.Steps {
		wf.CurrentStep = index
		wf.Log[wf.CurrentStep].Status = PENDING
		wf.Log[wf.CurrentStep].Message = "Ready"
		wf.CreateTrackingStep()
		if err := step.ProcessF(wf.Store); err != nil {
			wf.Log[wf.CurrentStep].Status = ERROR
			wf.Log[wf.CurrentStep].Message = fmt.Sprintf("Error: %s", err.Error())
			wf.CreateTrackingStep()
			wf.Revert()
			return err
		}
		wf.Log[wf.CurrentStep].Status = SUCCESS
		wf.Log[wf.CurrentStep].Message = "Successfully"
		wf.CreateTrackingStep()
	}
	return nil
}

func (wf *Workflow[T]) Revert() {
	for i := wf.CurrentStep; i >= 0; i-- {
		wf.CurrentStep = i
		if wf.Steps[i].CompensatingF == nil {
			wf.Log[wf.CurrentStep].Status = SKIPPED
			wf.Log[wf.CurrentStep].Message = "REVERTED: No compensating function provided"
			wf.CreateTrackingStep()
			continue
		}
		if err := wf.Steps[i].CompensatingF(wf.Store); err != nil {
			wf.Log[wf.CurrentStep].Status = ERROR
			wf.Log[wf.CurrentStep].Message = fmt.Sprintf("ERROR: Process function error : %s", err)
			wf.CreateTrackingStep()
			continue
		}

		wf.Log[wf.CurrentStep].Status = REVERTED
		wf.Log[wf.CurrentStep].Message = "REVERTED: revert done"
		wf.UpdateTrackingStep()
	}
}

func (wf *Workflow[T]) GetLog() []Log {
	return wf.Log
}
