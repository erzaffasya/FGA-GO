package books

import (
	"belajar_gin/repositories/books"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateBook(ctx *gin.Context) {
	var newBook books.Book

	err := ctx.ShouldBindJSON(&newBook)
	if err!=nil {
		ctx.AbortWithError(http.StatusBadRequest,err)
		return
	}
	newBook.Id = fmt.Sprintf("b%d",len(books.BookDatas)+1)
	books.BookDatas = append(books.BookDatas,newBook)

	ctx.JSON(http.StatusCreated,gin.H{
		"book":newBook,
	})
}