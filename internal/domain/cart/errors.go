package cart

import (
	"fmt"
	"github.com/google/uuid"
)

type OwnershipError struct {
	CustomerId int32
}

func (e *OwnershipError) Error() string {
	return fmt.Sprintf("Error: Cart doesn't belong to user with id %d.", e.CustomerId)
}

type ForbiddenError struct {
	Status StatusEnum
}

func (e *ForbiddenError) Error() string {
	return fmt.Sprintf("Error: Cart with status %s couldn't be modified.", e.Status)
}

type ItemAlreadyExistsError struct {
	ItemId int32
}

func (e *ItemAlreadyExistsError) Error() string {
	return fmt.Sprintf("Error: Item %d already in cart.", e.ItemId)
}

type ChangeStatusError struct {
	CartID  uuid.UUID
	Current StatusEnum
	Attempt StatusEnum
}

func (e *ChangeStatusError) Error() string {
	return fmt.Sprintf("Error: Cart %s change status failed! Current %s -> %s",
		e.CartID,
		e.Current,
		e.Attempt,
	)
}
