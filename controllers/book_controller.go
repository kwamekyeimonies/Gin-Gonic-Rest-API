package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kwamekyeimonies/Gin-Gonic-Rest-API/models"
)


func CreateBooks() gin.HandlerFunc{
	return func(ctx *gin.Context) {

		var book_input models.Book

		if err := ctx.ShouldBindJSON(&book_input); err != nil{
			ctx.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
			return
		}
		

	}
}


func GetAllBooks() gin.HandlerFunc{
	return func(ctx *gin.Context) {

	}
}


func GetBookById() gin.HandlerFunc{
	return func(ctx *gin.Context) {

	}
}