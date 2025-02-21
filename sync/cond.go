package sync

import (
    "sync"
)

type Cond interface {
    Broadcast()
    Signal()
    Wait()
}

type condFacade struct {
    realCond *sync.Cond
}

func (_ syncFacade) NewCond(l Locker) Cond {
    return condFacade{
        realCond: sync.NewCond(l),
    }
}

func (c condFacade) Broadcast() {
    return c.realCond.Broadcast()
}

func (c condFacade) Signal() {
    return c.realCond.Signal()
}

func (c condFacade) Wait() {
    return c.realCond.Wait()
}

