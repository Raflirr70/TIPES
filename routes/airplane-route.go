package routes

import (
	"tipes/internal/config"
	"tipes/internal/handler"
	"tipes/internal/repository"
	"tipes/internal/service"

	"github.com/gin-gonic/gin"
)

func AirplaneRoute(r *gin.Engine) {
	airplaneRepo := repository.NewAirplaneRepository(config.DB)
	airplaneService := service.NewAirplaneService(airplaneRepo)
	airplaneHandler := handler.NewAirplaneHandler(airplaneService)

	airplane := r.Group("/search-airplane")
	{
		airplane.GET("/", airplaneHandler.Ping)
		airplane.GET("/:name", airplaneHandler.Ping)
	}
}
