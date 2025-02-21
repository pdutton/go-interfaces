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

func (_ Sync) NewWaitGroup() waitGroupFacade {
    return waitGroupFacade{
        realWaitGroup: &sync.WaitGroup{},
    }
}

func (wg waitGroupFacade) Add(delta int) {
    return wg.realWaitGroup.Add(delta)
}

func (wg waitGroupFacade) Done() {
    wg.realWaitGroup.Done()
}

func (wg waitGroupFacade) Wait() {
    wg.realWaitGroup.Wait()
}

