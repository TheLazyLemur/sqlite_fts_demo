package main

import (
	"errors"
	"sync"
)

type SearchCache[T any] struct {
	lock  sync.Mutex
	cache map[string]T
}

func NewSearchCache[T any]() *SearchCache[T] {
	return &SearchCache[T]{
		lock:  sync.Mutex{},
		cache: make(map[string]T),
	}
}

func (sc *SearchCache[T]) Set(key string, v T) error {
	sc.lock.Lock()
	defer sc.lock.Unlock()

	if _, ok := sc.cache[key]; ok {
		return errors.New("key already exists")
	}

	sc.cache[key] = v
	return nil
}

func (sc *SearchCache[T]) Get(key string) (*T, error) {
	sc.lock.Lock()
	defer sc.lock.Unlock()

	if v, ok := sc.cache[key]; !ok {
		return nil, errors.New("key does not exist")
	} else {
		return &v, nil
	}
}
