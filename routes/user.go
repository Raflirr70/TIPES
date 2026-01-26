package routes

import (
	"tipes/internal/config"
	"tipes/internal/handler"
	"tipes/internal/repository"
	"tipes/internal/service"

	"github.com/gin-gonic/gin"
)

func GetAirplane(r *gin.Engine) {
	airplaneRepo := repository.NewAirplaneRepository(config.DB)
	airplaneService := service.NewAirplaneService(airplaneRepo)
	airplaneHandler := handler.NewAirplaneHandler(airplaneService)

	r.GET("/ping", airplaneHandler.Ping)
}

func Get(r *gin.Engine) {
	GetAirplane(r)
}
func Post(r *gin.Engine) {

}
