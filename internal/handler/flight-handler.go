package handler

import (
	"net/http"

	"tipes/internal/service"

	"github.com/gin-gonic/gin"
)

type FlightHandler struct {
	service service.FlightService
}

func NewFlightHandler(service service.FlightService) *FlightHandler {
	return &FlightHandler{service}
}

func (h *FlightHandler) Ping(c *gin.Context) {
	from := c.Query("city_from")
	to := c.Query("city_to")
	airplaneName := c.Query("airplane_name")
	seatClassName := c.Query("seat_class_name")

	Flights, err := h.service.GetAll(from, to, airplaneName, seatClassName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": Flights,
	})
}
