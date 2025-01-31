package rest_test

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"layered-arch/internal/adapters/rest"
	"layered-arch/internal/domain"
	"layered-arch/internal/domain/cart"
	"layered-arch/internal/domain/cartitem"
	"layered-arch/internal/domain/customer"
)

func TestCartItemResponse_String(t *testing.T) {
	item := rest.CartItemResponse{
		Id:       1,
		Name:     "Test Item",
		Price:    100,
		Quantity: 2,
		IsWeight: false,
	}
	expected := `{"id": 1, "name": "Test Item", "price": 100, "quantity": 2, "is_weight": false}`
	assert.Equal(t, expected, item.String())
}

func TestNewCartResponse(t *testing.T) {
	aCustomer := customer.NewCustomer(1)
	aCart := cart.NewCart(aCustomer)
	aPrice, _ := domain.NewPrice(100)
	aQty, _ := domain.NewQuantity(2)
	anItem := cartitem.NewCartItem(1, "Test Item", aPrice, aQty, false, aCart.Id)
	_ = aCart.AddNewItem(aCustomer, anItem)

	cartResponse := rest.NewCartResponse(aCart)

	assert.Equal(t, aCart.CreatedAt, cartResponse.CreatedAt)
	assert.Equal(t, aCart.Id.String(), cartResponse.Id)
	assert.Equal(t, int(aCart.Customer.Id), cartResponse.CustomerId)
	assert.Equal(t, string(aCart.Status), cartResponse.Status)
	assert.Len(t, cartResponse.Items, 1)
	assert.Equal(t, "Test Item", cartResponse.Items[0].Name)
}

func TestCartResponse_String(t *testing.T) {
	mockCreatedAt := time.Now().UTC().Truncate(time.Second)
	cartResponse := rest.CartResponse{
		CreatedAt:  mockCreatedAt,
		Id:         "1234",
		CustomerId: 1,
		Status:     "OPENED",
		Items:      []rest.CartItemResponse{{Id: 1, Name: "Test Item", Price: 100, Quantity: 2, IsWeight: false}},
	}

	expectedJSON, _ := json.Marshal(cartResponse)
	assert.JSONEq(t, string(expectedJSON), cartResponse.String())
}
