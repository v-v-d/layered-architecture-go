package main

import (
	"fmt"
	"log"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// CartItem represents an item that can be a part of a shopping cart. It also calculates the cost of the item based on its quantity and price.
type CartItem struct {
	ID       int
	Name     string
	Qty      decimal.Decimal
	Price    decimal.Decimal
	IsWeight bool
	CartID   uuid.UUID
}

// NewCartItem creates a new CartItem instance from given data.
func NewCartItem(id int, name string, qty, price decimal.Decimal, isWeight bool, cartID uuid.UUID) (*CartItem, error) {
	if qty.LessThanOrEqual(decimal.NewFromInt(0)) {
		log.Printf("Invalid item %d qty detected! Required > 0, got %s.", id, qty.String())
		return nil, fmt.Errorf("quantity must be greater than 0")
	}

	return &CartItem{
		ID:       id,
		Name:     name,
		Qty:      qty,
		Price:    price,
		IsWeight: isWeight,
		CartID:   cartID,
	}, nil
}

// Cost calculates the cost of the cart item based on its price and quantity.
func (c *CartItem) Cost() decimal.Decimal {
	return c.Price.Mul(c.Qty)
}

func main() {
	// Example of creating a CartItem
	cartID, _ := uuid.NewUUID()
	qty, _ := decimal.NewFromString("0")
	price, _ := decimal.NewFromString("19.99")

	item, err := NewCartItem(1, "Apple", qty, price, false, cartID)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Total cost of the item: %s\n", item.Cost().String())
}
