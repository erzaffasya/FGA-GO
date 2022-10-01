package cars

import (
	"belajar_gin/repositories/cars"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)
func UpdateCar(ctx *gin.Context){
	carID := ctx.Param(("car_id"))
	condition := false
	var updatedCar cars.Car

	if err := ctx.ShouldBindJSON(&updatedCar);err!= nil{
		ctx.AbortWithError(http.StatusBadRequest,err)
		return
	}

	for i, car := range cars.CarDatas {
		if carID == car.CarID{
			condition = true
			cars.CarDatas[i] = updatedCar
			cars.CarDatas[i].CarID = carID
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
	ctx.JSON(http.StatusOK,gin.H{
		"message":fmt.Sprintf("car with id %v has been successfully updated",carID),
	})
}