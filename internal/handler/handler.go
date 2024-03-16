package handler

import (
	"validate/internal/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	services service.Service
}

func NewHandler(service service.Service) *Handler {
	return &Handler{services: service}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	api := router.Group("/api")
	{
		card := api.Group("/cards")
		{
			card.POST("/", h.createCard)
			card.GET("/", h.getAllCards)
			card.GET("/:id", h.getCardById)
			card.DELETE("/:id", h.deleteCard)
			card.PUT("/:id", h.updateCard)
		}
	}

	return router
}
