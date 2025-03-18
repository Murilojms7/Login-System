package config

import (
	"fmt"
	"os"

	"github.com/Murilojms7/Login-System/schemas"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitializePostgre() (*gorm.DB, error) {
	logger := GetLogger("postgre")

	// Create and Connect DataBase
	err := godotenv.Load()
	if err != nil {
		logger.Errorf("Postgre opening error: %v", err)
		return nil, err
	}

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Errorf("Postgre opening error: %v", err)
		return nil, err
	}

	// Auto Migrate Model
	if err = db.AutoMigrate(&schemas.Users{}); err != nil {
		logger.Errorf("Postgre automigration error: %v", err)
		return nil, err
	}
	return db, nil
}
