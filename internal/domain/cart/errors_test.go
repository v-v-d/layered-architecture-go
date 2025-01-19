package cart

import (
	"errors"
	"fmt"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

func getRandomStatus() StatusEnum {
	statuses := []StatusEnum{StatusOpened, StatusDeactivated, StatusLocked, StatusCompleted}
	return statuses[gofakeit.Number(0, len(statuses)-1)]
}

func TestOwnershipError(t *testing.T) {
	randomCustomerId := int32(gofakeit.IntRange(1, 999))
	err := &OwnershipError{CustomerId: randomCustomerId}

	assert.NotNil(t, err)

	expectedMsg := fmt.Sprintf("Error: Cart doesn't belong to user with id %d.", randomCustomerId)
	assert.Equal(t, expectedMsg, err.Error())

	var ownershipErr *OwnershipError
	assert.True(t, errors.As(err, &ownershipErr), "Expected type: OwnershipError")
	assert.Equal(t, randomCustomerId, ownershipErr.CustomerId)
}

func TestForbiddenError(t *testing.T) {
	randomStatus := getRandomStatus()
	err := &ForbiddenError{Status: randomStatus}

	assert.NotNil(t, err)

	expectedMsg := fmt.Sprintf("Error: Cart with status %s couldn't be modified.", randomStatus)
	assert.Equal(t, expectedMsg, err.Error())

	var forbiddenErr *ForbiddenError
	assert.True(t, errors.As(err, &forbiddenErr), "Expected type: ForbiddenError")
	assert.Equal(t, randomStatus, forbiddenErr.Status)
}

func TestItemAlreadyExistsError(t *testing.T) {
	randomItemId := int32(gofakeit.IntRange(1, 999))
	err := &ItemAlreadyExistsError{ItemId: randomItemId}

	assert.NotNil(t, err)

	expectedMsg := fmt.Sprintf("Error: Item %d already in cart.", randomItemId)
	assert.Equal(t, expectedMsg, err.Error())

	var existsErr *ItemAlreadyExistsError
	assert.True(t, errors.As(err, &existsErr), "Expected type: ItemAlreadyExistsError")
	assert.Equal(t, randomItemId, existsErr.ItemId)
}

func TestChangeStatusError(t *testing.T) {
	randomStatus1 := getRandomStatus()
	randomStatus2 := getRandomStatus()
	randomCartId := uuid.New()
	err := &ChangeStatusError{
		CartID:  randomCartId,
		Current: randomStatus1,
		Attempt: randomStatus2,
	}

	assert.NotNil(t, err)

	expectedMsg := fmt.Sprintf(
		"Error: Cart %s change status failed! Current %s -> %s",
		randomCartId,
		randomStatus1,
		randomStatus2,
	)
	assert.Equal(t, expectedMsg, err.Error())

	var statusErr *ChangeStatusError
	assert.True(t, errors.As(err, &statusErr), "Expected type: ChangeStatusError")
	assert.Equal(t, randomCartId, statusErr.CartID)
	assert.Equal(t, randomStatus1, statusErr.Current)
	assert.Equal(t, randomStatus2, statusErr.Attempt)
}
