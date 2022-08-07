package main

import (
	"os"

	"github.com/FranciscoOrtizCastillo/go-react-project/config"
	"github.com/FranciscoOrtizCastillo/go-react-project/db"
	"github.com/FranciscoOrtizCastillo/go-react-project/routers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {

	config.Config()

	// Connect to database
	if err := db.Connect(); err != nil {
		panic(err)
	}

	app := fiber.New()

	app.Use(cors.New())

	app.Static("/", "./client/dist")

	routers.SetupRoutes(app)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	//fmt.Println("Server on port " + port)
	err := app.Listen(":" + port)

	if err != nil {
		panic(err)
	}
}
