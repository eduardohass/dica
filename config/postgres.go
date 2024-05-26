package config

import (
	"github.com/eduardohass/dica/schemas"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
	"fmt"
)

func InitializePostgres() (*gorm.DB, error) {
	logger := GetLogger("postgres")
	// dbURL := "postgres://postgres:changeme@postgres:5432/dica"
	// dbURL := os.Getenv("CON_STR")
	dbURL := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=America/Sao_Paulo",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	// Create DB and connect
	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		logger.Errorf("postgres opening error %v", err)
		return nil, err
	}
	// Migrate the schema
	err = db.AutoMigrate(&schemas.Opening{})
	if err != nil {
		logger.Errorf("postgres automigration error %v", err)
		return nil, err
	}
	// Return the DB
	return db, nil
}
