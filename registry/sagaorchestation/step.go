package sagaorchestation

type StepI interface {
	GetName() string
	ProcessFunc() error
	CompensatingFunc() error
}

type Step struct {
	Name string
}
