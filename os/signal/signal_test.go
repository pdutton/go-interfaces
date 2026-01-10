package signal

import (
	"context"
	"os"
	"syscall"
	"testing"
	"time"
)

func TestNewSignal(t *testing.T) {
	s := NewSignal()
	_ = s
}

func TestSignal_Notify(t *testing.T) {
	s := NewSignal()

	c := make(chan os.Signal, 1)
	defer s.Stop(c)

	// Register for SIGTERM (safe to use in tests)
	s.Notify(c, syscall.SIGTERM)

	// Just verify Notify doesn't panic
	// We won't actually send the signal
}

func TestSignal_Stop(t *testing.T) {
	s := NewSignal()

	c := make(chan os.Signal, 1)
	s.Notify(c, syscall.SIGTERM)

	// Stop should not panic
	s.Stop(c)

	// Calling Stop again should also not panic
	s.Stop(c)
}

func TestSignal_Ignore(t *testing.T) {
	s := NewSignal()

	// Ignore SIGHUP (safe to use in tests)
	// On Windows, only SIGINT is supported, but Ignore shouldn't panic
	s.Ignore(syscall.SIGHUP)

	// Verify Ignored returns true
	// Note: This may not work on all platforms
	// Just verify it doesn't panic
	_ = s.Ignored(syscall.SIGHUP)
}

func TestSignal_Ignored(t *testing.T) {
	s := NewSignal()

	// Test with a signal that's not ignored
	// SIGTERM is typically not ignored
	ignored := s.Ignored(syscall.SIGTERM)
	// Just verify the method works, result varies by platform
	_ = ignored
}

func TestSignal_Reset(t *testing.T) {
	s := NewSignal()

	// Ignore a signal first
	s.Ignore(syscall.SIGHUP)

	// Reset should restore default behavior
	s.Reset(syscall.SIGHUP)

	// Verify Reset doesn't panic
}

func TestSignal_NotifyContext(t *testing.T) {
	s := NewSignal()

	parent := context.Background()

	// Create a context that will be canceled on SIGTERM
	ctx, stop := s.NotifyContext(parent, syscall.SIGTERM)
	defer stop()

	if ctx == nil {
		t.Fatal("NotifyContext returned nil context")
	}
	if stop == nil {
		t.Fatal("NotifyContext returned nil stop function")
	}

	// Verify the context hasn't been canceled yet
	select {
	case <-ctx.Done():
		t.Error("Context was already done")
	default:
		// Context is not done, as expected
	}

	// Call stop function
	stop()

	// Give it a moment to cancel
	time.Sleep(10 * time.Millisecond)

	// Verify context is now done
	select {
	case <-ctx.Done():
		// Context is done, as expected
	case <-time.After(100 * time.Millisecond):
		t.Error("Context should have been canceled after calling stop()")
	}
}

func TestSignal_NotifyContext_Cancel(t *testing.T) {
	s := NewSignal()

	parent, parentCancel := context.WithCancel(context.Background())
	defer parentCancel()

	ctx, stop := s.NotifyContext(parent, syscall.SIGTERM)
	defer stop()

	// Cancel parent context
	parentCancel()

	// Child context should also be canceled
	select {
	case <-ctx.Done():
		// Expected
	case <-time.After(100 * time.Millisecond):
		t.Error("Child context should be canceled when parent is canceled")
	}
}

func TestSignal_MultipleNotify(t *testing.T) {
	s := NewSignal()

	c1 := make(chan os.Signal, 1)
	c2 := make(chan os.Signal, 1)
	defer s.Stop(c1)
	defer s.Stop(c2)

	// Register both channels
	s.Notify(c1, syscall.SIGTERM)
	s.Notify(c2, syscall.SIGTERM)

	// Both should be registered without error
}

func TestSignal_NotifyMultipleSignals(t *testing.T) {
	s := NewSignal()

	c := make(chan os.Signal, 2)
	defer s.Stop(c)

	// Register for multiple signals
	s.Notify(c, syscall.SIGTERM, syscall.SIGINT)

	// Should register without error
}

func TestSignal_ResetAll(t *testing.T) {
	s := NewSignal()

	// Ignore some signals
	s.Ignore(syscall.SIGHUP, syscall.SIGUSR1)

	// Reset all
	s.Reset()

	// Should not panic
}

func TestSignal_IgnoreAll(t *testing.T) {
	s := NewSignal()

	// Ignore without arguments should ignore no signals
	s.Ignore()

	// Should not panic
}

func TestSignal_StopNilChannel(t *testing.T) {
	s := NewSignal()

	// Calling Stop on nil channel should not panic
	// (the actual stdlib behavior)
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("Stop(nil) panicked: %v", r)
		}
	}()

	var c chan os.Signal
	s.Stop(c)
}
