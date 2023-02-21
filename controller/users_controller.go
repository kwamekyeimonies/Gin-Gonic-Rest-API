package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kwamekyeimonies/Gin-Gonic-Rest-API/data/response"
	"github.com/kwamekyeimonies/Gin-Gonic-Rest-API/repository"
)

type UserController struct{
	userRepository repository.UsersRepositories
}

func NewUsersController(repository repository.UsersRepositories) *UserController{
	return &UserController{userRepository: repository}
}


func (controller *UserController) GetUsers(ctx *gin.Context){
	users := controller.userRepository.FindAll()
	webResponse := response.Response{
		Code:200,
		Status: "Ok",
		Message: "Successfully Fetched all Users",
		Data:users,
	}

	ctx.JSON(http.StatusOK, webResponse)
}