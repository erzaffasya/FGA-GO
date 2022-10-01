package cars

import (
	"belajar_gin/repositories/cars"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)
func GetCar(ctx *gin.Context){
	carID := ctx.Param("car_id")
	condition :=false
	var carData cars.Car

	for i, car:= range cars.CarDatas {
		if carID == car.CarID{
			condition =true
			carData = cars.CarDatas[i]
			break
		}
	}

	if !condition {
		ctx.AbortWithStatusJSON(http.StatusNotFound,gin.H{
			"error_status": "Data Not Found",
			"error_message":fmt.Sprintf("car with id %v not found",carID),
		})
		return
	}
	ctx.JSON(http.StatusOK,gin.H{
		"car":carData,
	})
}