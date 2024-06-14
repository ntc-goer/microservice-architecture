package sagaorchestation

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
	Register(step StepI)
	Start()
	Revert()
	GetLog() []Log
}

type Workflow struct {
	Name         string
	CurrentStep  int
	ResultStatus ProcessStatus
	Log          []Log
	Steps        []StepI
}

func NewWorkflow(name string) WorkflowI {
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

func (wf *Workflow) Register(step StepI) {
	wf.Steps = append(wf.Steps, step)
}

func (wf *Workflow) Start() {
	// Init Process Log
	wf.initProcess()
	for index, step := range wf.Steps {
		wf.CurrentStep = index
		wf.Log[wf.CurrentStep].Status = PENDING
		if err := step.ProcessFunc(); err != nil {
			wf.Revert()
			return
		}
		wf.Log[wf.CurrentStep].Status = SUCCESS
	}
}

func (wf *Workflow) Revert() {
	for i := wf.CurrentStep; i >= 0; i-- {
		if wf.Steps[i].CompensatingFunc == nil {
			wf.Log[wf.CurrentStep].Status = SKIPPED
			continue
		}
		if err := wf.Steps[i].CompensatingFunc(); err != nil {
			wf.Log[wf.CurrentStep].Status = REVERT_FAILED
		}
		wf.Log[wf.CurrentStep].Status = REVERTED
	}
}

func (wf *Workflow) GetLog() []Log {
	return wf.Log
}
