package cartitem

import (
	"fmt"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"layered-arch/internal/domain"
)

func generateFakePrice() (domain.Price, error) {
	return domain.NewPrice(int32(gofakeit.Number(100, 100000)))
}

func generateFakeWeightQuantity() (domain.Quantity, error) {
	return domain.NewQuantity(int32(gofakeit.Number(100, 5000)))
}

func generateFakePieceQuantity() (domain.Quantity, error) {
	return domain.NewQuantity(int32(gofakeit.Number(1, 10)))
}

func generateFakeWeightCartItem() (CartItem, error) {
	price, err := generateFakePrice()
	if err != nil {
		return CartItem{}, err
	}

	quantity, err := generateFakeWeightQuantity()
	if err != nil {
		return CartItem{}, err
	}

	return NewCartItem(
		gofakeit.Int32(),
		gofakeit.ProductName(),
		price,
		quantity,
		true, // IsWeight = true
		uuid.New(),
	), nil
}

func generateFakePieceCartItem() (CartItem, error) {
	price, err := generateFakePrice()
	if err != nil {
		return CartItem{}, err
	}

	quantity, err := generateFakePieceQuantity()
	if err != nil {
		return CartItem{}, err
	}

	return NewCartItem(
		gofakeit.Int32(),
		gofakeit.ProductName(),
		price,
		quantity,
		false, // IsWeight = false
		uuid.New(),
	), nil
}

func TestNewCartItem(t *testing.T) {
	item, err := generateFakePieceCartItem()
	assert.NoError(t, err, "Error generating CartItem")
	assert.NotNil(t, item, "CartItem should not be nil")
}

func TestCartItem_Cost_WeightProduct(t *testing.T) {
	item, err := generateFakeWeightCartItem()
	assert.NoError(t, err, "Error generating CartItem")

	expectedCost := (item.Price.Value() * item.Quantity.Value()) / 1000
	assert.Equal(t, expectedCost, item.Cost(), "Cost calculation is incorrect")
}

func TestCartItem_Cost_PieceProduct(t *testing.T) {
	item, err := generateFakePieceCartItem()
	assert.NoError(t, err, "Error generating CartItem")

	expectedCost := item.Price.Value() * item.Quantity.Value()
	assert.Equal(t, expectedCost, item.Cost(), "Cost calculation is incorrect")
}

func TestCartItem_String(t *testing.T) {
	price, _ := domain.NewPrice(1500)
	quantity, _ := domain.NewQuantity(3)
	cartId := uuid.New()

	item := NewCartItem(1, "Test Product", price, quantity, false, cartId)

	expected := fmt.Sprintf(
		"CartItem{ID: %d, Name: %s, Price: %s, Quantity: %s, IsWeight: %t}",
		item.Id, item.Name, item.Price, item.Quantity, item.IsWeight,
	)
	assert.Equal(t, expected, item.String())
}

func TestCartItem_UUID(t *testing.T) {
	item, err := generateFakePieceCartItem()
	assert.NoError(t, err, "Error generating CartItem")

	assert.NotEqual(t, uuid.Nil, item.CartId, "UUID should not be nil")
}
