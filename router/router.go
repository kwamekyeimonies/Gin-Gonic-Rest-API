package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kwamekyeimonies/Gin-Gonic-Rest-API/controller"
	"github.com/kwamekyeimonies/Gin-Gonic-Rest-API/middleware"
	"github.com/kwamekyeimonies/Gin-Gonic-Rest-API/repository"
)

func NewRouter(userRepository repository.UsersRepositories, authenticationController *controller.AuthenticationController, userController *controller.UserController) *gin.Engine {
	service := gin.Default()

	service.GET("", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "Welcome Home")
	})

	router := service.Group("/api")
	authenticationRouter := router.Group("/auth")
	authenticationRouter.POST("/register", authenticationController.Register)
	authenticationRouter.POST("/login", authenticationController.Login)

	usersRouter := router.Group("/users")
	usersRouter.GET("", middleware.DeserializeUsers(userRepository), userController.GetUsers)

	return service
}
