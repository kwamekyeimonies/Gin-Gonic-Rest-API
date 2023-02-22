package models

import "gorm.io/gorm"

type Users struct {
	gorm.Model
	Id          string
	Email       string `gorm:"varchar(255;not null;unique"`
	Password    string `gorm:"required"`
	PhoneNumber string `gorm:"required"`
}
