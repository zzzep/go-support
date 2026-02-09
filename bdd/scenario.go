package bdd

import (
	"fmt"
	"reflect"
	"sync"
	"testing"
)

var (
	stepCounter int
	counterMutex sync.Mutex
)

type Scenario struct {
	T     *testing.T
	Steps []Step
	Mocks map[any]Mock
}

func NewScenario() *Scenario {
	return &Scenario{
		Mocks: make(map[any]Mock),
	}
}

func (sc *Scenario) AddStep(step Step) *Scenario {
	sc.Steps = append(sc.Steps, step)
	return sc
}

func (sc *Scenario) AddMock(key string, mock Mock) *Scenario {
	sc.Mocks[key] = mock
	return sc
}

func (sc *Scenario) Run() {
	name := ""
	for _, step := range sc.Steps {
		name += step.GetFunctionName(step) + " "
	}
	sc.T.Run(name, func(t *testing.T) {
		for _, step := range sc.Steps {
			// Get and increment step counter
			counterMutex.Lock()
			stepCounter++
			currentStep := stepCounter
			counterMutex.Unlock()
			
			// Get step name
			stepName := step.GetFunctionName(step)
			
			// Log with counter format: #N# - STEP_NAME
			fmt.Printf("#%d# - %s\n", currentStep, stepName)
			
			// Execute the step
			step(*sc, t)
		}
	})
}

func GenerateMock[T any](mock Mock) (i T, m Mock) {
	if !reflect.TypeOf(mock).Implements(reflect.TypeOf((*T)(nil)).Elem()) {
		panic("Mock does not implement interface")
	}
	return i, mock
}
