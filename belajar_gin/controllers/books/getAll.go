package books

import (
	"belajar_gin/repositories/books"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllBook(ctx *gin.Context)  {
	if len(books.BookDatas)>0{
		ctx.JSON(http.StatusOK, gin.H{
			"books": books.BookDatas,
		})		
	}else{
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "Data Not Found",
			"error_message": "Data book masih kosong",
		})		
	}
	
} 

