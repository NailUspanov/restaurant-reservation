package handler

import (
	"github.com/gin-gonic/gin"
	"restaurant-reservation/pkg/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	restaurants := router.Group("/restaurants", h.timeValidation)
	{
		restaurants.POST("/", h.createReservation)
		restaurants.POST("/available", h.getAvailable)
	}

	return router
}
