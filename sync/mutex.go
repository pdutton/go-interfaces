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

func (_ Sync) NewMutex() mutexFacade {
    return mutexFacade{
        realMutex: &sync.Mutex{},
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

