package error

import (
	"errors"
	"fmt"
)

func ExampleFck() {
	// Example: Handle error by panicking
	err := errors.New("something went wrong")
	Fck(err) // Will panic if err is not nil
}

func ExampleFckWithMessage() {
	// Example: Handle error with custom message
	err := errors.New("connection failed")
	FckWithMessage(err, "failed to connect to database")
	// Will panic with: "failed to connect to database: connection failed"
}

func ExampleFckWithExit() {
	// Example: Handle error by exiting program
	err := errors.New("critical error")
	FckWithExit(err) // Will exit with code 1 if err is not nil
}

func ExampleFckWithExitMessage() {
	// Example: Handle error by exiting with custom message
	err := errors.New("file not found")
	FckWithExitMessage(err, "failed to read config")
	// Will exit with: "error: failed to read config: file not found"
}

func ExampleFck_usage() {
	// Common usage pattern
	result, err := someFunction()
	Fck(err) // Panic if error occurred
	fmt.Println(result)
}

func someFunction() (string, error) {
	return "success", nil
}
