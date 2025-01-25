package repositories

import (
	"github.com/google/uuid"
	"layered-arch/internal/application"
	"layered-arch/internal/domain/cart"
	"layered-arch/internal/domain/cartitem"
)

type InMemoryCartsRepository struct {
	storage *map[uuid.UUID]cart.Cart
}

func NewInMemoryCartsRepository(storage *map[uuid.UUID]cart.Cart) *InMemoryCartsRepository {
	return &InMemoryCartsRepository{storage: storage}
}

func (r *InMemoryCartsRepository) Create(aCart cart.Cart) (cart.Cart, error) {
	(*r.storage)[aCart.Id] = aCart
	return aCart, nil
}

func (r *InMemoryCartsRepository) Retrieve(cartId uuid.UUID) (cart.Cart, error) {
	storage := *r.storage
	aCart, exists := storage[cartId]

	if !exists {
		return aCart, &application.CartNotFoundError{CartId: cartId}
	}

	return aCart, nil
}

func (r *InMemoryCartsRepository) Update(aCart cart.Cart) (cart.Cart, error) {
	storage := *r.storage

	if _, exists := storage[aCart.Id]; !exists {
		return aCart, &application.CartNotFoundError{CartId: aCart.Id}
	}

	storage[aCart.Id] = aCart

	return aCart, nil
}

func (r *InMemoryCartsRepository) Clear(aCart cart.Cart) error {
	storage := *r.storage

	if _, exists := storage[aCart.Id]; !exists {
		return &application.CartNotFoundError{CartId: aCart.Id}
	}

	aCart.Items = []cartitem.CartItem{}
	storage[aCart.Id] = aCart

	return nil
}
