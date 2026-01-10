package sync

import (
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

func TestNewWaitGroup(t *testing.T) {
	s := NewSync()
	wg := s.NewWaitGroup()
	_ = wg
}

func TestWaitGroup_Add_Done_Wait(t *testing.T) {
	s := NewSync()
	wg := s.NewWaitGroup()

	completed := false

	wg.Add(1)
	go func() {
		time.Sleep(10 * time.Millisecond)
		completed = true
		wg.Done()
	}()

	wg.Wait()

	if !completed {
		t.Error("expected goroutine to complete before Wait returns")
	}
}

func TestWaitGroup_MultipleGoroutines(t *testing.T) {
	s := NewSync()
	wg := s.NewWaitGroup()

	const numGoroutines = 10
	counter := int32(0)

	wg.Add(numGoroutines)
	for i := 0; i < numGoroutines; i++ {
		go func() {
			defer wg.Done()
			atomic.AddInt32(&counter, 1)
		}()
	}

	wg.Wait()

	if counter != numGoroutines {
		t.Errorf("expected counter to be %d, got %d", numGoroutines, counter)
	}
}

func TestWaitGroup_Concurrent(t *testing.T) {
	s := NewSync()
	wg := s.NewWaitGroup()

	const numGoroutines = 100
	completed := make([]bool, numGoroutines)

	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		index := i
		go func() {
			defer wg.Done()
			time.Sleep(1 * time.Millisecond)
			completed[index] = true
		}()
	}

	wg.Wait()

	// Verify all goroutines completed
	for i, done := range completed {
		if !done {
			t.Errorf("goroutine %d did not complete", i)
		}
	}
}

func TestWaitGroup_MultipleWaiters(t *testing.T) {
	s := NewSync()
	wg := s.NewWaitGroup()

	waiter1Done := false
	waiter2Done := false
	var mu sync.Mutex

	// Start a goroutine that will be waited on
	wg.Add(1)
	go func() {
		time.Sleep(50 * time.Millisecond)
		wg.Done()
	}()

	// First waiter
	go func() {
		wg.Wait()
		mu.Lock()
		waiter1Done = true
		mu.Unlock()
	}()

	// Second waiter
	go func() {
		wg.Wait()
		mu.Lock()
		waiter2Done = true
		mu.Unlock()
	}()

	// Give time for all waiters to complete
	time.Sleep(100 * time.Millisecond)

	mu.Lock()
	done1 := waiter1Done
	done2 := waiter2Done
	mu.Unlock()

	if !done1 {
		t.Error("expected first waiter to complete")
	}
	if !done2 {
		t.Error("expected second waiter to complete")
	}
}

func TestWaitGroup_AddDoneSequence(t *testing.T) {
	s := NewSync()
	wg := s.NewWaitGroup()

	// Test sequential Add/Done calls
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			time.Sleep(1 * time.Millisecond)
			wg.Done()
		}()
	}

	wg.Wait()

	// If we reach here, all goroutines completed
}

func TestWaitGroup_ZeroValue(t *testing.T) {
	s := NewSync()
	wg := s.NewWaitGroup()

	// Wait on a WaitGroup with counter 0 should return immediately
	done := make(chan bool)
	go func() {
		wg.Wait()
		done <- true
	}()

	select {
	case <-done:
		// Success - Wait returned immediately
	case <-time.After(100 * time.Millisecond):
		t.Error("expected Wait to return immediately for zero counter")
	}
}

func TestWaitGroup_WithCount_Option(t *testing.T) {
	const initialCount = 5
	completed := int32(0)

	s := NewSync()
	wg := s.NewWaitGroup(WithCount(initialCount))

	// Launch goroutines without calling Add
	for i := 0; i < initialCount; i++ {
		go func() {
			defer wg.Done()
			atomic.AddInt32(&completed, 1)
		}()
	}

	wg.Wait()

	if completed != initialCount {
		t.Errorf("expected %d goroutines to complete, got %d", initialCount, completed)
	}
}

func TestWaitGroup_AddNegative(t *testing.T) {
	s := NewSync()
	wg := s.NewWaitGroup()

	// Add 2, then Done 2 using Add(-1)
	wg.Add(2)

	go func() {
		time.Sleep(10 * time.Millisecond)
		wg.Add(-1)
	}()

	go func() {
		time.Sleep(10 * time.Millisecond)
		wg.Add(-1)
	}()

	wg.Wait()
}
