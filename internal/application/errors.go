package application

import (
	"fmt"
	"github.com/google/uuid"
)

type TrxManagerError struct {
	Operation string
	ErrMsg    string
}

func (e *TrxManagerError) Error() string {
	return fmt.Sprintf("Error: %s - %s", e.Operation, e.ErrMsg)
}

type CartNotFoundError struct {
	CartId uuid.UUID
}

func (e *CartNotFoundError) Error() string {
	return fmt.Sprintf("Cart %s doesn't exist.", e.CartId)
}

type CartItemNotFoundError struct {
	ItemId int32
}

func (e *CartItemNotFoundError) Error() string {
	return fmt.Sprintf("Cart item %d doesn't exist.", e.ItemId)
}
