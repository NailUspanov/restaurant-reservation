package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) getAvailable(c *gin.Context) {

	type AvailableRestaurantRequest struct {
		PeopleQuantity int    `json:"people_quantity"`
		Time           string `json:"time"`
	}

	var input AvailableRestaurantRequest
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	availableRestaurants, err := h.services.Restaurant.GetAvailable(input.PeopleQuantity, input.Time)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]any{
		"restaurants": availableRestaurants,
	})
}
