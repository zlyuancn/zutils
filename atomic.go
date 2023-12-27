package zutils

import (
	"sync/atomic"
)

func NewAtomic[T any](init T) *AtomicValue[T] {
	a := &AtomicValue[T]{
		value: atomic.Value{},
	}
	a.value.Store(init)
	return a
}

type AtomicValue[T any] struct {
	value atomic.Value
}

func (a *AtomicValue[T]) Set(t T) {
	a.value.Store(t)
}

func (a *AtomicValue[T]) Get() T {
	v := a.value.Load()
	ret, _ := v.(T)
	return ret
}

func (a *AtomicValue[T]) Swap(newV T) (oldV T) {
	v := a.value.Swap(newV)
	ret, _ := v.(T)
	return ret
}
func (a *AtomicValue[T]) CAS(oldV T, newV T) (ok bool) {
	return a.value.CompareAndSwap(oldV, newV)
}
