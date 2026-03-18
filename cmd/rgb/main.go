package main

import (
	"log/slog"
	"os"
	"rgb/internal/handlers"
	"rgb/internal/repositories"
	"rgb/internal/server"
	"rgb/internal/services"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var validate = validator.New()

func main() {
	// 1. Build infrastructure
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	// Load .env file
	err := godotenv.Load()
	if err != nil {
		logger.Error("Error loading .env file", "error", err)
	}

	// Access variables with os.Getenv()
	db_host := os.Getenv("DB_HOST")
	db_user := os.Getenv("DB_USER")
	db_password := os.Getenv("DB_PWD")
	db_name := os.Getenv("DB_NAME")
	db_port := os.Getenv("DB_PORT")

	// Define the DSN (Data Source Name)
	// The driver uses pgx as the underlying driver by default.
	dsn := "host=" + db_host + " user=" + db_user + " password=" + db_password + " dbname=" + db_name + " port=" + db_port + " sslmode=disable TimeZone=US/Eastern"

	// Open the connection
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Error("Failed to connect to database", "error", err)
	}

	validate = validator.New()

	// 2. Create repositories
	itemRepository := repositories.NewItemRepository(db)
	shipmentRepository := repositories.NewShipmentRepository(db)

	// 3. Create services
	itemService := services.NewItemService(itemRepository, logger)
	shipmentService := services.NewShipmentService(shipmentRepository, logger)

	// 4. Create handlers
	itemHandler := handlers.NewItemHandler(itemService, validate, logger)
	shipmentHandler := handlers.NewShipmentHandler(shipmentService, validate, logger)

	// 5. Start server and pass in handlers
	server.Start(itemHandler, shipmentHandler)
}
