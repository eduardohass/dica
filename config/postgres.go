package config

import (
	"github.com/eduardohass/dica/schemas"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitializePostgres() (*gorm.DB, error) {
	logger := GetLogger("postgres")
	dbURL := "postgres://postgres:changeme@postgres:5432/dica"
	// dbURL := os.Getenv("CON_STR")

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
