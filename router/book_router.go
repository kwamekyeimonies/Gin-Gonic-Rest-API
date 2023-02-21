package router

import (
	"github.com/gin-gonic/gin"
	"github.com/kwamekyeimonies/Gin-Gonic-Rest-API/controllers"
)

func BookRouter(incomingRoutes *gin.Engine){
	incomingRoutes.POST("/createbooks",controllers.CreateBooks())
	incomingRoutes.GET("/getbooks", controllers.GetAllBooks())
	incomingRoutes.GET("/getbook/:id",controllers.GetBookById())
}