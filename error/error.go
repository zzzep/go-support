package error

import (
	"fmt"
	"os"
	"sync"
)

// HandlerType defines the type of error handler
type HandlerType int

const (
	// HandlerPanic panics when an error occurs (default)
	HandlerPanic HandlerType = iota
	// HandlerPanicWithMessage panics with a custom message when an error occurs
	HandlerPanicWithMessage
	// HandlerExit exits the program when an error occurs
	HandlerExit
	// HandlerExitWithMessage exits the program with a custom message when an error occurs
	HandlerExitWithMessage
	// HandlerIgnore does nothing when an error occurs
	HandlerIgnore
)

var (
	defaultHandler     HandlerType = HandlerPanic
	defaultMessage     string      = ""
	handlerMutex       sync.RWMutex
)

// SetDefaultHandler sets the default handler type for Fck
// Usage: error.SetDefaultHandler(error.HandlerExit)
func SetDefaultHandler(handler HandlerType) {
	handlerMutex.Lock()
	defer handlerMutex.Unlock()
	defaultHandler = handler
}

// SetDefaultMessage sets the default message for handlers that support messages
// Usage: error.SetDefaultMessage("failed to process")
func SetDefaultMessage(message string) {
	handlerMutex.Lock()
	defer handlerMutex.Unlock()
	defaultMessage = message
}

// GetDefaultHandler returns the current default handler type
func GetDefaultHandler() HandlerType {
	handlerMutex.RLock()
	defer handlerMutex.RUnlock()
	return defaultHandler
}

// ResetDefaultHandler resets the default handler to HandlerPanic and clears the default message
func ResetDefaultHandler() {
	handlerMutex.Lock()
	defer handlerMutex.Unlock()
	defaultHandler = HandlerPanic
	defaultMessage = ""
}

// Fck handles an error using the configured default handler
// Usage: error.Fck(err)
// You can configure the default behavior with SetDefaultHandler()
func Fck(err error) {
	if err == nil {
		return
	}

	handlerMutex.RLock()
	handler := defaultHandler
	message := defaultMessage
	handlerMutex.RUnlock()

	switch handler {
	case HandlerPanic:
		panic(err)
	case HandlerPanicWithMessage:
		if message != "" {
			panic(fmt.Errorf("%s: %w", message, err))
		}
		panic(err)
	case HandlerExit:
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	case HandlerExitWithMessage:
		if message != "" {
			fmt.Fprintf(os.Stderr, "error: %s: %v\n", message, err)
		} else {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
		}
		os.Exit(1)
	case HandlerIgnore:
		// Do nothing
	default:
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
