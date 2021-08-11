/*
Copyright 2020 RS4
@Author: Weny Xu
@Date: 2021/03/06 20:37
*/

package storage

import (
	"sync"
)

type memoryStorage struct {
	mu   sync.Mutex
	data map[interface{}]interface{}
}

func (b *memoryStorage) Set(key, value interface{}) error {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.data[key] = value
	return nil
}

func (b *memoryStorage) Get(key interface{}) (interface{}, bool) {
	b.mu.Lock()
	defer b.mu.Unlock()
	v, ok := b.data[key]
	return v, ok
}

func NewMemoryStorage() Storage {
	return &memoryStorage{
		mu:   sync.Mutex{},
		data: make(map[interface{}]interface{}),
	}
}
