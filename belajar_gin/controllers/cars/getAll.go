package cars

import (
	"belajar_gin/repositories/cars"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllCar(ctx *gin.Context)  {
	if len(cars.CarDatas)>0{
		ctx.JSON(http.StatusOK, gin.H{
			"cars": cars.CarDatas,
		})		
	}else{
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "Data Not Found",
			"error_message": "Data car masih kosong",
		})		
	}	
} 

