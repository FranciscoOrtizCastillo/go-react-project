package routers

import (
	"context"

	"github.com/FranciscoOrtizCastillo/go-react-project/db"
	"github.com/FranciscoOrtizCastillo/go-react-project/models"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

// SetupRoutes func
func SetupRoutes(app *fiber.App) {
	// Middleware
	api := app.Group("/api")

	// routes
	api.Post("/users", func(c *fiber.Ctx) error {
		var user models.User

		c.BodyParser(&user)

		coll := db.DB.Database("gomongodb").Collection("users")
		result, err := coll.InsertOne(context.TODO(), bson.D{{
			Key:   "name",
			Value: user.Name,
		}})

		if err != nil {
			panic(err)
		}

		return c.JSON(&fiber.Map{
			"data": result,
		})
	})

	api.Get("/users", func(c *fiber.Ctx) error {
		var users []models.User

		coll := db.DB.Database("gomongodb").Collection("users")
		results, error := coll.Find(context.TODO(), bson.M{})

		if error != nil {
			panic(error)
		}

		for results.Next(context.TODO()) {
			var user models.User
			results.Decode(&user)
			users = append(users, user)
		}

		return c.JSON(&fiber.Map{
			"users": users,
		})

	})
}
