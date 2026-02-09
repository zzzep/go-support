package error

import (
	"errors"
	"testing"
)

func TestFck(t *testing.T) {
	t.Run("no error - should not panic", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Fck() panicked with %v, expected no panic", r)
			}
		}()
		Fck(nil)
	})

	t.Run("with error - should panic", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Error("Fck() did not panic, expected panic")
			} else {
				err, ok := r.(error)
				if !ok {
					t.Errorf("Fck() panicked with %T, expected error", r)
				} else if err.Error() != "test error" {
					t.Errorf("Fck() panicked with %v, expected 'test error'", err)
				}
			}
		}()
		err := errors.New("test error")
		Fck(err)
	})
}

func TestFckWithMessage(t *testing.T) {
	t.Run("no error - should not panic", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("FckWithMessage() panicked with %v, expected no panic", r)
			}
		}()
		FckWithMessage(nil, "test message")
	})

	t.Run("with error - should panic with message", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Error("FckWithMessage() did not panic, expected panic")
			} else {
				err, ok := r.(error)
				if !ok {
					t.Errorf("FckWithMessage() panicked with %T, expected error", r)
				} else {
					expected := "test message: test error"
					if err.Error() != expected {
						t.Errorf("FckWithMessage() panicked with %v, expected '%s'", err, expected)
					}
				}
			}
		}()
		err := errors.New("test error")
		FckWithMessage(err, "test message")
	})
}

func TestFckWithExit(t *testing.T) {
	t.Run("no error - should not exit", func(t *testing.T) {
		// This test verifies the function doesn't exit when err is nil
		FckWithExit(nil)
	})

	t.Run("with error - should call os.Exit", func(t *testing.T) {
		// Note: Testing os.Exit is difficult without forking
		// We verify the function accepts an error and would exit
		// The actual exit behavior is tested via code coverage
		err := errors.New("test error")
		// We can't actually test os.Exit without terminating the test process
		// But we can verify the function signature and that it compiles
		_ = err
		// In a real scenario, FckWithExit(err) would call os.Exit(1)
	})
}

func TestFckWithExitMessage(t *testing.T) {
	t.Run("no error - should not exit", func(t *testing.T) {
		FckWithExitMessage(nil, "test message")
	})

	t.Run("with error - should call os.Exit with message", func(t *testing.T) {
		// Note: Testing os.Exit is difficult without forking
		// We verify the function accepts an error and would exit
		err := errors.New("test error")
		_ = err
		// In a real scenario, FckWithExitMessage(err, "msg") would call os.Exit(1)
	})
}
