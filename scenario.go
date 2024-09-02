package main

import "testing"

type BDDScenario struct {
	T     *testing.T
	Steps []BDDStep
}

func NewScenario() *BDDScenario {
	return &BDDScenario{}
}

func (sc *BDDScenario) AddStep(step BDDStep) *BDDScenario {
	sc.Steps = append(sc.Steps, step)
	return sc
}

func (sc *BDDScenario) Run() {
	name := ""
	for _, step := range sc.Steps {
		name += step.GetFunctionName(step) + " "
	}
	sc.T.Run(name, func(t *testing.T) {
		for _, step := range sc.Steps {
			step(t)
		}
	})
}
