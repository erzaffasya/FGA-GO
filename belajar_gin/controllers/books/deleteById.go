package books

import (
	"belajar_gin/repositories/books"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeleteBook(ctx *gin.Context){
	book_id := ctx.Param(("book_id"))
	condition := false
	var bookIndex int
	
	for i, book := range books.BookDatas {
		if book_id == book.Id{
			condition = true
			bookIndex = i
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

	copy(books.BookDatas[bookIndex:],books.BookDatas[bookIndex+1:])
	books.BookDatas[len(books.BookDatas)-1] = books.Book{}
	books.BookDatas = books.BookDatas[:len(books.BookDatas)-1]

	ctx.JSON(http.StatusOK,gin.H{
		"message":fmt.Sprintf("book with id %v has been successfully deleted",book_id),
	})
}