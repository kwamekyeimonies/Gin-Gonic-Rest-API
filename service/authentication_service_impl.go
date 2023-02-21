package service

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/kwamekyeimonies/Gin-Gonic-Rest-API/config"
	"github.com/kwamekyeimonies/Gin-Gonic-Rest-API/data/request"
	"github.com/kwamekyeimonies/Gin-Gonic-Rest-API/helper"
	"github.com/kwamekyeimonies/Gin-Gonic-Rest-API/model"
	"github.com/kwamekyeimonies/Gin-Gonic-Rest-API/repository"
	"github.com/kwamekyeimonies/Gin-Gonic-Rest-API/utils"
)

type AuthenticationServiceImpl struct {
	UsersRepoitory repository.UsersRepositories
	Validate       *validator.Validate
}

func NewAuthenticationServiceImpl(usersRepository repository.UsersRepositories, validate *validator.Validate) AuthenticationService {
	return &AuthenticationServiceImpl{
		UsersRepoitory: usersRepository,
		Validate:       validate,
	}
}

// Login implements AuthenticationService
func (auth *AuthenticationServiceImpl) Login(users request.LoginRequest) (string, error) {
	//Find Username in database
	new_user, user_err := auth.UsersRepoitory.FindByUsername(users.Username)
	if user_err != nil {
		fmt.Println(http.StatusBadRequest, user_err)
	}

	config, _ := config.LoadConfig(".")
	verify_error := utils.VerifyPassword(new_user.Password, users.Password)
	// helper.Error_Log(verify_error)
	if verify_error != nil {
		fmt.Println(http.StatusBadRequest, gin.H{"error": verify_error})
	}

	//Generate Token
	token, err_token := utils.GenerateToken(config.TokenExpirationTime, new_user.Id, config.TokenSecret)
	helper.Error_Log(err_token)
	return token, nil
}

// Register implements AuthenticationService
func (auth *AuthenticationServiceImpl) Register(users request.CreateUserRequest) {
	hashedPassword, err := utils.HashPassword(users.Password)
	helper.Error_Log(err)

	newUser := model.Users{
		Username: users.Username,
		Email:    users.Email,
		Password: hashedPassword,
	}

	auth.UsersRepoitory.Save(newUser)
}
