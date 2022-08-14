package sync

import "sync"

type SaveMap[K comparable, V any] struct {
	m     map[K]V
	mutex sync.RWMutex
}

// LoadOrStore loaded 代表是返回老的对象，还是返回了新的对象
func (s *SaveMap[K, V]) LoadOrStore(key K, newVal V) (val V, loaded bool) {
	s.mutex.RLock()
	val, ok := s.m[key]
	s.mutex.RUnlock()
	if ok {
		return val, true
	}
	// double check
	s.mutex.Lock()
	defer s.mutex.Unlock()
	val, ok = s.m[key]
	if ok {
		return val, true
	}
	s.m[key] = newVal
	return newVal, false
}

type valProvider[V any] func() V

func (s *SaveMap[K, V]) LoadOrStoreHeavy(key K, p valProvider[V]) (val V, loaded bool) {
	s.mutex.RLock()
	val, ok := s.m[key]
	s.mutex.RUnlock()
	if ok {
		return val, true
	}

	// double check
	s.mutex.Lock()
	defer s.mutex.Unlock()
	val, ok = s.m[key]
	if ok {
		return val, true
	}

	val = p()
	s.m[key] = val
	return val, false
}
