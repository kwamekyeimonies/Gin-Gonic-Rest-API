package router

import (
	"github.com/gin-gonic/gin"
	"github.com/kwamekyeimonies/Gin-Gonic-Rest-API/controllers"
)

func UserRouter(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("/signup", controllers.SignUp)
	incomingRoutes.POST("/login",controllers.Login)
}
