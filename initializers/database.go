package initializers

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB // Declare a package-level variable

func ConnectToDatabase() {
	var err error
	dsn := os.Getenv("DB_STRING")
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{}) // Assign the value to the package-level variable

	if err != nil {
		log.Fatal("Failed to connect to the database")
	}
}
