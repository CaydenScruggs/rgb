package main

// import "rgb/internal/server"

// func main() {
// 	server.Start()
// }

// import (
// 	"database/sql"
// 	"fmt"

// 	_ "github.com/lib/pq"
// )

// func main() {
// 	// Connect to the PostgreSQL database
// 	connStr := "user=rgbuser dbname=rgb password=rgbuser sslmode=disable"
// 	db, err := sql.Open("postgres", connStr)
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer db.Close()

// 	// Perform a sample query
// 	rows, err := db.Query("SELECT * FROM cars")
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer rows.Close()

// 	// Iterate through the results
// 	for rows.Next() {
// 		var brand string
// 		var model string
// 		var year uint32
// 		if err := rows.Scan(&brand, &model, &year); err != nil {
// 			panic(err)
// 		}
// 		fmt.Printf("Brand: %s, Model: %s, Year: %d\n", brand, model, year)
// 	}
// }

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Car struct {
	Brand string
	Model string
	Year  uint
}

func main() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
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
		log.Fatal("Failed to connect to database:", err)
	}

	fmt.Println("Successfully connected to the PostgreSQL database!")

	// You can then use the 'db' object to perform operations
	// e.g., AutoMigrate, Create, Find, etc.

	var car Car

	result := db.Take(&car)

	if result.Error != nil {
		log.Fatal("Query failed:", result.Error)
	}

	fmt.Printf("First car: %s %s (%d)\n", car.Brand, car.Model, car.Year)
}
