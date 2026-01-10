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
	c.realCond.Broadcast()
}

func (c condFacade) Signal() {
	c.realCond.Signal()
}

func (c condFacade) Wait() {
	c.realCond.Wait()
}
