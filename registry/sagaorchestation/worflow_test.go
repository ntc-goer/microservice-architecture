package sagaorchestration

//
//import (
//	"fmt"
//	"github.com/stretchr/testify/assert"
//	"testing"
//)
//
//type TestCase struct {
//	Name           string
//	Steps          []Step
//	ExpectedResult error
//}
//
//func Test_Workflow(t *testing.T) {
//	wl := NewWorkflow("Test_Workflow")
//	testCases := []TestCase{
//		{
//			Name: "Success Case",
//			Steps: []Step{
//				{
//					Name: "Action 1",
//					ProcessF: func() error {
//						fmt.Println("ProcessF 1")
//						return nil
//					},
//					CompensatingF: func() error {
//						fmt.Println("ProcessF 1")
//						return nil
//					},
//				},
//				{
//					Name: "Action 2",
//					ProcessF: func() error {
//						fmt.Println("ProcessF 2")
//						return nil
//					},
//					CompensatingF: func() error {
//						fmt.Println("ProcessF 2")
//						return nil
//					},
//				},
//			},
//		},
//	}
//
//	for _, testCase := range testCases {
//		wl.RegisterSteps(testCase.Steps)
//		wl.Start()
//		assert.Equal(t, len(wl.GetLog()), len(testCase.Steps))
//	}
//
//}
