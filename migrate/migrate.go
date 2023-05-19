package main

import (
	"fmt"

	"github.com/beslintony/go-crud/initializers"
	"github.com/beslintony/go-crud/models"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectToDatabase()
}

func main() {
	// Migrate the schema
	initializers.DB.AutoMigrate(&models.Post{})
	fmt.Println("Migration completed")
}
