package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
	"restaurant-reservation/pkg/models"
)

func (h *Handler) createReservation(c *gin.Context) {

	var input models.ReservationRequest
	if err := c.ShouldBindBodyWith(&input, binding.JSON); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	newReservationId, err := h.services.Reservation.Create(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]any{
		"id":             newReservationId,
		"restaurant":     input.Restaurant,
		"customer_name":  input.CustomerName,
		"customer_phone": input.CustomerPhone,
		"table":          input.Table,
		"time":           input.Time,
	})
}
