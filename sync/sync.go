package sync

import (
    "sync"
)

type Locker = sync.Locker

type Sync interface {
    // Functions:
    OnceFunc(func()) func()
    // For simplicity, these functions are not implemented, as
    // they require typed arguments:
    // OnceValue(func() T0) func() T0
    // OnceValues(func() (T1, T2)) func() (T1, T2)

    // Constructors:
    NewCond(Locker) Cond
    NewMap() Map
    NewMutex() Mutex
    NewOnce() Once
    NewPool(func() any) Pool
    NewRWMutex() RWMutex
    NewWaitGroup() WaitGroup
}

type syncFacade struct {
}

func NewSync() syncFacade {
    return syncFacade{}
}

func (_ syncFacade) OnceFunc(f func()) func() {
    return sync.OnceFunc(f)
}
