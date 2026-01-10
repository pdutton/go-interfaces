package sync

import (
	"sync"
	"testing"
	"time"
)

func TestNewMutex(t *testing.T) {
	s := NewSync()
	m := s.NewMutex()
	_ = m
}

func TestMutex_Lock_Unlock(t *testing.T) {
	s := NewSync()
	m := s.NewMutex()

	// Test basic locking and unlocking
	m.Lock()
	// Critical section
	m.Unlock()

	// Verify we can lock again (wasn't deadlocked)
	m.Lock()
	m.Unlock()
}

func TestMutex_TryLock(t *testing.T) {
	s := NewSync()
	m := s.NewMutex()

	// Test TryLock when lock is available
	if !m.TryLock() {
		t.Fatal("expected lock to be available")
	}
	m.Unlock()

	// Test TryLock when lock is held
	m.Lock()
	if m.TryLock() {
		t.Error("expected TryLock to fail when lock is held")
		m.Unlock() // unlock the try-lock
	}
	m.Unlock()
}

func TestMutex_Concurrent(t *testing.T) {
	s := NewSync()
	m := s.NewMutex()

	counter := 0
	const numGoroutines = 100
	const numIncrements = 100

	var wg sync.WaitGroup
	wg.Add(numGoroutines)

	for i := 0; i < numGoroutines; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < numIncrements; j++ {
				m.Lock()
				counter++
				m.Unlock()
			}
		}()
	}

	wg.Wait()

	expected := numGoroutines * numIncrements
	if counter != expected {
		t.Errorf("expected counter to be %d, got %d (mutex failed to protect)", expected, counter)
	}
}

func TestMutex_MultipleGoroutines(t *testing.T) {
	s := NewSync()
	m := s.NewMutex()

	locked := false
	conflicts := 0
	var conflictMutex sync.Mutex

	done := make(chan bool)

	// Launch multiple goroutines trying to acquire the lock
	for i := 0; i < 10; i++ {
		go func() {
			m.Lock()

			// Check if another goroutine has the lock
			conflictMutex.Lock()
			if locked {
				conflicts++
			}
			locked = true
			conflictMutex.Unlock()

			// Hold the lock briefly
			time.Sleep(1 * time.Millisecond)

			conflictMutex.Lock()
			locked = false
			conflictMutex.Unlock()

			m.Unlock()
			done <- true
		}()
	}

	// Wait for all goroutines
	for i := 0; i < 10; i++ {
		<-done
	}

	// Mutex should prevent conflicts
	if conflicts > 0 {
		t.Errorf("expected no conflicts, but got %d", conflicts)
	}
}

func TestMutex_WithLocked_Option(t *testing.T) {
	s := NewSync()
	m := s.NewMutex(WithLocked())

	// Verify the mutex is initially locked by trying to TryLock
	if m.TryLock() {
		t.Error("expected mutex to be initially locked with WithLocked option")
		m.Unlock()
	}

	// Unlock and verify we can lock it again
	m.Unlock()
	if !m.TryLock() {
		t.Error("expected to be able to lock after unlocking")
	}
	m.Unlock()
}
