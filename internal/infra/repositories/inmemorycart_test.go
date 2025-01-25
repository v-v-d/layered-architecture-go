package repositories_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"layered-arch/internal/application"
	"layered-arch/internal/domain/cart"
	"layered-arch/internal/domain/cartitem"
	"layered-arch/internal/infra/repositories"
)

func TestNewInMemoryCartsRepository(t *testing.T) {
	storage := make(map[uuid.UUID]cart.Cart)
	repo := repositories.NewInMemoryCartsRepository(storage)

	assert.NotNil(t, repo, "Repository should not be nil")
}

func TestInMemoryCartsRepository_Create_Success(t *testing.T) {
	storage := make(map[uuid.UUID]cart.Cart)
	repo := repositories.NewInMemoryCartsRepository(storage)

	// Create a cart
	cartID := uuid.New()
	testCart := cart.Cart{Id: cartID, Items: []cartitem.CartItem{}}

	createdCart, err := repo.Create(testCart)
	require.NoError(t, err, "Create should not return an error")

	// Validate cart was added
	assert.Equal(t, testCart, createdCart, "Created cart should match input")
	assert.Contains(t, storage, cartID, "Cart should exist in storage")
}

func TestInMemoryCartsRepository_Retrieve_Success(t *testing.T) {
	storage := make(map[uuid.UUID]cart.Cart)
	repo := repositories.NewInMemoryCartsRepository(storage)

	// Create a cart and store it
	cartID := uuid.New()
	testCart := cart.Cart{Id: cartID, Items: []cartitem.CartItem{}}
	storage[cartID] = testCart

	// Retrieve the cart
	retrievedCart, err := repo.Retrieve(cartID)
	require.NoError(t, err, "Retrieve should not return an error")

	assert.Equal(t, testCart, retrievedCart, "Retrieved cart should match stored cart")
}

func TestInMemoryCartsRepository_Retrieve_CartNotFound(t *testing.T) {
	storage := make(map[uuid.UUID]cart.Cart)
	repo := repositories.NewInMemoryCartsRepository(storage)

	// Attempt to retrieve a non-existing cart
	_, err := repo.Retrieve(uuid.New())

	assert.Error(t, err, "Expected an error when retrieving a non-existent cart")

	// Ensure it's the correct error type
	var notFoundErr *application.CartNotFoundError
	assert.ErrorAs(t, err, &notFoundErr, "Expected CartNotFoundError")
}

func TestInMemoryCartsRepository_Update_Success(t *testing.T) {
	storage := make(map[uuid.UUID]cart.Cart)
	repo := repositories.NewInMemoryCartsRepository(storage)

	// Create and store a cart
	cartID := uuid.New()
	testCart := cart.Cart{Id: cartID, Items: []cartitem.CartItem{}}
	storage[cartID] = testCart

	// Update the cart
	updatedCart := cart.Cart{Id: cartID, Items: []cartitem.CartItem{{Id: 1, Name: "Updated Item"}}}
	_, err := repo.Update(updatedCart)
	require.NoError(t, err, "Update should not return an error")

	// Validate update
	assert.Equal(t, updatedCart, storage[cartID], "Cart should be updated in storage")
}

func TestInMemoryCartsRepository_Update_CartNotFound(t *testing.T) {
	storage := make(map[uuid.UUID]cart.Cart)
	repo := repositories.NewInMemoryCartsRepository(storage)

	// Attempt to update a non-existing cart
	_, err := repo.Update(cart.Cart{Id: uuid.New()})

	assert.Error(t, err, "Expected an error when updating a non-existent cart")

	// Ensure it's the correct error type
	var notFoundErr *application.CartNotFoundError
	assert.ErrorAs(t, err, &notFoundErr, "Expected CartNotFoundError")
}

func TestInMemoryCartsRepository_Clear_Success(t *testing.T) {
	storage := make(map[uuid.UUID]cart.Cart)
	repo := repositories.NewInMemoryCartsRepository(storage)

	// Create a cart with items
	cartID := uuid.New()
	testCart := cart.Cart{Id: cartID, Items: []cartitem.CartItem{
		{Id: 1, Name: "Test Item"},
	}}
	storage[cartID] = testCart

	// Clear cart
	err := repo.Clear(testCart)
	require.NoError(t, err, "Clear should not return an error")

	// Validate that the cart is empty
	assert.Empty(t, storage[cartID].Items, "Cart should be empty after clearing")
}

func TestInMemoryCartsRepository_Clear_CartNotFound(t *testing.T) {
	storage := make(map[uuid.UUID]cart.Cart)
	repo := repositories.NewInMemoryCartsRepository(storage)

	// Attempt to clear a non-existing cart
	err := repo.Clear(cart.Cart{Id: uuid.New()})

	assert.Error(t, err, "Expected an error when clearing a non-existent cart")

	// Ensure it's the correct error type
	var notFoundErr *application.CartNotFoundError
	assert.ErrorAs(t, err, &notFoundErr, "Expected CartNotFoundError")
}
