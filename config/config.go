package config

import (
	"fmt"

	"gorm.io/gorm"
)

var (
	db     *gorm.DB
	logger *Logger
)

func Init() error {
	var err error

	// // Initialize SQLite
	// db, err = InitializeSQLite()

	// Initialize SQLite
	db, err = InitializePostgres()

	// if err != nil {
	// 	return fmt.Errorf("Error initializing sqlite %v", err)
	// }

	if err != nil {
		return fmt.Errorf("Error initializing postgres %v", err)
	}

	// return errors.New("fake error")
	return nil
}

func GetDatabase() *gorm.DB {
	return db
}

func GetLogger(p string) *Logger {
	// Initialize Logger
	logger = NewLogger(p)
	return logger
}
