package database

import (
	"log"
	"os"

	"github.com/robjullian/fffgrenoble.fr/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// DB Global DB pointer
var DB *gorm.DB

// Check if database file exist
func check() {
	// if exist
	_, err := os.Stat(os.Getenv("DB_PATH") + os.Getenv("DB_FILE"))
	if err != nil {
		err := os.MkdirAll(os.Getenv("DB_PATH"), 0777) // create directories if don't exist
		if err != nil {
			log.Fatal(err)
		}

		file, err := os.Create(os.Getenv("DB_PATH") + os.Getenv("DB_FILE")) // Create file
		if err != nil {
			log.Fatal(err.Error())
		}
		file.Close()
	}
}

func Connect() error {
	check() // Check database existance

	// Create connection to DB
	connection, err := gorm.Open(sqlite.Open(os.Getenv("DB_PATH")+os.Getenv("DB_FILE")), &gorm.Config{
		FullSaveAssociations: true,
	})
	if err != nil {
		return err
	}

	DB = connection // Set global pointer

	connection.Logger.LogMode(3)

	// Remove all tables

	/*
		_ = connection.Migrator().DropTable(
		&models.Event{},
		&models.Link{},
		&models.Social{},
		)
	*/

	// Setup all tables
	_ = connection.AutoMigrate(
		&models.Event{},
		&models.Link{},
		&models.Social{},
	)

	return nil
}
