package sync

import (
	"sync"
)

type Mutex interface {
	Lock()
	TryLock() bool
	Unlock()
}

type mutexFacade struct {
	realMutex *sync.Mutex
}

// MutexOption allows you to set options on a mutex in the NewMutex constructor
type MutexOption func(mut *sync.Mutex)

// Create a locked Mutex
func WithLocked() MutexOption {
	return func(mut *sync.Mutex) {
		mut.Lock()
	}
}

func (_ syncFacade) NewMutex(options ...MutexOption) mutexFacade {
	var mut sync.Mutex

	for _, opt := range options {
		opt(&mut)
	}

	return mutexFacade{
		realMutex: &mut,
	}
}

func (m mutexFacade) Lock() {
	m.realMutex.Lock()
}

func (m mutexFacade) TryLock() bool {
	return m.realMutex.TryLock()
}

func (m mutexFacade) Unlock() {
	m.realMutex.Unlock()
}
