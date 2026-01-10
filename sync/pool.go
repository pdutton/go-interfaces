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

// PoolOption allows you to set options on a pool in the NewPool constructor
type PoolOption func(pool *sync.Pool)

// Create a pool with a given New function
func WithNew(f func() any) PoolOption {
	return func(pool *sync.Pool) {
		pool.New = f
	}
}

func (_ syncFacade) NewPool(options ...PoolOption) poolFacade {
	var pool sync.Pool

	for _, opt := range options {
		opt(&pool)
	}

	return poolFacade{
		realPool: &pool,
	}
}

func (p poolFacade) Get() any {
	return p.realPool.Get()
}

func (p poolFacade) Put(x any) {
	p.realPool.Put(x)
}
