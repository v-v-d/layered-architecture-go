package cart

import (
	"layered-arch/internal/application"
	"layered-arch/internal/domain/cart"
)

// ICreateCartUseCase is used for testing purposes.
type ICreateCartUseCase interface {
	Execute(authData string) (cart.Cart, error)
}

type CreateCartUseCase struct {
	trxManager application.TrxManager
	authSystem application.AuthSystem
}

func NewCreateCartUseCase(
	trxManager application.TrxManager,
	authSystem application.AuthSystem,
) ICreateCartUseCase {
	return &CreateCartUseCase{trxManager: trxManager, authSystem: authSystem}
}

func (uc *CreateCartUseCase) Execute(authData string) (cart.Cart, error) {
	var newCart cart.Cart

	aCustomer, err := uc.authSystem.GetCustomer(authData)

	if err != nil {
		return newCart, err
	}

	newCart = cart.NewCart(aCustomer)

	err = uc.trxManager.Run(func() error {
		repo := uc.trxManager.Carts()

		if _, err = repo.Create(newCart); err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return newCart, err
	}

	return newCart, nil
}
