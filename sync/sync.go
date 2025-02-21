package sync

import (
    "sync"
)

type Locker = sync.Locker

type Sync interface {
    // Functions:
    OnceFunc(func()) func()
    OnceValue[T any](func() T) func() T
    OnceValues[T1, T2 any](func() (T1, T2)) func() (T1, T2)

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
    return OnceFunc(f)
}

func (_ syncFacade) OnceValue[T any](f func() T) func() T {
    return OnceValue(f)
}

func (_ syncFacade) OnceValues[T1, T2 any](f func() (T1, T2)) func() (T1, T2) {
    return OnceValues(f)
}

