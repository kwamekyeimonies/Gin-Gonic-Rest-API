package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/kwamekyeimonies/Gin-Gonic-Rest-API/database"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	database.Database_Connection()

	gin_server := gin.Default()

	gin_server.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "Welcome to Gin-Gonic Framework"})
	})

	gin_server.Run()
}
