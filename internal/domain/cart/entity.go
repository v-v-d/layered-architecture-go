package cart

import (
	"fmt"
	"github.com/google/uuid"
	"layered-arch/internal/domain/cartitem"
	"layered-arch/internal/domain/customer"
	"time"
)

const weightItemDefaultQty = 1
const minCostForCheckout = 50000

var statusTransitionRuleset = map[StatusEnum]map[StatusEnum]struct{}{
	StatusOpened: {
		StatusDeactivated: {},
		StatusLocked:      {},
	},
	StatusDeactivated: {},
	StatusLocked: {
		StatusOpened:    {},
		StatusCompleted: {},
	},
	StatusCompleted: {},
}

type Cart struct {
	CreatedAt time.Time
	Id        uuid.UUID
	Customer  customer.Customer
	Status    StatusEnum
	Items     []cartitem.CartItem
}

func NewCart(aCustomer customer.Customer) (Cart, error) {
	return Cart{
		CreatedAt: time.Now(),
		Id:        uuid.New(),
		Customer:  aCustomer,
		Status:    StatusOpened,
		Items:     []cartitem.CartItem{},
	}, nil
}

func (c *Cart) ItemsQty() int32 {
	var totalQty int32

	for _, item := range c.Items {
		if item.IsWeight {
			totalQty += weightItemDefaultQty
			continue
		}

		totalQty += item.Quantity.Value()
	}

	return totalQty
}

func (c *Cart) Cost() int32 {
	var cost int32

	for _, item := range c.Items {
		cost += item.Cost()
	}

	return cost
}

func (c *Cart) CheckoutEnabled() bool {
	return c.Cost() >= minCostForCheckout
}

func (c *Cart) AddNewItem(customer customer.Customer, item cartitem.CartItem) error {
	if err := c.checkOwnership(customer); err != nil {
		return err
	}

	if err := c.checkCanBeModified(); err != nil {
		return err
	}

	if err := c.checkItemExistence(item); err != nil {
		return err
	}

	c.Items = append(c.Items, item)

	return nil
}

func (c *Cart) Deactivate() error {
	if err := c.changeStatus(StatusDeactivated); err != nil {
		return err
	}

	return nil
}

func (c *Cart) Lock() error {
	if !c.CheckoutEnabled() {
		return &ChangeStatusError{CartID: c.Id, Current: c.Status, Attempt: StatusLocked}
	}

	if err := c.changeStatus(StatusLocked); err != nil {
		return err
	}

	return nil
}

func (c *Cart) Complete() error {
	if err := c.changeStatus(StatusCompleted); err != nil {
		return err
	}

	return nil
}

func (c *Cart) Unlock() error {
	if err := c.changeStatus(StatusOpened); err != nil {
		return err
	}

	return nil
}

func (c *Cart) checkOwnership(customer customer.Customer) error {
	if c.Customer.Id != customer.Id {
		return &OwnershipError{CustomerId: customer.Id}
	}

	return nil
}

func (c *Cart) checkCanBeModified() error {
	if c.Status != StatusOpened {
		return &ForbiddenError{Status: c.Status}
	}

	return nil
}

func (c *Cart) checkItemExistence(item cartitem.CartItem) error {
	for _, existingItem := range c.Items {
		if existingItem.Id == item.Id {
			return &ItemAlreadyExistsError{ItemId: existingItem.Id}
		}
	}

	return nil
}

func (c *Cart) changeStatus(newStatus StatusEnum) error {
	if err := c.checkCanChangeStatus(newStatus); err != nil {
		return err
	}

	c.Status = newStatus

	return nil
}

func (c *Cart) checkCanChangeStatus(newStatus StatusEnum) error {
	allowedTransitions, _ := statusTransitionRuleset[c.Status]

	if _, valid := allowedTransitions[newStatus]; !valid {
		return &ChangeStatusError{CartID: c.Id, Current: c.Status, Attempt: newStatus}
	}

	return nil
}

func (c *Cart) String() string {
	return fmt.Sprintf(
		"Cart{CreatedAt: %s, Id: %s, CustomerId: %d, Status: %s}",
		c.CreatedAt,
		c.Id,
		c.Customer.Id,
		c.Status,
	)
}
