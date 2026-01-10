package sync

import (
	"sync"
)

type RWMutex interface {
	Lock()
	RLock()
	RLocker() Locker
	RUnlock()
	TryLock() bool
	TryRLock() bool
	Unlock()
}

type rwMutexFacade struct {
	realRWMutex *sync.RWMutex
}

// RWMutexOption allows you to set options on a rw mutex in the NewRWMutex constructor
type RWMutexOption func(mut *sync.RWMutex)

// Create a read locked RWMutex
func WithRLocked() RWMutexOption {
	return func(mut *sync.RWMutex) {
		mut.RLock()
	}
}

// Create a write locked RWMutex
func WithWLocked() RWMutexOption {
	return func(mut *sync.RWMutex) {
		mut.Lock()
	}
}

func (_ syncFacade) NewRWMutex(options ...RWMutexOption) rwMutexFacade {
	var mut sync.RWMutex

	for _, opt := range options {
		opt(&mut)
	}

	return rwMutexFacade{
		realRWMutex: &mut,
	}
}

func (m rwMutexFacade) Lock() {
	m.realRWMutex.Lock()
}

func (m rwMutexFacade) RLock() {
	m.realRWMutex.RLock()
}

func (m rwMutexFacade) RLocker() Locker {
	return m.realRWMutex.RLocker()
}

func (m rwMutexFacade) RUnlock() {
	m.realRWMutex.RUnlock()
}

func (m rwMutexFacade) TryLock() bool {
	return m.realRWMutex.TryLock()
}

func (m rwMutexFacade) TryRLock() bool {
	return m.realRWMutex.TryRLock()
}

func (m rwMutexFacade) Unlock() {
	m.realRWMutex.Unlock()
}
