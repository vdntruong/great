package repositories

import (
	"context"
	"sync"

	"oop/internal/domain"
)

// InMemoryProductRepository implements ProductRepository for testing
// This demonstrates the Liskov Substitution Principle: we can substitute
// the real repository with this in-memory version
type InMemoryProductRepository struct {
	products map[string]*domain.Product
	mutex    sync.RWMutex
}

func NewInMemoryProductRepository() *InMemoryProductRepository {
	return &InMemoryProductRepository{
		products: make(map[string]*domain.Product),
	}
}

func (r *InMemoryProductRepository) FindByID(ctx context.Context, id string) (*domain.Product, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	product, exists := r.products[id]
	if !exists {
		return nil, nil
	}
	return product, nil
}

// Other methods would be implemented similarly...
