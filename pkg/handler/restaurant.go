package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"net/http"
	"regexp"
	"restaurant-reservation/pkg/models"
	"strings"
	time2 "time"
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

	parsedTime := strings.Trim(time, ")")
	parsedTime = strings.Trim(parsedTime, "[")

	times := strings.Split(parsedTime, ", ")
	timeStart, err := time2.Parse("2006-01-2 15:04", times[0])
	if err != nil {
		return false
	}
	timeEnd, err := time2.Parse("2006-01-2 15:04", times[1])
	if err != nil {
		return false
	}

	return reger.MatchString(time) && timeEnd.Sub(timeStart).Hours() == 2 && isStartTimeValid(timeStart)

}

func isStartTimeValid(timeStart time2.Time) bool {
	return !timeStart.Before(time2.Now()) && timeStart.Before(time2.Now().AddDate(0, 2, 0)) &&
		timeStart.Before(time2.Date(timeStart.Year(), timeStart.Month(), timeStart.Day(), 20, 1, 0, 0, timeStart.Location())) &&
		timeStart.After(time2.Date(timeStart.Year(), timeStart.Month(), timeStart.Day(), 8, 59, 0, 0, timeStart.Location()))
}
