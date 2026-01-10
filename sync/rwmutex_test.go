package sync

import (
	"sync"
	"testing"
	"time"
)

func TestNewRWMutex(t *testing.T) {
	s := NewSync()
	rw := s.NewRWMutex()
	_ = rw
}

func TestRWMutex_Lock_Unlock(t *testing.T) {
	s := NewSync()
	rw := s.NewRWMutex()

	// Test basic write locking
	rw.Lock()
	// Critical section
	rw.Unlock()

	// Verify we can lock again
	rw.Lock()
	rw.Unlock()
}

func TestRWMutex_RLock_RUnlock(t *testing.T) {
	s := NewSync()
	rw := s.NewRWMutex()

	// Test basic read locking
	rw.RLock()
	// Read critical section
	rw.RUnlock()

	// Verify we can lock again
	rw.RLock()
	rw.RUnlock()
}

func TestRWMutex_TryLock(t *testing.T) {
	s := NewSync()
	rw := s.NewRWMutex()

	// Test TryLock when lock is available
	if !rw.TryLock() {
		t.Fatal("expected write lock to be available")
	}
	rw.Unlock()

	// Test TryLock when write lock is held
	rw.Lock()
	if rw.TryLock() {
		t.Error("expected TryLock to fail when write lock is held")
		rw.Unlock() // unlock the try-lock
	}
	rw.Unlock()
}

func TestRWMutex_TryRLock(t *testing.T) {
	s := NewSync()
	rw := s.NewRWMutex()

	// Test TryRLock when lock is available
	if !rw.TryRLock() {
		t.Fatal("expected read lock to be available")
	}
	rw.RUnlock()

	// Test TryRLock when write lock is held
	rw.Lock()
	if rw.TryRLock() {
		t.Error("expected TryRLock to fail when write lock is held")
		rw.RUnlock() // unlock the try-rlock
	}
	rw.Unlock()

	// Test TryRLock when another read lock is held (should succeed)
	rw.RLock()
	if !rw.TryRLock() {
		t.Error("expected TryRLock to succeed when read lock is held")
	} else {
		rw.RUnlock() // unlock the second read lock
	}
	rw.RUnlock() // unlock the first read lock
}

func TestRWMutex_MultipleReaders(t *testing.T) {
	s := NewSync()
	rw := s.NewRWMutex()

	const numReaders = 10
	readersActive := 0
	var counterMutex sync.Mutex

	done := make(chan bool)

	// Launch multiple readers
	for i := 0; i < numReaders; i++ {
		go func() {
			rw.RLock()

			counterMutex.Lock()
			readersActive++
			active := readersActive
			counterMutex.Unlock()

			// Multiple readers should be able to hold the lock simultaneously
			time.Sleep(10 * time.Millisecond)

			counterMutex.Lock()
			readersActive--
			counterMutex.Unlock()

			rw.RUnlock()

			// Report the max active readers we saw
			done <- active > 1
		}()
	}

	// Check if we had concurrent readers
	hadConcurrentReaders := false
	for i := 0; i < numReaders; i++ {
		if <-done {
			hadConcurrentReaders = true
		}
	}

	if !hadConcurrentReaders {
		t.Error("expected multiple readers to be active concurrently")
	}
}

func TestRWMutex_WriterBlocksReaders(t *testing.T) {
	s := NewSync()
	rw := s.NewRWMutex()

	writerHasLock := false
	readerGotLock := false
	var statusMutex sync.Mutex

	done := make(chan bool)

	// Writer acquires lock first
	go func() {
		rw.Lock()
		statusMutex.Lock()
		writerHasLock = true
		statusMutex.Unlock()

		time.Sleep(50 * time.Millisecond)

		rw.Unlock()
		done <- true
	}()

	// Give writer time to acquire lock
	time.Sleep(10 * time.Millisecond)

	// Reader tries to acquire lock
	go func() {
		// Verify writer has lock before we try to read
		statusMutex.Lock()
		writerHadLock := writerHasLock
		statusMutex.Unlock()

		rw.RLock()

		// If writer had lock, reader should be blocked until writer releases
		statusMutex.Lock()
		readerGotLock = writerHadLock
		statusMutex.Unlock()

		rw.RUnlock()
		done <- true
	}()

	// Wait for both
	<-done
	<-done

	statusMutex.Lock()
	gotLockAfterWriter := readerGotLock
	statusMutex.Unlock()

	if !gotLockAfterWriter {
		t.Error("expected reader to be blocked by writer")
	}
}

func TestRWMutex_ReaderBlocksWriter(t *testing.T) {
	s := NewSync()
	rw := s.NewRWMutex()

	readerHasLock := false
	writerGotLock := false
	var statusMutex sync.Mutex

	done := make(chan bool)

	// Reader acquires lock first
	go func() {
		rw.RLock()
		statusMutex.Lock()
		readerHasLock = true
		statusMutex.Unlock()

		time.Sleep(50 * time.Millisecond)

		rw.RUnlock()
		done <- true
	}()

	// Give reader time to acquire lock
	time.Sleep(10 * time.Millisecond)

	// Writer tries to acquire lock
	go func() {
		// Verify reader has lock
		statusMutex.Lock()
		readerHadLock := readerHasLock
		statusMutex.Unlock()

		rw.Lock()

		// If reader had lock, writer should be blocked until reader releases
		statusMutex.Lock()
		writerGotLock = readerHadLock
		statusMutex.Unlock()

		rw.Unlock()
		done <- true
	}()

	// Wait for both
	<-done
	<-done

	statusMutex.Lock()
	gotLockAfterReader := writerGotLock
	statusMutex.Unlock()

	if !gotLockAfterReader {
		t.Error("expected writer to be blocked by reader")
	}
}

func TestRWMutex_Concurrent(t *testing.T) {
	s := NewSync()
	rw := s.NewRWMutex()

	counter := 0
	const numReaders = 50
	const numWriters = 10
	const numOps = 50

	var wg sync.WaitGroup

	// Launch readers
	wg.Add(numReaders)
	for i := 0; i < numReaders; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < numOps; j++ {
				rw.RLock()
				_ = counter // read
				rw.RUnlock()
			}
		}()
	}

	// Launch writers
	wg.Add(numWriters)
	for i := 0; i < numWriters; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < numOps; j++ {
				rw.Lock()
				counter++ // write
				rw.Unlock()
			}
		}()
	}

	wg.Wait()

	expected := numWriters * numOps
	if counter != expected {
		t.Errorf("expected counter to be %d, got %d", expected, counter)
	}
}
