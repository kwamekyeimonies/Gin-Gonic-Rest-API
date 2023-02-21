package repository

import (
	"errors"

	"github.com/kwamekyeimonies/Gin-Gonic-Rest-API/data/request"
	"github.com/kwamekyeimonies/Gin-Gonic-Rest-API/helper"
	"github.com/kwamekyeimonies/Gin-Gonic-Rest-API/model"
	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	Db *gorm.DB
}


func NewusersRepositoryImpl(Db *gorm.DB) UsersRepositories {
	return &UserRepositoryImpl{Db: Db}
}


// Delete implements UsersRepositories
func (user *UserRepositoryImpl) Delete(userId int) {
	var users model.Users
	result := user.Db.Where("id = ?",userId).Delete(&users)
	helper.Error_Log(result.Error)
}

// FindAll implements UsersRepositories
func (user *UserRepositoryImpl) FindAll() []model.Users {
	var users []model.Users
	result := user.Db.Find(&users)
	helper.Error_Log(result.Error)
	return users
}

// FindById implements UsersRepositories
func (user *UserRepositoryImpl) FindById(userId int) (model.Users, error) {
	var users model.Users
	result := user.Db.Find(&users, userId)
	if result != nil{
		return users,nil
	}else{
		return users, errors.New("User Not Found")
	}
}

// FindByUsername implements UsersRepositories
func (user *UserRepositoryImpl) FindByUsername(username string) (model.Users, error) {
	var users model.Users
	result := user.Db.First(&users, "username = ?",username)
	if result.Error != nil{
		return users, errors.New("Invalid Username or Password")
	}
	return users, nil
}

// Save implements UsersRepositories
func (user *UserRepositoryImpl) Save(users model.Users) {
	result := user.Db.Create(&users)
	helper.Error_Log(result.Error)
}

// Update implements UsersRepositories
func (user *UserRepositoryImpl) Update(users model.Users) {
	var updateUsers = request.UpdateUserRequest{
		Id:users.Id,
		Username: users.Username,
		Email: users.Email,
		Password: users.Password,
	}

	result := user.Db.Model(&users).Updates(updateUsers)
	helper.Error_Log(result.Error)
}

