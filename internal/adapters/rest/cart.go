package rest

import (
	"github.com/gin-gonic/gin"
	"layered-arch/internal/application/usecases/cart"
	"net/http"
)

type CartController struct {
	createCartUseCase cart.ICreateCartUseCase
}

func NewCartController(createCartUseCase cart.ICreateCartUseCase) *CartController {
	return &CartController{createCartUseCase: createCartUseCase}
}

func (cc *CartController) CreateCart(c *gin.Context) {
	authToken := c.GetHeader("Authorization")

	if authToken == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "missing authorization token"})
		return
	}

	aCart, err := cc.createCartUseCase.Execute(authToken)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response := NewCartResponse(aCart)

	c.JSON(http.StatusOK, response)
}
