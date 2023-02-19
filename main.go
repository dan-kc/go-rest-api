package main

import (
	"github.com/dan-kc/go-rest-api/packages/controllers"
	"github.com/dan-kc/go-rest-api/packages/initializers"
	"github.com/gofiber/fiber/v2"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDatabase()
	initializers.SyncDatabase()
}

func main() {

	// setup app
	app := fiber.New()

	// routes
	app.Get("/posts", controllers.GetAllPosts)
	app.Get("/post/:id", controllers.GetPost)
	app.Post("/post", controllers.CreatePost)
	app.Put("/post/:id", controllers.UpdatePost)
	app.Delete("/post/:id", controllers.DeletePost)
	app.Get("/healthcheck", controllers.CheckHealth)

	// start app
	app.Listen(":3000")
}
