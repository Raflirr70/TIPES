package routes

import (
	"tipes/internal/config"
	"tipes/internal/handler"
	"tipes/internal/repository"
	"tipes/internal/service"

	"github.com/gin-gonic/gin"
)

func FlightRoute(r *gin.Engine) {
	flightRepo := repository.NewFlightRepository(config.DB)
	flightService := service.NewFlightService(flightRepo)
	flightHandler := handler.NewFlightHandler(flightService)

	flight := r.Group("/search-flight")
	{
		flight.GET("/", flightHandler.Ping)
		flight.GET("/:name", flightHandler.Ping)
	}
}
