package zutils

import (
	"sync/atomic"
)

func NewAtomic[T any](init T) *atomicCli[T] {
	a := &atomicCli[T]{
		value: atomic.Value{},
	}
	a.value.Store(init)
	return a
}

type atomicCli[T any] struct {
	value atomic.Value
}

func (a *atomicCli[T]) Set(t T) {
	a.value.Store(t)
}

func (a *atomicCli[T]) Get() T {
	v := a.value.Load()
	ret, _ := v.(T)
	return ret
}

func (a *atomicCli[T]) Swap(newV T) (oldV T) {
	v := a.value.Swap(newV)
	ret, _ := v.(T)
	return ret
}
func (a *atomicCli[T]) CAS(oldV T, newV T) (ok bool) {
	return a.value.CompareAndSwap(oldV, newV)
}
