package rest_test

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"layered-arch/internal/domain"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"layered-arch/internal/adapters/rest"
	"layered-arch/internal/domain/cart"
	"layered-arch/internal/domain/cartitem"
	"layered-arch/internal/domain/customer"
	"layered-arch/internal/mocks"
)

func generateCart() cart.Cart {
	aCustomer := customer.NewCustomer(1)
	aCart := cart.NewCart(aCustomer)
	aPrice, _ := domain.NewPrice(1)
	aQuantity, _ := domain.NewQuantity(1)
	anItem := cartitem.NewCartItem(1, "test", aPrice, aQuantity, false, aCart.Id)
	_ = aCart.AddNewItem(aCustomer, anItem)

	return aCart
}

func TestCreateCart_Unauthorized(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	mockUseCase := mocks.NewICreateCartUseCase(t)
	controller := rest.NewCartController(mockUseCase)
	r.POST("/api/v1/carts", controller.CreateCart)

	req, _ := http.NewRequest(http.MethodPost, "/api/v1/carts", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnauthorized, w.Code)
	assert.JSONEq(t, `{"error": "missing authorization token"}`, w.Body.String())
}

func TestCreateCart_BadRequest(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	mockUseCase := mocks.NewICreateCartUseCase(t)
	controller := rest.NewCartController(mockUseCase)
	r.POST("/api/v1/carts", controller.CreateCart)

	mockUseCase.On("Execute", "valid_token").Return(cart.Cart{}, assert.AnError).Once()

	req, _ := http.NewRequest(http.MethodPost, "/api/v1/carts", nil)
	req.Header.Set("Authorization", "valid_token")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.JSONEq(t, `{"error": "assert.AnError general error for testing"}`, w.Body.String())
}

func TestCreateCart_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	mockUseCase := mocks.NewICreateCartUseCase(t)
	controller := rest.NewCartController(mockUseCase)
	r.POST("/api/v1/carts", controller.CreateCart)

	aCart := generateCart()

	mockUseCase.On("Execute", "valid_token").Return(aCart, nil)

	req, _ := http.NewRequest(http.MethodPost, "/api/v1/carts", nil)
	req.Header.Set("Authorization", "valid_token")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)

	assert.NoError(t, err)

	createdAt, ok := response["created_at"].(string)
	assert.True(t, ok, "CreatedAt should be a string")
	parsedTime, err := time.Parse(time.RFC3339, createdAt)
	assert.NoError(t, err, "CreatedAt should be in RFC3339 format")
	assert.Equal(t, aCart.CreatedAt.Truncate(time.Second), parsedTime.Truncate(time.Second), "CreatedAt should match the expected value")

	assert.Equal(t, aCart.Id.String(), response["id"])
	assert.Equal(t, float64(aCart.Customer.Id), response["customer_id"])
	assert.Equal(t, string(aCart.Status), response["status"])

	items, ok := response["items"].([]interface{})
	assert.True(t, ok, "Items should be an array")
	assert.Len(t, items, 1, "Items should contain exactly one element")

	itemMap, valid := items[0].(map[string]interface{})
	assert.True(t, valid, "Item should be a JSON object")

	assert.Equal(t, float64(aCart.Items[0].Id), itemMap["id"])
	assert.Equal(t, aCart.Items[0].Name, itemMap["name"])
	assert.Equal(t, float64(aCart.Items[0].Price.Value()), itemMap["price"])
	assert.Equal(t, float64(aCart.Items[0].Quantity.Value()), itemMap["quantity"])
	assert.Equal(t, aCart.Items[0].IsWeight, itemMap["is_weight"])
}
