package sagaorchestration

type StepI interface {
	GetName() string
}

type Step struct {
	Name          string
	ProcessF      func() error
	CompensatingF func() error
}

func (step *Step) GetName() string {
	return step.Name
}
