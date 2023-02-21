package main

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/kwamekyeimonies/Gin-Gonic-Rest-API/config"
	"github.com/kwamekyeimonies/Gin-Gonic-Rest-API/controller"
	"github.com/kwamekyeimonies/Gin-Gonic-Rest-API/database"
	"github.com/kwamekyeimonies/Gin-Gonic-Rest-API/helper"
	"github.com/kwamekyeimonies/Gin-Gonic-Rest-API/model"
	"github.com/kwamekyeimonies/Gin-Gonic-Rest-API/repository"
	"github.com/kwamekyeimonies/Gin-Gonic-Rest-API/router"
	"github.com/kwamekyeimonies/Gin-Gonic-Rest-API/service"
)

func main() {

	// if err := godotenv.Load(); err != nil {
	// 	log.Fatal(err)
	// }

	loadConfig, err := config.LoadConfig(".")
	helper.Error_Log(err)

	// gin_server := gin.Default()

	db := database.DB_Connector(&loadConfig)
	validate := validator.New()
	db.Table("users").AutoMigrate(&model.Users{})

	usersRepository := repository.NewusersRepositoryImpl(db)

	authenticationService := service.NewAuthenticationServiceImpl(usersRepository, validate)

	authenticationController := controller.NewAuthenticationController(authenticationService)
	userController := controller.NewUsersController(usersRepository)

	routes := router.NewRouter(usersRepository,authenticationController,userController)

	// gin_server.GET("/", func(ctx *gin.Context) {
	// 	ctx.JSON(http.StatusOK, gin.H{"message": "Welcome to Gin-Gonic Framework"})
	// })

	server := &http.Server{
		Addr:    ":8888",
		Handler: routes,
	}

	server_err := server.ListenAndServe()
	helper.Error_Log(server_err)
}
