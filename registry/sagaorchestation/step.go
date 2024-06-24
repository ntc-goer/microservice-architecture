package sagaorchestration

type StepI interface {
	GetName() string
}

type Step[T any] struct {
	Name          string
	ProcessF      func(store T) error
	CompensatingF func(store T) error
}

func (step *Step[T]) GetName() string {
	return step.Name
}
