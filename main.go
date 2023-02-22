package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/kwamekyeimonies/Gin-Gonic-Rest-API/database"
	"github.com/kwamekyeimonies/Gin-Gonic-Rest-API/initiallizers"
	"github.com/kwamekyeimonies/Gin-Gonic-Rest-API/models"
	"github.com/kwamekyeimonies/Gin-Gonic-Rest-API/router"
)

func main() {

	loadConfig, err := initiallizers.LoadConfig(".")
	if err != nil {
		fmt.Println(err.Error())
	}

	port := loadConfig.ServerPort
	database.Database_Connector()
	database.DB.Table("book").AutoMigrate(&models.Book{})
	database.DB.Table("users").AutoMigrate(&models.Users{})

	// gin_server := gin.Default()

	// gin_server.GET("/test", func(ctx *gin.Context){
	// 	ctx.JSON(http.StatusOK,gin.H{
	// 		"Message":"Welcome to Gin-Framework",
	// 	})
	// })

	gin_server := gin.Default()
	gin_server.Use(gin.Logger())
	router.UserRouter(gin_server)

	gin_server.Run(":" + port)
}
