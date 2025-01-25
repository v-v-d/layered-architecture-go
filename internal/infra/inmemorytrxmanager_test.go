package infra_test

import (
	"layered-arch/internal/domain/cartitem"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"layered-arch/internal/domain/cart"
	"layered-arch/internal/infra"
)

func TestNewInMemoryTrxManager(t *testing.T) {
	storage := make(map[uuid.UUID]cart.Cart)
	trxManager := infra.NewInMemoryTrxManager(storage)

	assert.NotNil(t, trxManager)
}

func TestInMemoryTrxManager_Commit(t *testing.T) {
	storage := make(map[uuid.UUID]cart.Cart)
	trxManager := infra.NewInMemoryTrxManager(storage)

	err := trxManager.Commit()
	assert.NoError(t, err, "Commit should not return an error")
}

func TestInMemoryTrxManager_Rollback(t *testing.T) {
	storage := make(map[uuid.UUID]cart.Cart)
	trxManager := infra.NewInMemoryTrxManager(storage)

	err := trxManager.Rollback()
	assert.NoError(t, err, "Rollback should not return an error")
}

func TestInMemoryTrxManager_Shutdown(t *testing.T) {
	storage := make(map[uuid.UUID]cart.Cart)
	trxManager := infra.NewInMemoryTrxManager(storage)

	err := trxManager.Shutdown()
	assert.NoError(t, err, "Shutdown should not return an error")
}

func TestInMemoryTrxManager_Carts(t *testing.T) {
	storage := make(map[uuid.UUID]cart.Cart)
	trxManager := infra.NewInMemoryTrxManager(storage)
	repo := trxManager.Carts()

	assert.NotNil(t, repo, "Carts repository should not be nil")
}

func TestInMemoryTrxManager_Items(t *testing.T) {
	storage := make(map[uuid.UUID]cart.Cart)
	trxManager := infra.NewInMemoryTrxManager(storage)
	repo := trxManager.Items()

	assert.NotNil(t, repo, "Items repository should not be nil")
}

func TestInMemoryTrxManager_StorageConsistency(t *testing.T) {
	storage := make(map[uuid.UUID]cart.Cart)

	trxManager := infra.NewInMemoryTrxManager(storage)
	cartsRepo := trxManager.Carts()

	// Create a cart
	cartID := uuid.New()
	testCart := cart.Cart{Id: cartID, Items: []cartitem.CartItem{}}
	_, err := cartsRepo.Create(testCart)
	assert.NoError(t, err, "Creating a cart should not return an error")

	_, exists := storage[cartID]
	assert.True(t, exists, "Cart should exist in the shared storage")

	// Update the cart through the repository
	updatedCart := cart.Cart{Id: cartID, Items: []cartitem.CartItem{{Id: 1, Name: "Test Item"}}}
	_, err = cartsRepo.Update(updatedCart)
	assert.NoError(t, err, "Updating a cart should not return an error")

	// Validate update persistence
	assert.Equal(t, updatedCart, storage[cartID], "Cart update should be reflected in shared storage")
}

func TestInMemoryTrxManager_MultipleRepositories_SameStorage(t *testing.T) {
	storage := make(map[uuid.UUID]cart.Cart)

	trxManager := infra.NewInMemoryTrxManager(storage)

	cartsRepo := trxManager.Carts()
	itemsRepo := trxManager.Items()

	// Create a cart
	cartID := uuid.New()
	testCart := cart.Cart{Id: cartID, Items: []cartitem.CartItem{}}
	_, err := cartsRepo.Create(testCart)
	assert.NoError(t, err, "Creating a cart should not return an error")

	// Create a cart item
	item := cartitem.CartItem{Id: 1, Name: "Item 1", CartId: cartID}
	err = itemsRepo.Add(item)
	assert.NoError(t, err, "Adding an item should not return an error")

	// Validate item persistence in storage
	assert.Len(t, storage[cartID].Items, 1, "Cart should have 1 item after addition")
	assert.Equal(t, item, storage[cartID].Items[0], "Item should be correctly added")
}
