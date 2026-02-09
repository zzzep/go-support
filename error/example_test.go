package error

import (
	"errors"
	"fmt"
)

func ExampleFck() {
	// Example: Handle error by panicking (default behavior)
	err := errors.New("something went wrong")
	Fck(err) // Will panic if err is not nil
}

func ExampleFck_withCustomHandler() {
	// Example: Configure Fck to exit instead of panic
	SetDefaultHandler(HandlerExit)
	defer ResetDefaultHandler()

	err := errors.New("critical error")
	Fck(err) // Will exit with code 1 instead of panic
}

func ExampleFck_withMessage() {
	// Example: Configure Fck to panic with custom message
	SetDefaultHandler(HandlerPanicWithMessage)
	SetDefaultMessage("failed to process")
	defer ResetDefaultHandler()

	err := errors.New("connection failed")
	Fck(err) // Will panic with: "failed to process: connection failed"
}

func ExampleFck_ignoreErrors() {
	// Example: Configure Fck to ignore errors
	SetDefaultHandler(HandlerIgnore)
	defer ResetDefaultHandler()

	err := errors.New("non-critical error")
	Fck(err) // Will do nothing
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
	Fck(err) // Uses configured default handler
	fmt.Println(result)
}

func someFunction() (string, error) {
	return "success", nil
}
