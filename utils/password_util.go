package utils

import (
	"github.com/kwamekyeimonies/Gin-Gonic-Rest-API/helper"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string)(string,error){
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	helper.Error_Log(err)
	return string(hashedPassword),nil
}


func VerifyPassword(hashedPassword string,userPassword string) error{

	return bcrypt.CompareHashAndPassword([]byte(hashedPassword),[]byte(userPassword))
}