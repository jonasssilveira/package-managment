package database

import (
	"order-package/internal/domain/optimalpackage/entity"
	"sync"
)

type Collection interface {
	Find() []*entity.PackDocument
	CreateMany(docs []entity.PackDocument)
	Delete(size int64)
}

type InMemoryPackRepository struct {
	store map[int64]int64
	mu    sync.RWMutex
}

func NewInMemoryPackRepository() *InMemoryPackRepository {
	return &InMemoryPackRepository{
		store: map[int64]int64{23: 23, 31: 31, 53: 53},
	}
}

func (r *InMemoryPackRepository) Find() []*entity.PackDocument {
	r.mu.RLock()
	defer r.mu.RUnlock()

	var packs []*entity.PackDocument
	for size := range r.store {
		packs = append(packs, &entity.PackDocument{Size: size})
	}
	return packs
}

func (r *InMemoryPackRepository) CreateMany(docs []entity.PackDocument) {
	r.mu.Lock()
	defer r.mu.Unlock()

	for _, doc := range docs {
		r.store[doc.Size] = doc.Size
	}
}

func (r *InMemoryPackRepository) Delete(size int64) {
	r.mu.Lock()
	defer r.mu.Unlock()

	delete(r.store, size)
}
