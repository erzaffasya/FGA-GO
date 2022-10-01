package routers

import (
	"belajar_gin/controllers/books"

	"github.com/gin-gonic/gin"
)

func BookRouter(router gin.Engine) *gin.Engine{	
	router.POST("/books",books.CreateBook)
	router.GET("/books",books.GetAllBook)
	router.GET("/books/:book_id",books.GetBook)
	router.PUT("/books/:book_id",books.UpdateBook)
	router.DELETE("/books/:book_id",books.DeleteBook)
	return &router
}
