package handler

import (
	"net/http"

	"tipes/internal/service"

	"github.com/gin-gonic/gin"
)

type AirplaneHandler struct {
	service service.AirplaneService
}

func NewAirplaneHandler(service service.AirplaneService) *AirplaneHandler {
	return &AirplaneHandler{service}
}

func (h *AirplaneHandler) Ping(c *gin.Context) {
	airplanes, err := h.service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": airplanes,
	})
}
