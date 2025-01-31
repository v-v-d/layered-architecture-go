package cart

import (
	"errors"
	"fmt"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"layered-arch/internal/domain"
	"layered-arch/internal/domain/cartitem"
	"layered-arch/internal/domain/customer"
	"testing"
)

type cartItemFactory func(
	itemQuantity domain.Quantity,
	itemPrice domain.Price,
) cartitem.CartItem

func generateUserId() int32 {
	return int32(gofakeit.Number(1, 10))
}

func generateCustomer() customer.Customer {
	userId := generateUserId()
	aCustomer := customer.NewCustomer(userId)

	return aCustomer
}

func generateWeightCartItem(quantity domain.Quantity, price domain.Price) cartitem.CartItem {
	item := cartitem.NewCartItem(
		gofakeit.Int32(),
		gofakeit.ProductName(),
		price,
		quantity,
		true,
		uuid.New(),
	)

	return item
}

func generatePieceCartItem(quantity domain.Quantity, price domain.Price) cartitem.CartItem {
	item := cartitem.NewCartItem(
		gofakeit.Int32(),
		gofakeit.ProductName(),
		price,
		quantity,
		false,
		uuid.New(),
	)

	return item
}

func generateCartItem() cartitem.CartItem {
	itemQuantity, _ := domain.NewQuantity(int32(gofakeit.Number(1, 10)))
	itemPrice, _ := domain.NewPrice(int32(gofakeit.Number(100, 100000)))

	return generatePieceCartItem(itemQuantity, itemPrice)
}

func generateCart(
	numberOfItems int32,
	itemQuantity domain.Quantity,
	itemPrice domain.Price,
	cartItemFactory cartItemFactory,
) Cart {
	aCustomer := generateCustomer()
	items := make([]cartitem.CartItem, numberOfItems)

	for i := 0; i < int(numberOfItems); i++ {
		items[i] = cartItemFactory(itemQuantity, itemPrice)
	}

	cart := NewCart(aCustomer)
	cart.Items = items

	return cart
}

func generateEmptyCart() Cart {
	aCustomer := generateCustomer()
	cart := NewCart(aCustomer)

	return cart
}

func TestItemsQty_PieceItemsOnly(t *testing.T) {
	numberOfItems := int32(gofakeit.Number(2, 5))
	itemQuantity, _ := domain.NewQuantity(int32(gofakeit.Number(1, 10)))
	itemPrice, _ := domain.NewPrice(int32(gofakeit.Number(100, 100000)))
	cart := generateCart(numberOfItems, itemQuantity, itemPrice, generatePieceCartItem)

	result := cart.ItemsQty()

	expected := itemQuantity.Value() * numberOfItems
	assert.Equal(t, result, expected)
}

func TestItemsQty_WeightItemsOnly(t *testing.T) {
	numberOfItems := int32(gofakeit.Number(2, 5))
	itemQuantity, _ := domain.NewQuantity(int32(gofakeit.Number(100, 99999)))
	itemPrice, _ := domain.NewPrice(int32(gofakeit.Number(100, 100000)))
	cart := generateCart(numberOfItems, itemQuantity, itemPrice, generateWeightCartItem)

	result := cart.ItemsQty()

	expected := weightItemDefaultQty * numberOfItems
	assert.Equal(t, result, expected)
}

func TestCost(t *testing.T) {
	numberOfItems := int32(gofakeit.Number(2, 5))
	itemQuantity, _ := domain.NewQuantity(int32(gofakeit.Number(1, 10)))
	itemPrice, _ := domain.NewPrice(int32(gofakeit.Number(100, 100000)))
	cart := generateCart(numberOfItems, itemQuantity, itemPrice, generatePieceCartItem)

	result := cart.Cost()

	var expected int32

	for _, item := range cart.Items {
		expected += item.Cost()
	}

	assert.Equal(t, result, expected)
}

func TestCheckoutEnabled_EmptyCart(t *testing.T) {
	cart := generateEmptyCart()
	assert.False(t, cart.CheckoutEnabled())
}

func TestCheckoutEnabled_MinCostReached(t *testing.T) {
	numberOfItems := int32(1)
	itemQuantity, _ := domain.NewQuantity(1)
	itemPrice, _ := domain.NewPrice(minCostForCheckout)
	cart := generateCart(numberOfItems, itemQuantity, itemPrice, generatePieceCartItem)

	assert.True(t, cart.CheckoutEnabled())
}

func TestCheckoutEnabled_MoreThanMinCost(t *testing.T) {
	numberOfItems := int32(1)
	itemQuantity, _ := domain.NewQuantity(1)
	itemPrice, _ := domain.NewPrice(minCostForCheckout + 1)
	cart := generateCart(numberOfItems, itemQuantity, itemPrice, generatePieceCartItem)

	assert.True(t, cart.CheckoutEnabled())
}

func TestAddNewItem_Ok(t *testing.T) {
	item := generateCartItem()
	aCustomer := generateCustomer()
	cart := NewCart(aCustomer)

	err := cart.AddNewItem(aCustomer, item)

	assert.NoError(t, err)
	assert.Len(t, cart.Items, 1)
	assert.Equal(t, cart.Items[0], item)
}

func TestAddNewItem_OwnershipError(t *testing.T) {
	item := generateCartItem()
	aCustomer := generateCustomer()
	cart := NewCart(aCustomer)
	anotherCustomer := customer.NewCustomer(int32(gofakeit.Number(-99, -1)))

	err := cart.AddNewItem(anotherCustomer, item)

	assert.NotNil(t, err)

	var ownershipErr *OwnershipError
	assert.True(t, errors.As(err, &ownershipErr))
	assert.Equal(t, anotherCustomer.Id, ownershipErr.CustomerId)

	assert.Len(t, cart.Items, 0)
}

func TestAddNewItem_ForbiddenError(t *testing.T) {
	item := generateCartItem()
	aCustomer := generateCustomer()
	cart := NewCart(aCustomer)

	_ = cart.Deactivate()

	err := cart.AddNewItem(aCustomer, item)

	assert.NotNil(t, err)

	var forbiddenErr *ForbiddenError
	assert.True(t, errors.As(err, &forbiddenErr))
	assert.Equal(t, cart.Status, forbiddenErr.Status)

	assert.Len(t, cart.Items, 0)
}

func TestAddNewItem_ItemAlreadyExistsError(t *testing.T) {
	item := generateCartItem()
	aCustomer := generateCustomer()
	cart := NewCart(aCustomer)

	_ = cart.AddNewItem(aCustomer, item)
	err := cart.AddNewItem(aCustomer, item)

	assert.NotNil(t, err)

	var existsErr *ItemAlreadyExistsError
	assert.True(t, errors.As(err, &existsErr))
	assert.Equal(t, item.Id, existsErr.ItemId)

	assert.Len(t, cart.Items, 1)
	assert.Equal(t, cart.Items[0], item)
}

func TestDeactivate_Ok(t *testing.T) {
	cart := generateEmptyCart()

	err := cart.Deactivate()

	assert.NoError(t, err)
	assert.Equal(t, cart.Status, StatusDeactivated)
}

func TestDeactivate_Failed(t *testing.T) {
	cart := generateEmptyCart()

	tests := map[string]struct {
		from StatusEnum
		to   StatusEnum
	}{
		"DEACTIVATED → DEACTIVATED": {StatusDeactivated, StatusDeactivated},
		"LOCKED → DEACTIVATED":      {StatusLocked, StatusDeactivated},
		"COMPLETED → DEACTIVATED":   {StatusCompleted, StatusDeactivated},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			cart.Status = tt.from

			err := cart.Deactivate()

			assert.NotNil(t, err)

			var statusErr *ChangeStatusError
			assert.True(t, errors.As(err, &statusErr))
			assert.Equal(t, cart.Status, statusErr.Current)
			assert.Equal(t, tt.to, statusErr.Attempt)

			assert.Equal(t, tt.from, cart.Status)
		})
	}
}

func TestLock_Ok(t *testing.T) {
	numberOfItems := int32(1)
	itemQuantity, _ := domain.NewQuantity(1)
	itemPrice, _ := domain.NewPrice(minCostForCheckout)
	cart := generateCart(numberOfItems, itemQuantity, itemPrice, generatePieceCartItem)

	err := cart.Lock()

	assert.NoError(t, err)
	assert.Equal(t, cart.Status, StatusLocked)
}

func TestLock_CheckoutDisabled(t *testing.T) {
	cart := generateEmptyCart()

	err := cart.Lock()

	var statusErr *ChangeStatusError
	assert.True(t, errors.As(err, &statusErr))
	assert.Equal(t, cart.Status, statusErr.Current)
	assert.Equal(t, StatusLocked, statusErr.Attempt)

	assert.Equal(t, cart.Status, StatusOpened)
}

func TestLock_Failed(t *testing.T) {
	numberOfItems := int32(1)
	itemQuantity, _ := domain.NewQuantity(1)
	itemPrice, _ := domain.NewPrice(minCostForCheckout)
	cart := generateCart(numberOfItems, itemQuantity, itemPrice, generatePieceCartItem)

	tests := map[string]struct {
		from StatusEnum
		to   StatusEnum
	}{
		"DEACTIVATED → LOCKED": {StatusDeactivated, StatusLocked},
		"LOCKED → LOCKED":      {StatusLocked, StatusLocked},
		"COMPLETED → LOCKED":   {StatusCompleted, StatusLocked},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			cart.Status = tt.from

			err := cart.Lock()

			assert.NotNil(t, err)

			var statusErr *ChangeStatusError
			assert.True(t, errors.As(err, &statusErr))
			assert.Equal(t, cart.Status, statusErr.Current)
			assert.Equal(t, tt.to, statusErr.Attempt)

			assert.Equal(t, tt.from, cart.Status)
		})
	}
}

func TestComplete_Ok(t *testing.T) {
	numberOfItems := int32(1)
	itemQuantity, _ := domain.NewQuantity(1)
	itemPrice, _ := domain.NewPrice(minCostForCheckout)
	cart := generateCart(numberOfItems, itemQuantity, itemPrice, generatePieceCartItem)

	_ = cart.Lock()

	err := cart.Complete()

	assert.NoError(t, err)
	assert.Equal(t, cart.Status, StatusCompleted)
}

func TestComplete_Failed(t *testing.T) {
	cart := generateEmptyCart()

	tests := map[string]struct {
		from StatusEnum
		to   StatusEnum
	}{
		"OPENED → COMPLETED":      {StatusOpened, StatusCompleted},
		"DEACTIVATED → COMPLETED": {StatusDeactivated, StatusCompleted},
		"COMPLETED → COMPLETED":   {StatusCompleted, StatusCompleted},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			cart.Status = tt.from

			err := cart.Complete()

			assert.NotNil(t, err)

			var statusErr *ChangeStatusError
			assert.True(t, errors.As(err, &statusErr))
			assert.Equal(t, cart.Status, statusErr.Current)
			assert.Equal(t, tt.to, statusErr.Attempt)

			assert.Equal(t, tt.from, cart.Status)
		})
	}
}

func TestUnlock_Ok(t *testing.T) {
	numberOfItems := int32(1)
	itemQuantity, _ := domain.NewQuantity(1)
	itemPrice, _ := domain.NewPrice(minCostForCheckout)
	cart := generateCart(numberOfItems, itemQuantity, itemPrice, generatePieceCartItem)

	_ = cart.Lock()

	err := cart.Unlock()

	assert.NoError(t, err)
	assert.Equal(t, cart.Status, StatusOpened)
}

func TestUnlock_Failed(t *testing.T) {
	cart := generateEmptyCart()

	tests := map[string]struct {
		from StatusEnum
		to   StatusEnum
	}{
		"OPENED → OPENED":      {StatusOpened, StatusOpened},
		"DEACTIVATED → OPENED": {StatusDeactivated, StatusOpened},
		"COMPLETED → OPENED":   {StatusCompleted, StatusOpened},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			cart.Status = tt.from

			err := cart.Unlock()

			assert.NotNil(t, err)

			var statusErr *ChangeStatusError
			assert.True(t, errors.As(err, &statusErr))
			assert.Equal(t, cart.Status, statusErr.Current)
			assert.Equal(t, tt.to, statusErr.Attempt)

			assert.Equal(t, tt.from, cart.Status)
		})
	}
}

func TestString(t *testing.T) {
	aCustomer := generateCustomer()
	cart := NewCart(aCustomer)

	expected := fmt.Sprintf(
		"Cart{CreatedAt: %s, Id: %s, CustomerId: %d, Status: %s, Items: %s}",
		cart.CreatedAt,
		cart.Id,
		cart.Customer.Id,
		cart.Status,
		cart.Items,
	)
	assert.Equal(t, cart.String(), expected)
}
