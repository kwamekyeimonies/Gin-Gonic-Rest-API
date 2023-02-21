package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kwamekyeimonies/Gin-Gonic-Rest-API/data/request"
	"github.com/kwamekyeimonies/Gin-Gonic-Rest-API/data/response"
	"github.com/kwamekyeimonies/Gin-Gonic-Rest-API/helper"
	"github.com/kwamekyeimonies/Gin-Gonic-Rest-API/service"
)

type AuthenticationController struct {
	authenticationService service.AuthenticationService
}

func NewAuthenticationController(service service.AuthenticationService) *AuthenticationController {
	return &AuthenticationController{authenticationService: service}
}

func (controller *AuthenticationController) Login(ctx *gin.Context) {
	loginRequest := request.LoginRequest{}
	err := ctx.ShouldBindJSON(&loginRequest)
	helper.Error_Log(err)

	token, err_token := controller.authenticationService.Login(loginRequest)
	fmt.Println(err_token)

	if err_token != nil {
		webResponse := response.Response{
			Code:    http.StatusBadRequest,
			Status:  "Bad Request",
			Message: "Invalide Username of password",
		}
		ctx.JSON(http.StatusBadRequest, webResponse)
		return
	}

	resp := response.LoginRespone{
		TokenType: "Bearer",
		Token:     token,
	}

	webResponse := response.Response{
		Code:    200,
		Status:  "Ok",
		Message: "Successfully Logged in",
		Data:    resp,
	}

	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *AuthenticationController) Register(ctx *gin.Context) {
	createUsersRequest := request.CreateUserRequest{}
	err := ctx.ShouldBindJSON(&createUsersRequest)
	helper.Error_Log(err)

	controller.authenticationService.Register(createUsersRequest)
	webResponse := response.Response{
		Code:    200,
		Status:  "Ok",
		Message: "Successfully Created User",
		Data:    nil,
	}

	ctx.JSON(http.StatusOK, webResponse)
}
