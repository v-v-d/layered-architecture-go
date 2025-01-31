package rest

import (
	"fmt"
	"time"

	"layered-arch/internal/domain/cart"
)

type CartItemResponse struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Price    int    `json:"price"`
	Quantity int    `json:"quantity"`
	IsWeight bool   `json:"is_weight"`
}

func (c CartItemResponse) String() string {
	return fmt.Sprintf(
		"{\"id\": %d, \"name\": \"%s\", \"price\": %d, \"quantity\": %d, \"is_weight\": %t}",
		c.Id,
		c.Name,
		c.Price,
		c.Quantity,
		c.IsWeight,
	)
}

type CartResponse struct {
	CreatedAt  time.Time          `json:"created_at"`
	Id         string             `json:"id"`
	CustomerId int                `json:"customer_id"`
	Status     string             `json:"status"`
	Items      []CartItemResponse `json:"items"`
}

func NewCartResponse(aCart cart.Cart) CartResponse {
	items := make([]CartItemResponse, 0)

	for _, item := range aCart.Items {
		items = append(
			items,
			CartItemResponse{
				Id:       int(item.Id),
				Name:     item.Name,
				Price:    int(item.Price.Value()),
				Quantity: int(item.Quantity.Value()),
				IsWeight: item.IsWeight,
			},
		)
	}

	return CartResponse{
		CreatedAt:  aCart.CreatedAt,
		Id:         aCart.Id.String(),
		CustomerId: int(aCart.Customer.Id),
		Status:     string(aCart.Status),
		Items:      items,
	}
}

func (c CartResponse) String() string {
	return fmt.Sprintf(
		"{\"created_at\": \"%s\", \"id\": \"%s\", \"customer_id\": %d, \"status\": \"%s\", \"items\": %s}",
		c.CreatedAt.Format(time.RFC3339),
		c.Id,
		c.CustomerId,
		c.Status,
		c.Items,
	)
}
