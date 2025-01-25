package repositories

import (
	"github.com/google/uuid"
	"layered-arch/internal/application"
	"layered-arch/internal/domain/cart"
	"layered-arch/internal/domain/cartitem"
)

type InMemoryItemsRepository struct {
	storage *map[uuid.UUID]cart.Cart
}

func NewInMemoryItemsRepository(storage *map[uuid.UUID]cart.Cart) *InMemoryItemsRepository {
	return &InMemoryItemsRepository{storage: storage}
}

func (r *InMemoryItemsRepository) Add(item cartitem.CartItem) error {
	storage := *r.storage
	aCart, exists := storage[item.CartId]

	if !exists {
		return &application.CartNotFoundError{CartId: item.CartId}
	}

	aCart.Items = append(aCart.Items, item)
	storage[item.CartId] = aCart

	return nil
}

func (r *InMemoryItemsRepository) Update(item cartitem.CartItem) (cartitem.CartItem, error) {
	storage := *r.storage
	aCart, exists := storage[item.CartId]

	if !exists {
		return item, &application.CartNotFoundError{CartId: item.CartId}
	}

	for i := 0; i < len(aCart.Items); i++ {
		if aCart.Items[i].Id != item.Id {
			continue
		}

		aCart.Items[i] = item
		storage[item.CartId] = aCart

		return item, nil
	}

	return item, &application.CartItemNotFoundError{ItemId: item.Id}
}

func (r *InMemoryItemsRepository) Delete(item cartitem.CartItem) error {
	storage := *r.storage
	aCart, exists := storage[item.CartId]

	if !exists {
		return &application.CartNotFoundError{CartId: item.CartId}
	}

	for i := 0; i < len(aCart.Items); i++ {
		if aCart.Items[i].Id != item.Id {
			continue
		}

		aCart.Items[i] = aCart.Items[len(aCart.Items)-1]
		aCart.Items = aCart.Items[:len(aCart.Items)-1]
		storage[item.CartId] = aCart

		return nil
	}

	return &application.CartItemNotFoundError{ItemId: item.Id}
}
