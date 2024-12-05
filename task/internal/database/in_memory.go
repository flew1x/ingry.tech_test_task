package database

import "sync"

type MemoryDatabase interface {
	Get(key string) (any, bool)
	GetAll() []any
	Set(key string, value any)
	Delete(key string)
	Clear()
}

type InMemoryDatabase[K comparable, V any] struct {
	storage map[K]V
	mu      sync.RWMutex
}

func NewInMemoryDatabase[K comparable, V any]() *InMemoryDatabase[K, V] {
	return &InMemoryDatabase[K, V]{
		storage: make(map[K]V),
	}
}

func (db *InMemoryDatabase[K, V]) Get(key K) (V, bool) {
	db.mu.RLock()
	defer db.mu.RUnlock()
	value, ok := db.storage[key]

	return value, ok
}

func (db *InMemoryDatabase[K, V]) Set(key K, value V) {
	db.mu.Lock()
	defer db.mu.Unlock()
	db.storage[key] = value
}

func (db *InMemoryDatabase[K, V]) Delete(key K) {
	db.mu.Lock()
	defer db.mu.Unlock()
	delete(db.storage, key)
}

func (db *InMemoryDatabase[K, V]) Clear() {
	db.mu.Lock()
	defer db.mu.Unlock()
	db.storage = make(map[K]V)
}

func (db *InMemoryDatabase[K, V]) GetAll() []V {
	db.mu.RLock()
	defer db.mu.RUnlock()

	var all []V

	for _, value := range db.storage {
		all = append(all, value)
	}

	return all
}
