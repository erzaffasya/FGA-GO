package cars

import (
	"belajar_gin/repositories/cars"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateCar(ctx *gin.Context){
	var newCar cars.Car
	err := ctx.ShouldBindJSON(&newCar)
	if err!=nil {
		ctx.AbortWithError(http.StatusBadRequest,err)
		return
	}
	newCar.CarID = fmt.Sprintf("c%d",len(cars.CarDatas)+1)
	cars.CarDatas = append(cars.CarDatas,newCar)

	ctx.JSON(http.StatusCreated,gin.H{
		"car":newCar,
	})
}