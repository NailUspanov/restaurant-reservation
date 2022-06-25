package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/spf13/viper"
	"net/http"
	"regexp"
	"strings"
	time2 "time"
)

func (h *Handler) timeValidation(c *gin.Context) {
	type RequestType struct {
		Time string `json:"time"`
	}
	var input RequestType

	_ = c.ShouldBindBodyWith(&input, binding.JSON)
	if !isTimeValid(input.Time) {
		newErrorResponse(c, http.StatusBadRequest, "invalid time input")
		return
	}

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
	now := time2.Now()
	time, err := time2.Parse("2006-1-2 15:04", fmt.Sprintf("%d-%d-%d %d:%d", now.Year(), int(now.Month()), now.Day(), now.Hour(), now.Minute()))
	if err != nil {
		return false
	}
	return !timeStart.Before(time) && timeStart.Before(time2.Now().AddDate(0, 2, 0)) &&
		timeStart.Before(time2.Date(timeStart.Year(), timeStart.Month(), timeStart.Day(), 21, 1, 0, 0, timeStart.Location())) &&
		timeStart.After(time2.Date(timeStart.Year(), timeStart.Month(), timeStart.Day(), 8, 59, 0, 0, timeStart.Location()))
}
