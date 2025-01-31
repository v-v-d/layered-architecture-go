package main

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.uber.org/fx"
	"layered-arch/internal/adapters/rest"
	"layered-arch/internal/application"
	"layered-arch/internal/application/usecases/cart"
	domaincart "layered-arch/internal/domain/cart"
	"layered-arch/internal/infra"
)

func main() {
	app := fx.New(
		fx.Provide(
			NewRouter,
			NewStorage,
			fx.Annotate(infra.NewInMemoryTrxManager, fx.As(new(application.TrxManager))),
			fx.Annotate(infra.NewDummyAuthSystem, fx.As(new(application.AuthSystem))),
			cart.NewCreateCartUseCase,
			rest.NewCartController,
		),
		fx.Invoke(RegisterRoutes),
	)

	app.Run()
}

func NewRouter() *gin.Engine {
	return gin.Default()
}

func NewStorage() map[uuid.UUID]domaincart.Cart {
	return make(map[uuid.UUID]domaincart.Cart)
}

func RegisterRoutes(router *gin.Engine, cartController *rest.CartController) {
	router.POST("/api/v1/carts", cartController.CreateCart)
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})
	router.Run()
}
