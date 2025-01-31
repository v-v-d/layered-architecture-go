package application

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTrxManagerError(t *testing.T) {
	err := &TrxManagerError{Operation: "foo", ErrMsg: "bar"}

	assert.NotNil(t, err)

	expectedMsg := fmt.Sprintf("Error: %s - %s", err.Operation, err.ErrMsg)
	assert.Equal(t, expectedMsg, err.Error(), "Error message is incorrect")

	var trxErr *TrxManagerError
	assert.True(t, errors.As(err, &trxErr), "Error should be of type TrxManagerError")
	assert.Equal(t, "foo", trxErr.Operation)
	assert.Equal(t, "bar", trxErr.ErrMsg)
}

func TestCartNotFoundError(t *testing.T) {
	cartId := uuid.New()
	err := &CartNotFoundError{CartId: cartId}

	assert.NotNil(t, err)

	expectedMsg := fmt.Sprintf("Cart %s doesn't exist.", err.CartId)
	assert.Equal(t, expectedMsg, err.Error(), "Error message is incorrect")

	var notFoundErr *CartNotFoundError
	assert.True(t, errors.As(err, &notFoundErr), "Error should be of type CartNotFoundError")
	assert.Equal(t, cartId, notFoundErr.CartId)
}

func TestCartItemNotFoundError(t *testing.T) {
	itemId := int32(1)
	err := &CartItemNotFoundError{ItemId: itemId}

	assert.NotNil(t, err)

	expectedMsg := fmt.Sprintf("Cart item %d doesn't exist.", err.ItemId)
	assert.Equal(t, expectedMsg, err.Error(), "Error message is incorrect")

	var notFoundErr *CartItemNotFoundError
	assert.True(t, errors.As(err, &notFoundErr), "Error should be of type CartItemNotFoundError")
	assert.Equal(t, itemId, notFoundErr.ItemId)
}

func TestCustomerNotFoundError(t *testing.T) {
	AuthData := "test"
	err := &CustomerNotFoundError{AuthData: AuthData}

	assert.NotNil(t, err)

	expectedMsg := fmt.Sprintf("Customer %s doesn't exist.", err.AuthData)
	assert.Equal(t, expectedMsg, err.Error(), "Error message is incorrect")

	var notFoundErr *CustomerNotFoundError
	assert.True(t, errors.As(err, &notFoundErr), "Error should be of type CustomerNotFoundError")
	assert.Equal(t, AuthData, notFoundErr.AuthData)
}
