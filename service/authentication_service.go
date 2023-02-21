package service

import "github.com/kwamekyeimonies/Gin-Gonic-Rest-API/data/request"

type AuthenticationService interface{
	Login(users request.LoginRequest)(string,error)
	Register(users request.CreateUserRequest)
}