package sagaorchestration

import (
	"fmt"
	uuid2 "github.com/google/uuid"
	"github.com/ntc-goer/microservice-examples/registry/sagaorchestation/ent"
)

type ProcessStatus string

const (
	SUCCESS       ProcessStatus = "SUCCESS"
	REVERT_FAILED ProcessStatus = "REVERT_FAILED"
	REVERTED      ProcessStatus = "REVERTED"
	SKIPPED       ProcessStatus = "SKIPPED"
	PENDING       ProcessStatus = "PENDING"
	WAITING       ProcessStatus = "WAITING"
)

type Log struct {
	StepName string
	Status   ProcessStatus
}

type WorkflowI interface {
	Start() error
	Revert()
	GetLog() []Log
}

type Workflow[T any] struct {
	ID               uuid2.UUID
	Store            T
	Name             string
	CurrentStep      int
	ResultStatus     ProcessStatus
	Log              []Log
	Steps            []Step[T]
	TrackingDB       *DB
	TrackingDBClient *ent.Client
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
}

func NewWorkflow[T any](name string, cfg WorkflowConfig[T]) (*Workflow[T], error) {
	id := uuid2.New()
	wl := &Workflow[T]{
		ID:         id,
		Store:      cfg.Store,
		Name:       name,
		TrackingDB: cfg.TrackingDB,
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
		if err != nil {
			return nil, err
		}
	}

	return wl, nil
}

func (wf *Workflow[T]) initProcess() {
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

func (wf *Workflow[T]) Start() error {
	// Init Process Log
	wf.initProcess()
	for index, step := range wf.Steps {
		wf.CurrentStep = index
		wf.Log[wf.CurrentStep].Status = PENDING
		if err := step.ProcessF(wf.Store); err != nil {
			wf.Revert()
			return err
		}
		wf.Log[wf.CurrentStep].Status = SUCCESS
	}
	return nil
}

func (wf *Workflow[T]) Revert() {
	for i := wf.CurrentStep; i >= 0; i-- {
		if wf.Steps[i].CompensatingF(wf.Store) == nil {
			wf.Log[wf.CurrentStep].Status = SKIPPED
			continue
		}
		if err := wf.Steps[i].CompensatingF(wf.Store); err != nil {
			wf.Log[wf.CurrentStep].Status = REVERT_FAILED
		}
		wf.Log[wf.CurrentStep].Status = REVERTED
	}
}

func (wf *Workflow[T]) GetLog() []Log {
	return wf.Log
}
