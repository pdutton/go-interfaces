package sync

import (
	"sync"
	"testing"
)

func TestNewSync(t *testing.T) {
	impl := NewSync()
	// syncFacade doesn't implement Sync interface (returns concrete types)
	// Just verify we can create it
	_ = impl
}

func TestSync_OnceFunc(t *testing.T) {
	s := NewSync()

	callCount := 0
	f := func() {
		callCount++
	}

	onceFn := s.OnceFunc(f)

	// Call it multiple times
	onceFn()
	onceFn()
	onceFn()

	// Should only be called once
	if callCount != 1 {
		t.Errorf("expected function to be called once, but was called %d times", callCount)
	}
}

func TestSync_OnceFunc_Concurrent(t *testing.T) {
	s := NewSync()

	callCount := 0
	var mu sync.Mutex
	f := func() {
		mu.Lock()
		callCount++
		mu.Unlock()
	}

	onceFn := s.OnceFunc(f)

	// Call concurrently
	done := make(chan bool)
	for i := 0; i < 10; i++ {
		go func() {
			onceFn()
			done <- true
		}()
	}

	// Wait for all goroutines
	for i := 0; i < 10; i++ {
		<-done
	}

	// Should only be called once despite concurrent calls
	if callCount != 1 {
		t.Errorf("expected function to be called once, but was called %d times", callCount)
	}
}

func TestSync_NewCond(t *testing.T) {
	s := NewSync()
	m := s.NewMutex()

	cond := s.NewCond(m)
	// Just verify we can create it
	_ = cond
}

func TestSync_NewMap(t *testing.T) {
	s := NewSync()

	m := s.NewMap()
	_ = m
}

func TestSync_NewMutex(t *testing.T) {
	s := NewSync()

	m := s.NewMutex()
	_ = m
}

func TestSync_NewOnce(t *testing.T) {
	s := NewSync()

	o := s.NewOnce()
	_ = o
}

func TestSync_NewPool(t *testing.T) {
	s := NewSync()

	p := s.NewPool()
	_ = p
}

func TestSync_NewRWMutex(t *testing.T) {
	s := NewSync()

	rw := s.NewRWMutex()
	_ = rw
}

func TestSync_NewWaitGroup(t *testing.T) {
	s := NewSync()

	wg := s.NewWaitGroup()
	_ = wg
}
