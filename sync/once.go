package sync

import (
	"sync"
)

type Once interface {
	Do(func())
}

type onceFacade struct {
	realOnce *sync.Once
}

func (_ syncFacade) NewOnce() onceFacade {
	return onceFacade{
		realOnce: &sync.Once{},
	}
}

func (o onceFacade) Do(f func()) {
	o.realOnce.Do(f)
}
