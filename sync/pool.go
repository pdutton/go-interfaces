package sync

import (
    "sync"
)

type Pool interface {
    Get() any
    Put(any)
}

type poolFacade struct {
    realPool *sync.Pool
}

func (_ Sync) NewPool(new f() any) poolFacade {
    return poolFacade{
        realPool: sync.Pool{ New: new },
    }
}

func (p poolFacade) Get() any {
    return p.realPool.Get()
}

func (p poolFacade) Put(x any) {
    p.realPool.Put(x)
}

