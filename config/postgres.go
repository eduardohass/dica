package config

import (
	"github.com/eduardohass/dica/schemas"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitializePostgres() (*gorm.DB, error) {
	logger := GetLogger("postgres")
	dbURL := "postgres://postgres:changeme@0.0.0.0:5432/dica"

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
