package bdd

import (
	"reflect"
	"testing"
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
