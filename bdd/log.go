package bdd

import (
	"fmt"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
)

// Log logs a message with simplified relative path
// The Go testing framework will still show file:line, but this helper
// ensures consistent formatting and can be used for cleaner output
// Usage: bdd.Log(t, "Given: I have a user")
func Log(t *testing.T, msg string) {
	// Get caller file path
	_, file, _, ok := runtime.Caller(1)
	if !ok {
		t.Log(msg)
		return
	}

	// Get relative path from project root (go-support)
	relPath := getRelativePath(file)

	// Log with relative path - Go testing framework will format it
	t.Logf("%s: %s", relPath, msg)
}

// getRelativePath returns the relative path from the project root
func getRelativePath(fullPath string) string {
	parts := strings.Split(fullPath, "/")

	// Find "go-support" in the path
	for i := len(parts) - 1; i >= 0; i-- {
		if parts[i] == "go-support" {
			// Return path from go-support onwards
			return strings.Join(parts[i:], "/")
		}
	}

	// If not found, return just package/filename
	// Extract package and filename
	if len(parts) >= 2 {
		return strings.Join(parts[len(parts)-2:], "/")
	}

	return filepath.Base(fullPath)
}

// SimpleLog logs a message without file path prefix for cleaner output
// This uses fmt.Print which shows only the message, without file:line prefix
// Usage: bdd.SimpleLog("Given: I have a user")
func SimpleLog(msg string) {
	// Use fmt.Print for cleaner output without file:line prefix
	fmt.Println(msg)
}
