package infra

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"layered-arch/internal/application"
	"layered-arch/internal/domain/customer"
)

func TestNewDummyAuthSystem(t *testing.T) {
	authSystem := NewDummyAuthSystem()
	assert.NotNil(t, authSystem)
	assert.Len(t, authSystem.customerByToken, 3)
}

func TestDummyAuthSystem_GetCustomer_Success(t *testing.T) {
	authSystem := NewDummyAuthSystem()

	expectedCustomer := customer.NewCustomer(1)
	actualCustomer, err := authSystem.GetCustomer("customer.1")

	assert.NoError(t, err)
	assert.Equal(t, expectedCustomer, actualCustomer)
}

func TestDummyAuthSystem_GetCustomer_NotFound(t *testing.T) {
	authSystem := NewDummyAuthSystem()

	_, err := authSystem.GetCustomer("unknown_customer")
	assert.Error(t, err)
	assert.IsType(t, &application.CustomerNotFoundError{}, err)
}
