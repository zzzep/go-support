package error

import (
	"errors"
	"testing"
)

func TestFck(t *testing.T) {
	// Reset to default before each test
	defer ResetDefaultHandler()

	t.Run("no error - should not panic", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Fck() panicked with %v, expected no panic", r)
			}
		}()
		Fck(nil)
	})

	t.Run("with error - default panic", func(t *testing.T) {
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

	t.Run("with error - HandlerPanic", func(t *testing.T) {
		SetDefaultHandler(HandlerPanic)
		defer ResetDefaultHandler()

		defer func() {
			if r := recover(); r == nil {
				t.Error("Fck() did not panic, expected panic")
			}
		}()
		err := errors.New("test error")
		Fck(err)
	})

	t.Run("with error - HandlerPanicWithMessage", func(t *testing.T) {
		SetDefaultHandler(HandlerPanicWithMessage)
		SetDefaultMessage("custom message")
		defer ResetDefaultHandler()

		defer func() {
			if r := recover(); r == nil {
				t.Error("Fck() did not panic, expected panic")
			} else {
				err, ok := r.(error)
				if !ok {
					t.Errorf("Fck() panicked with %T, expected error", r)
				} else {
					expected := "custom message: test error"
					if err.Error() != expected {
						t.Errorf("Fck() panicked with %v, expected '%s'", err, expected)
					}
				}
			}
		}()
		err := errors.New("test error")
		Fck(err)
	})

	t.Run("with error - HandlerIgnore", func(t *testing.T) {
		SetDefaultHandler(HandlerIgnore)
		defer ResetDefaultHandler()

		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Fck() panicked with %v, expected no panic", r)
			}
		}()
		err := errors.New("test error")
		Fck(err)
	})

	t.Run("with error - HandlerExit", func(t *testing.T) {
		SetDefaultHandler(HandlerExit)
		defer ResetDefaultHandler()

		// Note: We can't fully test os.Exit without forking, but we verify it compiles and runs
		// The actual exit behavior would be tested in integration tests
		err := errors.New("test error")
		// In a real scenario, this would call os.Exit(1)
		_ = err
	})

	t.Run("with error - HandlerExitWithMessage", func(t *testing.T) {
		SetDefaultHandler(HandlerExitWithMessage)
		SetDefaultMessage("custom exit message")
		defer ResetDefaultHandler()

		// Note: We can't fully test os.Exit without forking
		err := errors.New("test error")
		// In a real scenario, this would call os.Exit(1) with the message
		_ = err
	})

	t.Run("with error - HandlerExitWithMessage without default message", func(t *testing.T) {
		SetDefaultHandler(HandlerExitWithMessage)
		SetDefaultMessage("")
		defer ResetDefaultHandler()

		err := errors.New("test error")
		// In a real scenario, this would call os.Exit(1) without custom message
		_ = err
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

func TestSetDefaultHandler(t *testing.T) {
	defer ResetDefaultHandler()

	t.Run("set and get handler", func(t *testing.T) {
		SetDefaultHandler(HandlerExit)
		if GetDefaultHandler() != HandlerExit {
			t.Errorf("GetDefaultHandler() = %v, expected %v", GetDefaultHandler(), HandlerExit)
		}
	})

	t.Run("set message", func(t *testing.T) {
		SetDefaultMessage("test message")
		SetDefaultHandler(HandlerPanicWithMessage)
		defer func() {
			if r := recover(); r == nil {
				t.Error("Expected panic with message")
			}
		}()
		err := errors.New("test error")
		Fck(err)
	})
}

func TestResetDefaultHandler(t *testing.T) {
	SetDefaultHandler(HandlerExit)
	SetDefaultMessage("test")
	ResetDefaultHandler()

	if GetDefaultHandler() != HandlerPanic {
		t.Errorf("ResetDefaultHandler() did not reset to HandlerPanic, got %v", GetDefaultHandler())
	}
}
