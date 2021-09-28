package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"github.com/robjullian/fffgrenoble.fr/database"
	"github.com/robjullian/fffgrenoble.fr/routes"
)

// Load env values
func init() {
	if err := godotenv.Load(); err != nil {
		log.Printf("Unable to load env values: %v\n", err)
		panic(err)
	} else {
		log.Println("Loaded env values successfully")
	}
}

func main() {
	app := fiber.New()
	app.Use(cors.New()) // Add cors

	// Create the connection to the database database
	if err := database.Connect(); err != nil {
		log.Printf("Unable to connect the database: %v\n", err.Error())
		panic(err)
	} else {
		log.Println("Successfully connected to the database")
	}

	// Serve Frontend
	app.Static("/", "./view/public")
	// Backend
	routes.SetupRoutes(app)

	// Get IP and Port to listen on
	listner := os.Getenv("SERVER_HOST") + ":" + os.Getenv("SERVER_PORT")

	// Create listener
	err := app.Listen(listner)
	if err != nil {
		log.Println("Unable to start server on [" + listner + "]")
		panic(err)
	} else {
		log.Println("Listening on [" + listner + "]")
	}
}
