package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"net/http"
	"regexp"
	"restaurant-reservation/pkg/models"
)

func (h *Handler) getAvailable(c *gin.Context) {

	type AvailableRestaurantRequest struct {
		PeopleQuantity int    `json:"people_quantity"`
		Time           string `json:"time"`
	}

	var availableRestaurants []models.AvailableRestaurantResponse
	var input AvailableRestaurantRequest
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if isTimeValid(input.Time) {
		var err error
		availableRestaurants, err = h.services.RestaurantService.GetAvailable(input.PeopleQuantity, input.Time)
		_ = availableRestaurants
		if err != nil {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}
	}
	c.JSON(http.StatusOK, map[string]any{
		"restaurants": availableRestaurants,
	})
}

func isTimeValid(time string) bool {
	timePattern := viper.GetString("time_pattern")
	reger, _ := regexp.Compile(timePattern)
	return reger.MatchString(time)
}
