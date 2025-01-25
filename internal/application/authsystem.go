package application

import "layered-arch/internal/domain/customer"

type AuthSystem interface {
	GetCustomer(data string) (customer.Customer, error)
}
