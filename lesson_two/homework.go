package main

import (
	"fmt"
)

// SetInterface 接口
type SetInterface interface {
	// Put 往里塞数据 不做判断
	Put(key string)
	// Keys 以切片形式返回所有MAP数据
	Keys() []string
	// Contains 查询操作，看看有没有对应的Item存在
	Contains(key string) bool
	// Remove 移除操作
	Remove(key string)
	// PutIfAbsent 如果之前已经有了，就返回旧的值，absent =false；如果之前没有，就塞下去，返回 absent = true
	PutIfAbsent(key string) (old string, absent bool)
}

type Set struct {
	m map[string]bool
}

func (s Set) Put(key string) {
	f := s.Contains(key)
	if !f {
		s.m[key] = true
	}
}

func (s Set) Keys() []string {
	fmt.Println(s.m)
	ss := []string{""}
	for k, _ := range s.m {
		ss = append(ss, k)
	}
	return ss
}

func (s Set) Contains(key string) bool {
	flag := false
	if s.m[key] {
		flag = true
	}

	return flag
}

func (s Set) Remove(key string) {
	f := s.Contains(key)
	if f {
		delete(s.m, key)
	}
}
func (s Set) PutIfAbsent(key string) (old string, absent bool) {
	f := s.Contains(key)
	old = key
	absent = false
	if !f {
		s.Put(key)
		absent = true
	}
	return
}

func main() {
	var s Set
	s.m = make(map[string]bool)
	s.m["default_key"] = true
	s.m["default_key2"] = true

	s.Put("test_key")
	s.Keys()

	s.Remove("default_key2")
	s.Keys()

	o, f := s.PutIfAbsent("test_key")
	println(o, f)

	s.Keys()
}
