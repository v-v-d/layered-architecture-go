package infra_test

import (
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
