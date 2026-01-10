package sync

import (
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

func TestNewOnce(t *testing.T) {
	s := NewSync()
	o := s.NewOnce()
	_ = o
}

func TestOnce_Do(t *testing.T) {
	s := NewSync()
	o := s.NewOnce()

	callCount := 0
	f := func() {
		callCount++
	}

	// Call Do multiple times
	o.Do(f)
	o.Do(f)
	o.Do(f)

	// Function should only be called once
	if callCount != 1 {
		t.Errorf("expected function to be called once, but was called %d times", callCount)
	}
}

func TestOnce_Do_DifferentFunctions(t *testing.T) {
	s := NewSync()
	o := s.NewOnce()

	call1Count := 0
	call2Count := 0

	f1 := func() {
		call1Count++
	}

	f2 := func() {
		call2Count++
	}

	// First function should be called
	o.Do(f1)

	// Second function should not be called
	o.Do(f2)

	if call1Count != 1 {
		t.Errorf("expected first function to be called once, got %d", call1Count)
	}
	if call2Count != 0 {
		t.Errorf("expected second function to not be called, got %d", call2Count)
	}
}

func TestOnce_Concurrent(t *testing.T) {
	s := NewSync()
	o := s.NewOnce()

	callCount := int32(0)
	f := func() {
		atomic.AddInt32(&callCount, 1)
	}

	const numGoroutines = 100
	var wg sync.WaitGroup
	wg.Add(numGoroutines)

	// Launch many goroutines calling Do concurrently
	for i := 0; i < numGoroutines; i++ {
		go func() {
			defer wg.Done()
			o.Do(f)
		}()
	}

	wg.Wait()

	// Despite concurrent calls, function should only be called once
	if callCount != 1 {
		t.Errorf("expected function to be called once despite %d concurrent calls, but was called %d times",
			numGoroutines, callCount)
	}
}

func TestOnce_PanicInDo(t *testing.T) {
	s := NewSync()
	o := s.NewOnce()

	callCount := 0

	// First call panics
	func() {
		defer func() {
			if r := recover(); r == nil {
				t.Error("expected panic")
			}
		}()

		o.Do(func() {
			callCount++
			panic("test panic")
		})
	}()

	// Subsequent calls should not execute the function
	// (Once is still considered "done" even if the function panicked)
	o.Do(func() {
		callCount++
	})

	if callCount != 1 {
		t.Errorf("expected function to be called once (even with panic), got %d", callCount)
	}
}

func TestOnce_LongRunningFunction(t *testing.T) {
	s := NewSync()
	o := s.NewOnce()

	started := make(chan bool)
	finished := make(chan bool)
	secondDone := make(chan bool)

	// First goroutine calls Do with a long-running function
	go func() {
		o.Do(func() {
			started <- true
			// Simulate long-running function
			<-finished
		})
	}()

	// Wait for first call to start
	<-started

	// Second goroutine tries to call Do
	// It should block until the first call completes
	go func() {
		o.Do(func() {
			t.Error("second function should not be called")
		})
		secondDone <- true
	}()

	// Give second goroutine time to block on Do
	time.Sleep(50 * time.Millisecond)

	// Finish the first call
	finished <- true

	// Second call should complete shortly after
	select {
	case <-secondDone:
		// Success - second Do call returned
	case <-time.After(100 * time.Millisecond):
		t.Error("expected second Do call to complete after first call finished")
	}
}
