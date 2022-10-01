package books

import (
	"belajar_gin/repositories/books"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UpdateBook(ctx *gin.Context){
	book_id := ctx.Param(("book_id"))
	condition := false
	var updatedBook books.Book

	if err := ctx.ShouldBindJSON(&updatedBook);err!= nil{
		ctx.AbortWithError(http.StatusBadRequest,err)
		return
	}

	for i, book := range books.BookDatas {
		if book_id == book.Id{
			condition = true
			books.BookDatas[i] = updatedBook
			books.BookDatas[i].Id = book_id
			break
		}
	}
	if !condition {
		ctx.AbortWithStatusJSON(http.StatusNotFound,gin.H{
			"error_status": "Data Not Found",
			"rrror_message":fmt.Sprintf("book with id %v not found",book_id),
		})
		return
	}
	ctx.JSON(http.StatusOK,gin.H{
		"message":fmt.Sprintf("book with id %v has been successfully updated",book_id),
	})
}