package config

import (
	"fmt"
	"os"

	"github.com/eduardohass/dica/schemas"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitializePostgres() (*gorm.DB, error) {
	logger := GetLogger("postgres")
	// dbURL := "postgres://postgres:changeme@postgres:5432/dica"
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

	// Migrate the answer schema
	err = db.AutoMigrate(&schemas.Answer{})
	if err != nil {
		logger.Errorf("postgres automigration error - answer %v", err)
		return nil, err
	}

	// Migrate the question schema
	err = db.AutoMigrate(&schemas.Question{})
	if err != nil {
		logger.Errorf("postgres automigration error - question %v", err)
		return nil, err
	}

	fmt.Println("DBG==Conectou no DB")
	fmt.Println("DBG==db: ", db)
	// Return the DB
	return db, nil
}
