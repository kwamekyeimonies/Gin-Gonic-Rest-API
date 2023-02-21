package repository

import "github.com/kwamekyeimonies/Gin-Gonic-Rest-API/model"

type UsersRepositories interface{

	Save(users model.Users)
	Update(users model.Users)
	Delete(userId int)
	FindById(userId int)(model.Users, error)
	FindAll() []model.Users
	FindByUsername(username string)(model.Users, error)
}