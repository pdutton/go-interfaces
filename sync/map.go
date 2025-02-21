package sync

import (
    "sync"
)

type Map interface {
    Clear()
    CompareAndDelete(any, any) bool
    CompareAndSwap(any, any, any) bool
    Delete(any)
    Load(any) (any, bool)
    LoadAndDelete(any) any, bool)
    LoadOrStore(any, any) (any, bool)
    Range(func(any, any) bool)
    Store(any, any)
    Swap(any, any) (any, bool)
}

type mapFacade struct {
    realMap *sync.Map
}

func (_ Sync) NewMap() Map {
    return mapFacade{
        realMap: &sync.Map{},
    }
}

func (m mapFacade) Clear() {
    m.realMap.Clear()
}

func (m mapFacade) CompareAndDelete(key, old any) bool {
    return m.realMap.CompareAndDelete(key, old)
}

func (m mapFacade) CompareAndSwap(key, old, new any) bool {
    return m.realMap.CompareAndSwap(key, old, new)
}

func (m mapFacade) Delete(key any) {
    return m.realMap.Delete(key)
}

func (m mapFacade) Load(key any) (any, bool) {
    return m.Load(key)
}

func (m mapFacade) LoadAndDelete(key any) any, bool) {
    return m.LoadAndDelete(key)
}

func (m mapFacade) LoadOrStore(key, value any) (any, bool) {
    return m.LoadOrStore(key, value)
}

func (m mapFacade) Range(f func(any, any) bool) {
    return m.Range(f)
}

func (m mapFacade) Store(key, value any) {
    return m.Store(key, value)
}

func (m mapFacade) Swap(key, value any) (any, bool) {
    return m.Swap(key, value)
}

