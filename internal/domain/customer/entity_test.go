package customer

import (
	"github.com/brianvoe/gofakeit/v6"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewCustomer(t *testing.T) {
	customerId := int32(gofakeit.Number(1, 999))
	aCustomer := NewCustomer(customerId)
	assert.Equal(t, customerId, aCustomer.Id)
}
