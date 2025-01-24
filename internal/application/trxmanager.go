package application

import (
	"fmt"
	"github.com/google/uuid"
	"layered-arch/internal/domain/cart"
	"layered-arch/internal/domain/cartitem"
)

type CartsRepository interface {
	Create(aCart cart.Cart) (cart.Cart, error)
	Retrieve(cartId uuid.UUID) (cart.Cart, error)
	Update(aCart cart.Cart) (cart.Cart, error)
	Clear(aCart cart.Cart) error
}

type ItemsRepository interface {
	Add(item cartitem.CartItem) error
	Update(item cartitem.CartItem) (cartitem.CartItem, error)
	Delete(item cartitem.CartItem) error
}

type TrxManager interface {
	Commit() error
	Rollback() error
	Shutdown() error
	Carts() CartsRepository
	Items() ItemsRepository
}

type BaseTrxManager struct {
	trxManager TrxManager
}

func NewBaseTrxManager(trxManager TrxManager) BaseTrxManager {
	return BaseTrxManager{trxManager: trxManager}
}

func (t *BaseTrxManager) Run(action func() error) error {
	defer func(trxManager TrxManager) {
		if err := trxManager.Shutdown(); err != nil {
			fmt.Printf("Failed to return connection back to pool! Shutdown error: %s", err.Error())
		}
	}(t.trxManager)

	err := action()

	if err != nil {
		if rollbackErr := t.trxManager.Rollback(); rollbackErr != nil {
			return &TrxManagerError{Operation: "Rollback", ErrMsg: rollbackErr.Error()}
		}
		return err
	}

	if commitErr := t.trxManager.Commit(); commitErr != nil {
		return &TrxManagerError{Operation: "Commit", ErrMsg: commitErr.Error()}
	}

	return nil
}
