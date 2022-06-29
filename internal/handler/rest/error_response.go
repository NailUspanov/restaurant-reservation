package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type errResponse struct {
	Message string `json:"message"`
}

func newErrorResponse(c *gin.Context, statusCode int, message string) {
	logrus.Error(message)
	c.AbortWithStatusJSON(statusCode, errResponse{message})
}
