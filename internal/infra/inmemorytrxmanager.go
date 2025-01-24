package infra

import (
	"github.com/google/uuid"
	"layered-arch/internal/application"
	"layered-arch/internal/domain/cart"
	"layered-arch/internal/infra/repositories"
)

type InMemoryTrxManager struct {
	storage map[uuid.UUID]cart.Cart
}

func NewInMemoryTrxManager(storage map[uuid.UUID]cart.Cart) InMemoryTrxManager {
	return InMemoryTrxManager{storage: storage}
}

func (t *InMemoryTrxManager) Commit() error {
	println("Commit")
	return nil
}

func (t *InMemoryTrxManager) Rollback() error {
	println("Rollback")
	return nil
}

func (t *InMemoryTrxManager) Shutdown() error {
	println("Shutdown")
	return nil
}

func (t *InMemoryTrxManager) Carts() application.CartsRepository {
	return repositories.NewInMemoryCartsRepository(t.storage)
}

func (t *InMemoryTrxManager) Items() application.ItemsRepository {
	return repositories.NewInMemoryItemsRepository(t.storage)
}
