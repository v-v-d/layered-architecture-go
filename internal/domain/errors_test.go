package domain

import (
	"errors"
	"fmt"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
)

func TestInvalidQuantityError(t *testing.T) {
	randomValue := int32(gofakeit.IntRange(-100, 0))
	err := &InvalidQuantityError{Value: randomValue}

	assert.NotNil(t, err)

	expectedMsg := fmt.Sprintf("Error: Quantity %d is invalid.", randomValue)
	assert.Equal(t, expectedMsg, err.Error(), "Error message is incorrect")

	var qtyErr *InvalidQuantityError
	assert.True(t, errors.As(err, &qtyErr), "Error should be of type InvalidQuantityError")
	assert.Equal(t, randomValue, qtyErr.Value)
}

func TestInvalidPriceError(t *testing.T) {
	randomValue := int32(gofakeit.IntRange(-1000, 0))
	err := &InvalidPriceError{Value: randomValue}

	assert.NotNil(t, err)

	expectedMsg := fmt.Sprintf("Error: Price %d is invalid.", randomValue)
	assert.Equal(t, expectedMsg, err.Error(), "Error message is incorrect")

	var priceErr *InvalidPriceError
	assert.True(t, errors.As(err, &priceErr), "Error should be of type InvalidPriceError")
	assert.Equal(t, randomValue, priceErr.Value)
}
