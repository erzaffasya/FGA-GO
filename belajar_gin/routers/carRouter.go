package routers

import (
	"belajar_gin/controllers/cars"

	"github.com/gin-gonic/gin"
)

func CarRouter(router gin.Engine) *gin.Engine{
	router.POST("/cars",cars.CreateCar)
	router.GET("/cars",cars.GetAllCar)
	router.GET("/cars/:car_id",cars.GetCar)
	router.PUT("/cars/:car_id",cars.UpdateCar)
	router.DELETE("/cars/:car_id",cars.DeleteCar)
	return &router
}