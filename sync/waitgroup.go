package sync

import (
	"sync"
)

type WaitGroup interface {
	Add(int)
	Done()
	Wait()
}

type waitGroupFacade struct {
	realWaitGroup *sync.WaitGroup
}

// waitGroupOption allows you to set options on a wait group in the NewWaitGroup constructor
type WaitGroupOption func(wg *sync.WaitGroup)

// Create a wait group with the given initial count
func WithCount(nbr uint8) WaitGroupOption {
	return func(wg *sync.WaitGroup) {
		wg.Add(int(nbr))
	}
}

func (_ syncFacade) NewWaitGroup(options ...WaitGroupOption) waitGroupFacade {
	var wg sync.WaitGroup

	for _, opt := range options {
		opt(&wg)
	}

	return waitGroupFacade{
		realWaitGroup: &wg,
	}
}

func (wg waitGroupFacade) Add(delta int) {
	wg.realWaitGroup.Add(delta)
}

func (wg waitGroupFacade) Done() {
	wg.realWaitGroup.Done()
}

func (wg waitGroupFacade) Wait() {
	wg.realWaitGroup.Wait()
}
