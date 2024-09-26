package handler

import (
	"github.com/biyoba1/redisProject/controllers"
	"github.com/biyoba1/redisProject/internal/services"
	"github.com/biyoba1/redisProject/middleware"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *services.Service
}

func NewHandler(services *services.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.Default()

	//=== products ===

	products := router.Group("/products")
	{
		products.POST("/", h.CreateProduct)
		products.GET("/", h.GetAllProducts)
		products.GET("/:name", h.GetByNameProduct)
		products.PUT("/:name", h.UpdateProduct)
		products.DELETE("/:name", h.DeleteProduct)
	}

	//=== users && orders ===

	users := router.Group("/users")
	{
		auth := &controllers.Auth{}
		users.POST("/signin", auth.SignUp)
		users.POST("/login", auth.Login)
		users.GET("/validate", middleware.RequireAuth, auth.Validate)

		orders := users.Group("/orders", middleware.RequireAuth)
		{
			orders.POST("/", h.CreateOrder)
			orders.GET("/", h.GetAllOrders)
			orders.GET("/:id", h.GetOrderByID)
			orders.PUT("/:id", h.UpdateOrder)
			orders.DELETE("/:id", h.DeleteOrder)
		}
	}

	return router
}
