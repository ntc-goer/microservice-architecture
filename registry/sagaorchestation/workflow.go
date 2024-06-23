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
	RegisterSteps(steps []Step) *Workflow
	Register(step Step) *Workflow
	Start() error
	Revert()
	GetLog() []Log
}

type Workflow struct {
	Name         string
	CurrentStep  int
	ResultStatus ProcessStatus
	Log          []Log
	Steps        []Step
}

func NewWorkflow(name string) *Workflow {
	return &Workflow{
		Name: name,
	}
}

func (wf *Workflow) initProcess() {
	wf.Log = make([]Log, len(wf.Steps))
	for i := range wf.Log {
		wf.Log[i] = Log{
			StepName: wf.Steps[i].GetName(),
			Status:   WAITING,
		}
	}
}

func (wf *Workflow) Register(step Step) *Workflow {
	wf.Steps = append(wf.Steps, step)
	return wf
}

func (wf *Workflow) RegisterSteps(steps []Step) *Workflow {
	wf.Steps = steps
	return wf
}

func (wf *Workflow) Start() error {
	// Init Process Log
	wf.initProcess()
	for index, step := range wf.Steps {
		wf.CurrentStep = index
		wf.Log[wf.CurrentStep].Status = PENDING
		if err := step.ProcessF(); err != nil {
			wf.Revert()
			return err
		}
		wf.Log[wf.CurrentStep].Status = SUCCESS
	}
	return nil
}

func (wf *Workflow) Revert() {
	for i := wf.CurrentStep; i >= 0; i-- {
		if wf.Steps[i].CompensatingF() == nil {
			wf.Log[wf.CurrentStep].Status = SKIPPED
			continue
		}
		if err := wf.Steps[i].CompensatingF(); err != nil {
			wf.Log[wf.CurrentStep].Status = REVERT_FAILED
		}
		wf.Log[wf.CurrentStep].Status = REVERTED
	}
}

func (wf *Workflow) GetLog() []Log {
	return wf.Log
}
