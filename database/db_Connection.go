package database

import (
	"fmt"

	"github.com/kwamekyeimonies/Gin-Gonic-Rest-API/config"
	"github.com/kwamekyeimonies/Gin-Gonic-Rest-API/helper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DB_Connector(config *config.Config) *gorm.DB {

	ConnectionString := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", config.DBHost, config.DBPort, config.DBUsername,config.DBName, config.DBPassword)
	fmt.Println(ConnectionString)

	db, err := gorm.Open(postgres.Open(ConnectionString), &gorm.Config{})
	helper.Error_Log(err)

	fmt.Println("Database Connected Successfully...")
	return db
}
