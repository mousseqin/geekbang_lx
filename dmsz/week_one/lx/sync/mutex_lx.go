package sync

import "sync"

type SafeResource struct {
	resource interface{}
	lock     sync.Mutex
}

func (s *SafeResource) DoSomethingToResource() {
	s.lock.Lock()
	defer s.lock.Unlock()
}

type List[T any] interface {
	Get(index int) T
	Set(index int, t T)
	Delete(index int) T
	Append(t T)
}

type ArrayList[T any] struct {
	val []T
}

func NewArrayList[T any](initCap int) *ArrayList[T] {
	return &ArrayList[T]{val: make([]T, 0, initCap)}
}

func (a *ArrayList[T]) Get(index int) T {
	return a.val[index]
}

func (a *ArrayList[T]) Set(index int, t T) {
	if index >= len(a.val) || index < 0 {
		panic("index 超出范围")
	}
	a.val[index] = t
}

func (a *ArrayList[T]) Delete(index int) T {
	if index >= len(a.val) || index < 0 {
		panic("index 超出范围")
	}
	res := a.val[index]
	a.val = append(a.val[:index], a.val[index+1:]...)
	return res
}

func (a *ArrayList[T]) Append(t T) {
	a.val = append(a.val, t)
}

// ----------------- SafeListDecorator -------------------------

type SafeListDecorator[T any] struct {
	l     List[T]
	mutex sync.RWMutex
}

func (s *SafeListDecorator[T]) Get(index int) T {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	return s.l.Get(index)
}

func (s *SafeListDecorator[T]) Set(index int, t T) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.l.Set(index, t)
}

func (s *SafeListDecorator[T]) Delete(index int) T {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	return s.l.Delete(index)
}

func (s *SafeListDecorator[T]) Append(t T) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.l.Append(t)
}
