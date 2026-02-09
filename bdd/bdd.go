package bdd

import (
	"path/filepath"
	"runtime"
	"strings"
	"testing"
)

type BDD struct {
	T         *testing.T
	Scenarios []Scenario
}

func NewBDD(t *testing.T) *BDD {
	return &BDD{
		T: t,
	}
}

func (b *BDD) AddScenario(scenario Scenario) *BDD {
	b.Scenarios = append(b.Scenarios, scenario)
	return b
}

func (b *BDD) Run() {
	// Store the test file path before running scenarios
	testFile, testLine := getTestLocation()

	for _, scenario := range b.Scenarios {
		if scenario.T == nil {
			scenario.T = b.T
		}
		scenario.Run()
	}

	// Log test location at the end for easy navigation in IDEs
	// Format: file:line (clickable in most IDEs)
	if testFile != "" {
		b.T.Logf("%s:%d", testFile, testLine)
	}
}

// getTestLocation finds the test function location in the call stack
func getTestLocation() (file string, line int) {
	// Go up the call stack to find the test function
	// Call stack: getTestLocation(0) -> Run(1) -> test function(2+)
	for i := 2; i <= 10; i++ {
		pc, file, line, ok := runtime.Caller(i)
		if !ok {
			break
		}

		// Check if this is a test function (starts with "Test")
		if fn := runtime.FuncForPC(pc); fn != nil {
			fnName := fn.Name()
			parts := strings.Split(fnName, ".")
			if len(parts) > 0 {
				lastPart := parts[len(parts)-1]
				if strings.HasPrefix(lastPart, "Test") {
					// Found the test function
					relPath := getRelativeTestPath(file)
					return relPath, line
				}
			}
		}
	}
	return "", 0
}

// getRelativeTestPath returns the relative path from the project root
func getRelativeTestPath(fullPath string) string {
	parts := strings.Split(fullPath, "/")

	// Find "go-support" in the path
	for i := len(parts) - 1; i >= 0; i-- {
		if parts[i] == "go-support" {
			return strings.Join(parts[i:], "/")
		}
	}

	// If not found, return relative path
	return filepath.Base(fullPath)
}
