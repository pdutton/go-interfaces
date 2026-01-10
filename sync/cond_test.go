package sync

import (
	"sync"
	"testing"
	"time"
)

func TestNewCond(t *testing.T) {
	s := NewSync()
	m := s.NewMutex()
	c := s.NewCond(m)

	if c == nil {
		t.Fatal("expected non-nil Cond")
	}
}

func TestCond_Wait_Signal(t *testing.T) {
	s := NewSync()
	m := s.NewMutex()
	c := s.NewCond(m)

	ready := false
	done := make(chan bool)

	// Waiter goroutine
	go func() {
		m.Lock()
		for !ready {
			c.Wait()
		}
		m.Unlock()
		done <- true
	}()

	// Give waiter time to start waiting
	time.Sleep(10 * time.Millisecond)

	// Signal the waiter
	m.Lock()
	ready = true
	c.Signal()
	m.Unlock()

	// Wait for waiter to complete
	select {
	case <-done:
		// Success
	case <-time.After(100 * time.Millisecond):
		t.Error("waiter did not wake up after signal")
	}
}

func TestCond_Wait_Broadcast(t *testing.T) {
	s := NewSync()
	m := s.NewMutex()
	c := s.NewCond(m)

	ready := false
	const numWaiters = 5
	done := make(chan bool, numWaiters)

	// Launch multiple waiters
	for i := 0; i < numWaiters; i++ {
		go func() {
			m.Lock()
			for !ready {
				c.Wait()
			}
			m.Unlock()
			done <- true
		}()
	}

	// Give waiters time to start waiting
	time.Sleep(20 * time.Millisecond)

	// Broadcast to all waiters
	m.Lock()
	ready = true
	c.Broadcast()
	m.Unlock()

	// Wait for all waiters to complete
	completed := 0
	timeout := time.After(200 * time.Millisecond)
	for i := 0; i < numWaiters; i++ {
		select {
		case <-done:
			completed++
		case <-timeout:
			t.Errorf("only %d of %d waiters woke up after broadcast", completed, numWaiters)
			return
		}
	}

	if completed != numWaiters {
		t.Errorf("expected %d waiters to wake up, got %d", numWaiters, completed)
	}
}

func TestCond_MultipleWaiters_Signal(t *testing.T) {
	s := NewSync()
	m := s.NewMutex()
	c := s.NewCond(m)

	counter := 0
	const numWaiters = 3
	done := make(chan bool, numWaiters)

	// Launch multiple waiters
	for i := 0; i < numWaiters; i++ {
		go func() {
			m.Lock()
			for counter == 0 {
				c.Wait()
			}
			counter--
			m.Unlock()
			done <- true
		}()
	}

	// Give waiters time to start waiting
	time.Sleep(20 * time.Millisecond)

	// Signal them one at a time
	for i := 0; i < numWaiters; i++ {
		m.Lock()
		counter++
		c.Signal()
		m.Unlock()

		// Wait for one waiter to complete
		select {
		case <-done:
			// Success
		case <-time.After(100 * time.Millisecond):
			t.Errorf("waiter %d did not wake up after signal", i+1)
			return
		}
	}
}

func TestCond_Signal_NoWaiters(t *testing.T) {
	s := NewSync()
	m := s.NewMutex()
	c := s.NewCond(m)

	// Signaling with no waiters should not block or panic
	c.Signal()
	c.Broadcast()
}

func TestCond_Wait_RequiresLock(t *testing.T) {
	s := NewSync()
	m := s.NewMutex()
	c := s.NewCond(m)

	ready := false

	m.Lock()

	go func() {
		time.Sleep(10 * time.Millisecond)
		m.Lock()
		ready = true
		c.Signal()
		m.Unlock()
	}()

	// Wait releases the lock and reacquires it
	for !ready {
		c.Wait()
	}

	m.Unlock()
}

func TestCond_Concurrent_ProducerConsumer(t *testing.T) {
	s := NewSync()
	m := s.NewMutex()
	c := s.NewCond(m)

	queue := []int{}
	const numItems = 10
	consumed := make(chan int, numItems)

	// Consumer goroutine
	go func() {
		for i := 0; i < numItems; i++ {
			m.Lock()
			for len(queue) == 0 {
				c.Wait()
			}
			item := queue[0]
			queue = queue[1:]
			m.Unlock()

			consumed <- item
		}
	}()

	// Producer goroutine
	go func() {
		for i := 0; i < numItems; i++ {
			m.Lock()
			queue = append(queue, i)
			c.Signal()
			m.Unlock()

			time.Sleep(1 * time.Millisecond)
		}
	}()

	// Verify all items were consumed
	seen := make(map[int]bool)
	timeout := time.After(500 * time.Millisecond)

	for i := 0; i < numItems; i++ {
		select {
		case item := <-consumed:
			seen[item] = true
		case <-timeout:
			t.Errorf("only consumed %d of %d items", len(seen), numItems)
			return
		}
	}

	if len(seen) != numItems {
		t.Errorf("expected to consume %d items, got %d", numItems, len(seen))
	}
}

func TestCond_Broadcast_WakesAll(t *testing.T) {
	s := NewSync()
	m := s.NewMutex()
	c := s.NewCond(m)

	const numWaiters = 10
	var wg sync.WaitGroup
	wg.Add(numWaiters)

	ready := false

	// Launch waiters
	for i := 0; i < numWaiters; i++ {
		go func() {
			defer wg.Done()
			m.Lock()
			for !ready {
				c.Wait()
			}
			m.Unlock()
		}()
	}

	// Give waiters time to start waiting
	time.Sleep(20 * time.Millisecond)

	// Broadcast
	m.Lock()
	ready = true
	c.Broadcast()
	m.Unlock()

	// Wait for all with timeout
	done := make(chan bool)
	go func() {
		wg.Wait()
		done <- true
	}()

	select {
	case <-done:
		// Success - all waiters woke up
	case <-time.After(200 * time.Millisecond):
		t.Error("not all waiters woke up after broadcast")
	}
}
