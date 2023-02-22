package utils

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)


func HashPassword(password string)(string,error){
	hahsed_password,err := bcrypt.GenerateFromPassword([]byte(password),bcrypt.DefaultCost)
	if err != nil{
		fmt.Println("Error Hashing Password")
		
	}
	return string(hahsed_password),nil
}

func VerfiyPassword(hahsed_password string,userInputPassword string) error{
	
	return bcrypt.CompareHashAndPassword([]byte(hahsed_password),[]byte(userInputPassword))
}