package handler

import (
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
	a := timeStart.Year() >= time2.Now().Year()
	b := timeStart.Month() >= time2.Now().Month()
	c := timeStart.Day() >= time2.Now().Day()
	d := timeStart.Hour() >= time2.Now().Hour()
	e := timeStart.Minute() >= time2.Now().Minute()
	return a && b && c && d && e && timeStart.Before(time2.Now().AddDate(0, 2, 0)) &&
		timeStart.Before(time2.Date(timeStart.Year(), timeStart.Month(), timeStart.Day(), 20, 1, 0, 0, timeStart.Location())) &&
		timeStart.After(time2.Date(timeStart.Year(), timeStart.Month(), timeStart.Day(), 8, 59, 0, 0, timeStart.Location()))
}
