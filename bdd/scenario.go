package bdd

import (
	"fmt"
	"reflect"
	"strings"
	"sync"
	"testing"
)

var (
	testCounter int
	counterMutex sync.Mutex
)

type stepType int

const (
	stepTypeNone stepType = iota
	stepTypeGiven
	stepTypeWhen
	stepTypeThen
)

type stepInfo struct {
	step Step
	typ  stepType
}
type Scenario struct {
	T        *testing.T
	Steps    []stepInfo
	Mocks    map[any]Mock
	lastType stepType
}

func NewScenario() *Scenario {
	return &Scenario{
		Mocks: make(map[any]Mock),
	}
}

func (sc *Scenario) AddStep(step Step) *Scenario {
	sc.Steps = append(sc.Steps, stepInfo{step: step, typ: stepTypeNone})
	return sc
}

// Given adds a Given step. If there's already a Given step, it's treated as "AND"
// Accepts Step directly for simplified usage: scenario.Given(metodo)
func (sc *Scenario) Given(step Step) *Scenario {
	if sc.lastType == stepTypeGiven {
		// Already a Given step, treat as "AND"
		sc.Steps = append(sc.Steps, stepInfo{step: step, typ: stepTypeGiven})
	} else {
		sc.Steps = append(sc.Steps, stepInfo{step: step, typ: stepTypeGiven})
		sc.lastType = stepTypeGiven
	}
	return sc
}

// When adds a When step. If there's already a When step, it's treated as "AND"
// Accepts Step directly for simplified usage: scenario.When(metodo)
func (sc *Scenario) When(step Step) *Scenario {
	if sc.lastType == stepTypeWhen {
		// Already a When step, treat as "AND"
		sc.Steps = append(sc.Steps, stepInfo{step: step, typ: stepTypeWhen})
	} else {
		sc.Steps = append(sc.Steps, stepInfo{step: step, typ: stepTypeWhen})
		sc.lastType = stepTypeWhen
	}
	return sc
}

// Then adds a Then step. If there's already a Then step, it's treated as "AND"
// Accepts Step directly for simplified usage: scenario.Then(metodo)
func (sc *Scenario) Then(step Step) *Scenario {
	if sc.lastType == stepTypeThen {
		// Already a Then step, treat as "AND"
		sc.Steps = append(sc.Steps, stepInfo{step: step, typ: stepTypeThen})
	} else {
		sc.Steps = append(sc.Steps, stepInfo{step: step, typ: stepTypeThen})
		sc.lastType = stepTypeThen
	}
	return sc
}

// AddMock adds a mock to the scenario using the interface type as key
// Usage: scenario.AddMock((*UserRepository)(nil), mockRepo)
func (sc *Scenario) AddMock(interfaceType interface{}, mock Mock) *Scenario {
	sc.Mocks[interfaceType] = mock
	return sc
}

func (sc *Scenario) Run() {
	name := ""
	for _, stepInfo := range sc.Steps {
		name += GetFunctionName(stepInfo.step) + " "
	}
	
	// Trim trailing space
	name = strings.TrimSpace(name)
	
	// Get and increment test counter (once per test/scenario)
	counterMutex.Lock()
	testCounter++
	currentTest := testCounter
	counterMutex.Unlock()
	
	// Log test with counter format: #N# - TEST_NAME
	fmt.Printf("#%d# - %s\n", currentTest, name)
	
	sc.T.Run(name, func(t *testing.T) {
		var lastStepType stepType = stepTypeNone
		for _, stepInfo := range sc.Steps {
			// Auto-log based on step type
			stepName := GetFunctionName(stepInfo.step)
			prefix := getStepPrefix(stepInfo.typ, lastStepType)
			SimpleLog(prefix + formatStepName(stepName))

			// Execute the step
			stepInfo.step(*sc, t)

			lastStepType = stepInfo.typ
		}
	})
}

// getStepPrefix returns the prefix (Given/When/Then/And) based on step type
func getStepPrefix(currentType, lastType stepType) string {
	if currentType == stepTypeNone {
		return ""
	}

	// If same type as last, it's "And"
	if currentType == lastType && lastType != stepTypeNone {
		return "And: "
	}

	switch currentType {
	case stepTypeGiven:
		return "Given: "
	case stepTypeWhen:
		return "When: "
	case stepTypeThen:
		return "Then: "
	default:
		return ""
	}
}

// formatStepName converts function name to readable format
// e.g., "IHaveAUser" -> "I have a user"
func formatStepName(fnName string) string {
	if fnName == "" {
		return ""
	}

	// Convert camelCase to readable text with spaces
	result := ""
	for i, r := range fnName {
		if i > 0 && r >= 'A' && r <= 'Z' {
			prev := rune(fnName[i-1])
			// Add space before uppercase if:
			// 1. Previous was lowercase, OR
			// 2. Previous was uppercase and next is lowercase (single letter like "O" followed by word)
			if prev >= 'a' && prev <= 'z' {
				result += " "
			} else if prev >= 'A' && prev <= 'Z' {
				// Check if this is a single uppercase letter followed by a word
				// e.g., "OUsuario" -> "O Usuario"
				if i < len(fnName)-1 {
					next := rune(fnName[i+1])
					// If next character is lowercase, this uppercase is start of new word
					if next >= 'a' && next <= 'z' {
						result += " "
					}
				}
			}
		}
		result += string(r)
	}

	// Convert to lowercase, but keep acronyms (ID, API, etc.) uppercase
	words := strings.Fields(result)
	for i, word := range words {
		// Keep short all-uppercase words as acronyms (but not single letters like "O", "A")
		// Single letters should be lowercase unless they're part of an acronym
		if len(word) > 1 && len(word) <= 3 && strings.ToUpper(word) == word && word != strings.ToLower(word) {
			words[i] = word
		} else {
			words[i] = strings.ToLower(word)
		}
	}

	// Capitalize first letter of first word
	if len(words) > 0 && len(words[0]) > 0 {
		words[0] = strings.ToUpper(words[0][:1]) + words[0][1:]
	}

	return strings.Join(words, " ")
}

func GenerateMock[T any](mock Mock) (i T, m Mock) {
	if !reflect.TypeOf(mock).Implements(reflect.TypeOf((*T)(nil)).Elem()) {
		panic("Mock does not implement interface")
	}
	return i, mock
}
