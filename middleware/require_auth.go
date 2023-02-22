package middleware

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/kwamekyeimonies/Gin-Gonic-Rest-API/database"
	"github.com/kwamekyeimonies/Gin-Gonic-Rest-API/initiallizers"
	"github.com/kwamekyeimonies/Gin-Gonic-Rest-API/models"
)

func RequireAuth(ctx *gin.Context) {
	//Get Cookies of Requrest
	tokenString, err := ctx.Cookie("Authorization")
	if err != nil {
		ctx.AbortWithStatus(http.StatusUnauthorized)
	}
	//Decode/Validate the Request
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		config, _ := initiallizers.LoadConfig(".")
		return config.TokenSecret, nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		//Check the expiration time
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// Find the User with Token Sub
		var user models.Users
		database.DB.First(&user,claims["sub"])
		if user.Id == ""{
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		ctx.Set("user",user)
		ctx.Next()
	} else {
		ctx.AbortWithStatus(http.StatusUnauthorized)
	}

}
