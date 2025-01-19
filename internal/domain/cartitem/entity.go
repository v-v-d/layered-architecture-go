package cartitem

import (
	"fmt"
	"github.com/google/uuid"
	"layered-arch/internal/domain"
)

type CartItem struct {
	Id       int32
	Name     string
	Price    domain.Price
	Quantity domain.Quantity
	IsWeight bool
	CartId   uuid.UUID
}

func NewCartItem(
	id int32,
	name string,
	price domain.Price,
	quantity domain.Quantity,
	isWeight bool,
	cartId uuid.UUID,
) (CartItem, error) {
	return CartItem{
		Id:       id,
		Name:     name,
		Price:    price,
		Quantity: quantity,
		IsWeight: isWeight,
		CartId:   cartId,
	}, nil
}

func (c CartItem) Cost() int32 {
	if c.IsWeight {
		return (c.Price.Value() * c.Quantity.Value()) / 1000
	}
	return c.Price.Value() * c.Quantity.Value()
}

func (c CartItem) String() string {
	return fmt.Sprintf(
		"CartItem{ID: %d, Name: %s, Price: %s, Quantity: %s, IsWeight: %t}",
		c.Id, c.Name, c.Price, c.Quantity, c.IsWeight,
	)
}
