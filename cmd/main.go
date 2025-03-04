package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"github.com/vishalpandhare01/Predictive-Maintenance-System/db"
	"github.com/vishalpandhare01/Predictive-Maintenance-System/routes"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Env not loaded", err)
	}

	app := fiber.New()
	db.ConnectDb()

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:3000/, http://localhost:3001/, http://localhost:3002/",
		AllowCredentials: true,
	}))

	routes.SetupRoutes(app)

	app.Listen(":8080")
}
