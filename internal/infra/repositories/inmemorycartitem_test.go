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

func TestNewInMemoryItemsRepository(t *testing.T) {
	storage := make(map[uuid.UUID]cart.Cart)
	repo := repositories.NewInMemoryItemsRepository(storage)

	assert.NotNil(t, repo, "Repository should not be nil")
}

func TestInMemoryItemsRepository_Add_Success(t *testing.T) {
	storage := make(map[uuid.UUID]cart.Cart)
	repo := repositories.NewInMemoryItemsRepository(storage)

	// Create a cart
	cartID := uuid.New()
	testCart := cart.Cart{Id: cartID, Items: []cartitem.CartItem{}}
	storage[cartID] = testCart

	// Create a cart item
	item := cartitem.CartItem{Id: 1, Name: "Test Item", CartId: cartID}

	// Add item
	err := repo.Add(item)
	require.NoError(t, err, "Add should not return an error")

	// Validate item was added
	updatedCart := storage[cartID]
	assert.Len(t, updatedCart.Items, 1, "Cart should have 1 item")
	assert.Equal(t, item, updatedCart.Items[0], "Item should match the added item")
}

func TestInMemoryItemsRepository_Add_CartNotFound(t *testing.T) {
	storage := make(map[uuid.UUID]cart.Cart)
	repo := repositories.NewInMemoryItemsRepository(storage)

	// Create an item for a non-existing cart
	item := cartitem.CartItem{Id: 1, Name: "Test Item", CartId: uuid.New()}

	// Try to add
	err := repo.Add(item)
	assert.Error(t, err, "Expected an error when adding to a non-existent cart")

	// Ensure it's the correct error type
	var notFoundErr *application.CartNotFoundError
	assert.ErrorAs(t, err, &notFoundErr, "Expected CartNotFoundError")
}

func TestInMemoryItemsRepository_Update_Success(t *testing.T) {
	storage := make(map[uuid.UUID]cart.Cart)
	repo := repositories.NewInMemoryItemsRepository(storage)

	// Create a cart
	cartID := uuid.New()
	item := cartitem.CartItem{Id: 1, Name: "Old Item", CartId: cartID}
	testCart := cart.Cart{Id: cartID, Items: []cartitem.CartItem{item}}
	storage[cartID] = testCart

	// Update item
	updatedItem := cartitem.CartItem{Id: 1, Name: "Updated Item", CartId: cartID}
	_, err := repo.Update(updatedItem)
	require.NoError(t, err, "Update should not return an error")

	// Validate update
	updatedCart := storage[cartID]
	assert.Equal(t, "Updated Item", updatedCart.Items[0].Name, "Item should be updated")
}

func TestInMemoryItemsRepository_Update_CartNotFound(t *testing.T) {
	storage := make(map[uuid.UUID]cart.Cart)
	repo := repositories.NewInMemoryItemsRepository(storage)

	// Try to update an item in a non-existing cart
	item := cartitem.CartItem{Id: 1, Name: "Test Item", CartId: uuid.New()}
	_, err := repo.Update(item)

	assert.Error(t, err, "Expected an error when updating a non-existent cart")

	// Ensure it's the correct error type
	var notFoundErr *application.CartNotFoundError
	assert.ErrorAs(t, err, &notFoundErr, "Expected CartNotFoundError")
}

func TestInMemoryItemsRepository_Update_CartItemNotFound(t *testing.T) {
	storage := make(map[uuid.UUID]cart.Cart)
	repo := repositories.NewInMemoryItemsRepository(storage)

	// Create a cart
	cartID := uuid.New()
	item := cartitem.CartItem{Id: 1, Name: "Old Item", CartId: cartID}
	testCart := cart.Cart{Id: cartID, Items: []cartitem.CartItem{item}}
	storage[cartID] = testCart

	// Try to update a non-existing item
	nonExistingItem := cartitem.CartItem{Id: 2, Name: "Test Item", CartId: cartID}
	_, err := repo.Update(nonExistingItem)

	assert.Error(t, err, "Expected an error when updating a non-existent cart item")

	// Ensure it's the correct error type
	var notFoundErr *application.CartItemNotFoundError
	assert.ErrorAs(t, err, &notFoundErr, "Expected CartItemNotFoundError")
}

func TestInMemoryItemsRepository_Delete_Success(t *testing.T) {
	storage := make(map[uuid.UUID]cart.Cart)
	repo := repositories.NewInMemoryItemsRepository(storage)

	// Create a cart
	cartID := uuid.New()
	item := cartitem.CartItem{Id: 1, Name: "Test Item", CartId: cartID}
	testCart := cart.Cart{Id: cartID, Items: []cartitem.CartItem{item}}
	storage[cartID] = testCart

	// Delete item
	err := repo.Delete(item)
	require.NoError(t, err, "Delete should not return an error")

	// Validate item was removed
	updatedCart := storage[cartID]
	assert.Empty(t, updatedCart.Items, "Cart should have no items after deletion")
}

func TestInMemoryItemsRepository_Delete_CartNotFound(t *testing.T) {
	storage := make(map[uuid.UUID]cart.Cart)
	repo := repositories.NewInMemoryItemsRepository(storage)

	item := cartitem.CartItem{Id: 1, Name: "Test Item", CartId: uuid.New()}
	err := repo.Delete(item)

	assert.Error(t, err, "Expected an error when deleting from a non-existent cart")

	// Ensure it's the correct error type
	var notFoundErr *application.CartNotFoundError
	assert.ErrorAs(t, err, &notFoundErr, "Expected CartNotFoundError")
}

func TestInMemoryItemsRepository_Delete_CartItemNotFound(t *testing.T) {
	storage := make(map[uuid.UUID]cart.Cart)
	repo := repositories.NewInMemoryItemsRepository(storage)

	// Create a cart
	cartID := uuid.New()
	item := cartitem.CartItem{Id: 1, Name: "Old Item", CartId: cartID}
	testCart := cart.Cart{Id: cartID, Items: []cartitem.CartItem{item}}
	storage[cartID] = testCart

	// Try to update a non-existing item
	nonExistingItem := cartitem.CartItem{Id: 2, Name: "Test Item", CartId: cartID}
	err := repo.Delete(nonExistingItem)

	assert.Error(t, err, "Expected an error when deleting from a non-existent cart")

	// Ensure it's the correct error type
	var notFoundErr *application.CartItemNotFoundError
	assert.ErrorAs(t, err, &notFoundErr, "Expected CartItemNotFoundError")
}
