package ctx

import (
	"context"
	"sync"
)

type Cache interface {
	Get(key string) (string, error)
}

type OtherCache interface {
	GetValue(ctx context.Context, key string) (any, error)
}

// CacheAdapter 适配器强调的是不同接口之间进行适配
// 装饰器强调的是添加额外的功能
type CacheAdapter struct {
	Cache
}

func (c *CacheAdapter) GetValue(ctx context.Context, key string) (any, error) {
	return c.Cache.Get(key)
}

type memoryMap struct {
	m map[string]string
}

func (m *memoryMap) Get(key string) (string, error) {
	return m.m[key], nil
}

var s = &SafeCache{
	Cache: &memoryMap{},
}

// SafeCache 我要改造为线程安全的
// 无侵入式地改造
type SafeCache struct {
	Cache
	lock sync.RWMutex
}

func (s *SafeCache) Get(key string) (string, error) {
	s.lock.RLock()
	defer s.lock.RUnlock()
	return s.Cache.Get(key)
}
