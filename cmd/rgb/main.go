package main

import (
	"log/slog"
	"os"
	"rgb/internal/handlers"
	"rgb/internal/server"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func main() {
	// 1. Build infrastructure
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	// // Load .env file
	// err := godotenv.Load()
	// if err != nil {
	// 	logger.Error("Error loading .env file: %s", err)
	// }

	// // Access variables with os.Getenv()
	// db_host := os.Getenv("DB_HOST")
	// db_user := os.Getenv("DB_USER")
	// db_password := os.Getenv("DB_PWD")
	// db_name := os.Getenv("DB_NAME")
	// db_port := os.Getenv("DB_PORT")

	// // Define the DSN (Data Source Name)
	// // The driver uses pgx as the underlying driver by default.
	// dsn := "host=" + db_host + " user=" + db_user + " password=" + db_password + " dbname=" + db_name + " port=" + db_port + " sslmode=disable TimeZone=US/Eastern"

	// // Open the connection
	// db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	// if err != nil {
	// 	logger.Error("Failed to connect to database:", err)
	// }

	validate = validator.New()

	// 2. Create handlers
	// TODO: Later, we will create repositories, services, etc.
	itemHandler := handlers.NewItemHandler(validate, logger)
	shipmentHandler := handlers.NewShipmentHandler(validate, logger)

	// 3. Start server and pass in handlers
	server.Start(itemHandler, shipmentHandler)
}

// import (
// 	"fmt"
// 	"log"
// 	"os"

// 	"github.com/joho/godotenv"
// 	"gorm.io/driver/postgres"
// 	"gorm.io/gorm"
// )

// type Car struct {
// 	Brand string
// 	Model string
// 	Year  uint
// }

// func main() {
// 	// Load .env file
// 	err := godotenv.Load()
// 	if err != nil {
// 		log.Fatalf("Error loading .env file: %s", err)
// 	}

// 	// Access variables with os.Getenv()
// 	db_host := os.Getenv("DB_HOST")
// 	db_user := os.Getenv("DB_USER")
// 	db_password := os.Getenv("DB_PWD")
// 	db_name := os.Getenv("DB_NAME")
// 	db_port := os.Getenv("DB_PORT")

// 	// Define the DSN (Data Source Name)
// 	// The driver uses pgx as the underlying driver by default.
// 	dsn := "host=" + db_host + " user=" + db_user + " password=" + db_password + " dbname=" + db_name + " port=" + db_port + " sslmode=disable TimeZone=US/Eastern"

// 	// Open the connection
// 	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 		log.Fatal("Failed to connect to database:", err)
// 	}

// 	fmt.Println("Successfully connected to the PostgreSQL database!")

// 	// You can then use the 'db' object to perform operations
// 	// e.g., AutoMigrate, Create, Find, etc.

// 	var car Car

// 	result := db.Take(&car)

// 	if result.Error != nil {
// 		log.Fatal("Query failed:", result.Error)
// 	}

// 	fmt.Printf("First car: %s %s (%d)\n", car.Brand, car.Model, car.Year)
// }
