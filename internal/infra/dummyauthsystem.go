package infra

import (
	"layered-arch/internal/application"
	"layered-arch/internal/domain/customer"
)

type DummyAuthSystem struct {
	customerByToken map[string]customer.Customer
}

func NewDummyAuthSystem() *DummyAuthSystem {
	customerByToken := map[string]customer.Customer{
		"customer.1": customer.NewCustomer(1),
		"customer.2": customer.NewCustomer(2),
		"customer.3": customer.NewCustomer(3),
	}

	return &DummyAuthSystem{customerByToken: customerByToken}
}

func (as *DummyAuthSystem) GetCustomer(data string) (customer.Customer, error) {
	aCustomer, ok := as.customerByToken[data]

	if !ok {
		return customer.Customer{}, &application.CustomerNotFoundError{AuthData: data}
	}

	return aCustomer, nil
}
