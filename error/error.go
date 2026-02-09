package error

import (
	"fmt"
	"os"
)

// Fck handles an error by panicking if the error is not nil
// Usage: error.Fck(err)
func Fck(err error) {
	if err != nil {
		panic(err)
	}
}

// FckWithMessage handles an error by panicking with a custom message if the error is not nil
// Usage: error.FckWithMessage(err, "failed to process")
func FckWithMessage(err error, message string) {
	if err != nil {
		panic(fmt.Errorf("%s: %w", message, err))
	}
}

// FckWithExit handles an error by exiting the program if the error is not nil
// Usage: error.FckWithExit(err)
func FckWithExit(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

// FckWithExitMessage handles an error by exiting the program with a custom message if the error is not nil
// Usage: error.FckWithExitMessage(err, "failed to process")
func FckWithExitMessage(err error, message string) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s: %v\n", message, err)
		os.Exit(1)
	}
}
