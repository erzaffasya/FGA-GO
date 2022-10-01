package routers

import (
	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine{
	router := gin.Default()
	CarRouter(*router)
	BookRouter(*router)
	return router
}