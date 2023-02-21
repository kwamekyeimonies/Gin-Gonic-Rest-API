package middleware

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/kwamekyeimonies/Gin-Gonic-Rest-API/config"
	"github.com/kwamekyeimonies/Gin-Gonic-Rest-API/helper"
	"github.com/kwamekyeimonies/Gin-Gonic-Rest-API/repository"
	"github.com/kwamekyeimonies/Gin-Gonic-Rest-API/utils"
)

func DeserializeUsers(usersRepository repository.UsersRepositories) gin.HandlerFunc{
	return func(ctx *gin.Context) {
		var token string
		authorizationHeader := ctx.Request.Header.Get("Authorization")
		fields := strings.Fields(authorizationHeader)

		if len(fields) != 0 && fields[0] == "Bearer"{
			token = fields[1]
		}

		if token == ""{
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status":"Failed","message":"You are not logged in"})
			return
		}

		config,_ := config.LoadConfig(".")
		sub,err := utils.ValidateToken(token,config.TokenSecret)
		if err != nil{
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status":"Failed","message":err.Error()})
			return
		}

		id,err_id := strconv.Atoi(fmt.Sprint(sub))
		helper.Error_Log(err_id)
		result,err := usersRepository.FindById(id)
		if err != nil{
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status":"Failed","message":err.Error()})
			return
		}

		ctx.Set("currentUser",result.Username)
		ctx.Next()


	}

}