package config

import (
	"github.com/Murilojms7/Login-System/schemas"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitializePostgre() (*gorm.DB, error) {
	logger := GetLogger("postgre")

	// Create and Connect DataBase
	dsn := "host=localhost user=postgres password=1234 dbname=postgres port=5432 sslmode=disable"
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
