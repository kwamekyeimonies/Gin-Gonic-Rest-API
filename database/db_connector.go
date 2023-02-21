package database

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/kwamekyeimonies/Gin-Gonic-Rest-API/models"
)

var DB *gorm.DB

func Database_Connection() {

	dbHost := os.Getenv("DATABASE_HOST")
	dbPort := os.Getenv("DATABASE_PORT")
	dbUser := os.Getenv("DATABASE_USERNAME")
	dbName := os.Getenv("DATABASE_NAME")
	dbPass := os.Getenv("DATABASE_PASSWORD")

	ConnectionString := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s", dbHost, dbPort, dbUser, dbName, dbPass)

	db, err := gorm.Open("postgres", ConnectionString)

	if err != nil {
		fmt.Println("Error Connecting to database...")
	} else {
		fmt.Println("Database Connected Successfully")
	}

	db.AutoMigrate(
		&models.Book{},
	)

	DB = db
}
