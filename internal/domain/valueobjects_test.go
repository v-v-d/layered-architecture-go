package domain

import (
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
)

func generateFakePrice() (Price, error) {
	return NewPrice(int32(gofakeit.Number(100, 100000)))
}

func generateFakeQuantity() (Quantity, error) {
	return NewQuantity(int32(gofakeit.Number(10, 5000)))
}

func TestNewPrice_Valid(t *testing.T) {
	price, err := generateFakePrice()
	assert.NoError(t, err, "There should be no error for a valid price")
	assert.Greater(t, price.Value(), int32(0), "Price value should be greater than 0")
}

func TestNewPrice_Invalid(t *testing.T) {
	_, err := NewPrice(0)
	assert.Error(t, err, "An error was expected for price 0")

	_, err = NewPrice(-int32(gofakeit.Number(1, 10000)))
	assert.Error(t, err, "An error was expected for a negative price")
}

func TestPrice_String(t *testing.T) {
	price, _ := generateFakePrice()
	assert.Contains(t, price.String(), string(price.Value()))
}

func TestNewQuantity_Valid(t *testing.T) {
	quantity, err := generateFakeQuantity()
	assert.NoError(t, err, "There should be no error for a valid quantity")
	assert.Greater(t, quantity.Value(), int32(0), "Quantity value should be greater than 0")
}

func TestNewQuantity_Invalid(t *testing.T) {
	_, err := NewQuantity(0)
	assert.Error(t, err, "An error was expected for quantity 0")

	_, err = NewQuantity(-int32(gofakeit.Number(1, 5000)))
	assert.Error(t, err, "An error was expected for a negative quantity")
}

func TestQuantity_String(t *testing.T) {
	quantity, _ := generateFakeQuantity()
	assert.Contains(t, quantity.String(), string(quantity.Value()))
}
