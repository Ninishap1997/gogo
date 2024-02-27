package handler

import (
	"awesomeProject/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sing-in", h.signIn)
	}

	api := router.Group("/api")
	{
		product := api.Group("/product")
		{
			product.POST("/", h.createProduct)
			product.GET("/", h.getAllProduct)
			product.GET("/:id", h.getProductId)
			product.PUT("/:id", h.updateProduct)
			product.DELETE("/:id", h.deleteProduct)
		}
	}
	return router
}
