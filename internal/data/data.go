package data

import (
	"errors"
	"sync"
)

type RingStorage struct {
	mu    sync.Mutex
	rings map[string]string
}

func (r *RingStorage) Create(key string, value string) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.rings[key] = value
}

func (r *RingStorage) Update(key string, newValue string) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.rings[key] = newValue
}

func (r *RingStorage) Get(key string) (string, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	value, ok := r.rings[key]
	if !ok {
		return "", errors.New("didnt Found")
	}

	return value, nil
}

type CrudForRing interface {
	Create(key string, value string)
	Update(key string, newValue string)
	Get(key string) (string, error)
}

type RingHandler struct {
	Storage CrudForRing
}
