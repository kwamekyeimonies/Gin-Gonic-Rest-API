package models

import "time"

type User_Auth struct {
	ID        string    `json:"id" gorm:"primary_Key"`
	Username  string    `json:"username" binding:"required"`
	Password  string    `json:"password" binding:"required"`
	Email     string    `gorm:"uniqueIndex;not null"`
	Role      string    `gorm:"type:varchar(200);not null"`
	Photo     string    `gorm:"not null"`
	Verified  bool      `gorm:"not null"`
	UpdatedAt time.Time `gorm:"not null"`
	CreatedAt time.Time `gorm:"not null"`
	Provider string `gorm:"not null"`
}
