package sagaorchestration

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
	Store        T
	Name         string
	CurrentStep  int
	ResultStatus ProcessStatus
	Log          []Log
	Steps        []Step[T]
}

func NewWorkflow[T any](name string, initStore T) *Workflow[T] {
	return &Workflow[T]{
		Store: initStore,
		Name:  name,
	}
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
