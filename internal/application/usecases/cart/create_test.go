package cart

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"layered-arch/internal/domain/cart"
	"layered-arch/internal/domain/customer"
	"layered-arch/internal/mocks"
)

func TestCreateCartUseCase_Execute(t *testing.T) {
	mockTrxManager := mocks.NewTrxManager(t)
	mockAuthSystem := mocks.NewAuthSystem(t)
	mockCartsRepo := mocks.NewCartsRepository(t)

	uc := NewCreateCartUseCase(mockTrxManager, mockAuthSystem)

	aCustomer := customer.Customer{Id: 1}
	mockAuthSystem.On("GetCustomer", "valid_token").Return(aCustomer, nil).Once()

	mockTrxManager.On("Run", mock.Anything).Return(nil).Run(func(args mock.Arguments) {
		txFunc := args.Get(0).(func() error)
		err := txFunc()
		if err != nil {
			return
		}
	}).Once()

	mockTrxManager.On("Carts").Return(mockCartsRepo).Once()
	mockCartsRepo.On("Create", mock.Anything).Return(func(aCart cart.Cart) cart.Cart {
		return aCart
	}, nil).Once()

	result, err := uc.Execute("valid_token")

	assert.NoError(t, err)
	assert.NotZero(t, result.Id)
	assert.Equal(t, aCustomer.Id, result.Customer.Id)
}

func TestCreateCartUseCase_Execute_AuthFailure(t *testing.T) {
	mockTrxManager := mocks.NewTrxManager(t)
	mockAuthSystem := mocks.NewAuthSystem(t)

	uc := NewCreateCartUseCase(mockTrxManager, mockAuthSystem)

	mockAuthSystem.On("GetCustomer", "invalid_token").Return(customer.Customer{}, errors.New("authentication failed")).Once()

	_, err := uc.Execute("invalid_token")

	assert.Error(t, err)
	assert.Equal(t, "authentication failed", err.Error())
}

func TestCreateCartUseCase_Execute_CreateFailure(t *testing.T) {
	mockTrxManager := mocks.NewTrxManager(t)
	mockAuthSystem := mocks.NewAuthSystem(t)
	mockCartsRepo := mocks.NewCartsRepository(t)

	uc := NewCreateCartUseCase(mockTrxManager, mockAuthSystem)

	aCustomer := customer.Customer{Id: 1}
	mockAuthSystem.On("GetCustomer", "valid_token").Return(aCustomer, nil).Once()

	mockTrxManager.On("Run", mock.Anything).Return(errors.New("failed to create cart")).Run(func(args mock.Arguments) {
		txFunc := args.Get(0).(func() error)
		err := txFunc()
		args[0] = err
	}).Once()

	mockTrxManager.On("Carts").Return(mockCartsRepo).Once()
	mockCartsRepo.On("Create", mock.Anything).Return(cart.Cart{}, errors.New("failed to create cart")).Once()

	_, err := uc.Execute("valid_token")

	assert.Error(t, err)
	assert.Equal(t, "failed to create cart", err.Error())
}
