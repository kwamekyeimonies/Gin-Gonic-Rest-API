package controllers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/kwamekyeimonies/Gin-Gonic-Rest-API/database"
	"github.com/kwamekyeimonies/Gin-Gonic-Rest-API/initiallizers"
	"github.com/kwamekyeimonies/Gin-Gonic-Rest-API/models"
	utils "github.com/kwamekyeimonies/Gin-Gonic-Rest-API/utils"
)

var validate = validator.New()

func SignUp(ctx *gin.Context) {

	var users models.Users
	if err := ctx.BindJSON(&users); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Error": err.Error(),
		})
		return
	}
	validationErr := validate.Struct(users)
	if validationErr != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Error": validationErr.Error(),
		})
		return
	}

	users.Id = uuid.New().String()

	//Hash Password
	// hahsed_password, err := bcrypt.GenerateFromPassword([]byte(users.Password), bcrypt.DefaultCost)
	// if err != nil {
	// 	ctx.JSON(http.StatusBadRequest, gin.H{
	// 		"Error": err.Error(),
	// 	})
	// 	return
	// }
	// users.Password = string(hahsed_password)
	hashed_password, err := utils.HashPassword(users.Password)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Error": err.Error(),
		})
		return
	}
	users.Password = hashed_password

	//Insert Into Database
	result := database.DB.Create(&users)
	if result.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Error": result.Error.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"Response": "Account Created Successfully",
	})
}

func Login(ctx *gin.Context) {
	var user models.Users
	var userInput models.Users

	if err := ctx.BindJSON(&userInput); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Error": err.Error(),
		})
		return
	}

	database.DB.First(&user, "email = ?", userInput.Email)

	if user.Id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Error": "Invalide Id",
		})
		return
	}

	//Compare Password
	verify_err := utils.VerfiyPassword(user.Password, userInput.Password)
	if verify_err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Error": "Invalid Password",
		})
		return
	}

	//Generate JWT Token
	config, _ := initiallizers.LoadConfig(".")
	// token, err_token : utils.GenerateToken(config.TokenExpirationTime, user.Id, config.TokenSecret)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.Id,
		"nbf": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(config.TokenSecret))

	fmt.Println(tokenString, err)

	ctx.JSON(http.StatusOK, gin.H{
		"Token": tokenString,
	})
}
