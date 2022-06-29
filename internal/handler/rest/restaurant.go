package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
	"restaurant-reservation/internal/domain/dto"
)

func (h *Handler) getAvailable(c *gin.Context) {

	type AvailableRestaurantRequest struct {
		PeopleQuantity int    `json:"people_quantity"`
		Time           string `json:"time"`
	}

	var availableRestaurants []dto.AvailableRestaurantResponse
	var input AvailableRestaurantRequest
	if err := c.ShouldBindBodyWith(&input, binding.JSON); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	availableRestaurants, err := h.services.RestaurantService.GetAvailable(input.PeopleQuantity, input.Time)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]any{
		"restaurants": availableRestaurants,
	})
}
