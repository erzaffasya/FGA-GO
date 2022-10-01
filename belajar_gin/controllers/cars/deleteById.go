package cars

import (
	"belajar_gin/repositories/cars"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)
func DeleteCar(ctx *gin.Context){
	carID := ctx.Param(("car_id"))
	condition := false
	var carIndex int
	
	for i, car := range cars.CarDatas {
		if carID == car.CarID{
			condition = true
			carIndex = i
			break
		}
	}
	if !condition {
		ctx.AbortWithStatusJSON(http.StatusNotFound,gin.H{
			"error_status": "Data Not Found",
			"rrror_message":fmt.Sprintf("car with id %v not found",carID),
		})
		return
	}

	copy(cars.CarDatas[carIndex:],cars.CarDatas[carIndex+1:])
	cars.CarDatas[len(cars.CarDatas)-1] = cars.Car{}
	cars.CarDatas = cars.CarDatas[:len(cars.CarDatas)-1]

	ctx.JSON(http.StatusOK,gin.H{
		"message":fmt.Sprintf("car with id %v has been successfully deleted",carID),
	})
}