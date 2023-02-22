package database

import (
	"fmt"

	"github.com/kwamekyeimonies/Gin-Gonic-Rest-API/initiallizers"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Database_Connector() {

	config,_ := initiallizers.LoadConfig(".")
	
	// Connection_URL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
	// config.DBHost,
	// config.DBPort,
	// config.DBUsername,
	// config.DBName,
	// config.DBPassword)
	// fmt.Println(Connection_URL)

	Connect_Elephant_SQL := fmt.Sprintf(config.ELEPHANTSQL_URL)

	db, err := gorm.Open(postgres.Open(Connect_Elephant_SQL), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())

	}

	fmt.Println("Database Connected successfully.....")

	DB = db

	
}
