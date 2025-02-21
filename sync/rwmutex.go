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

func (_ Sync) NewRWMutex() rwMutexFacade {
    return rwMutexFacade{
        realRWMutex: &sync.RWMutex{},
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

