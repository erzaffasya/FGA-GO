package books

import (
	"belajar_gin/repositories/books"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetBook(ctx *gin.Context) {
	id := ctx.Param("book_id")
	condition := false
	var bookData books.Book

	for i, book := range books.BookDatas {
		if id == book.Id {
			condition = true
			bookData = books.BookDatas[i]
			break
		}
	}

	if !condition{
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "Data Not Found",
			"error_message": fmt.Sprintf("book with id %v not found", id),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"book": bookData,
	})
}