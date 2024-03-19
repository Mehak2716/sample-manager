package config

import (
	"fmt"
	"log"

	"github.com/Mehak2716/sample-manager/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	DB_HOST     string = "localhost"
	DB_PORT     uint32 = 5432
	DB_USER     string = "postgres"
	DB_NAME     string = "instamart"
	DB_PASSWORD string = "postgres"
)

func DatabaseConnection() *gorm.DB {

	dsn := fmt.Sprintf("host=%v port=%v user=%v dbname=%v password=%v sslmode=disable", DB_HOST, DB_PORT, DB_USER, DB_NAME, DB_PASSWORD)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to the database...", err)
	}

	log.Println("Database connection successful...")
	db.AutoMigrate(&models.SampleMapping{})
	return db
}
