package utils

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string)(string,error){

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password),bcrypt.DefaultCost)

	if err != nil{
		fmt.Println("Unable to Hash Password")
	}
	return string(hashedPassword), nil
}

func VerifyPassword(hashedPassword string, userPassword string) error{

	return bcrypt.CompareHashAndPassword([]byte(hashedPassword),[]byte(userPassword))
}