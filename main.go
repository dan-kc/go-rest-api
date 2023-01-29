package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/dan-kc/go-rest-api/packages/initializers"
	"github.com/dan-kc/go-rest-api/packages/controllers"
	"os"
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
	app.Get("/post/:id", controllers.GetPost)
	app.Get("/posts", controllers.GetAllPosts)
	app.Post("/post", controllers.CreatePost)
  app.Put("/post/:id", controllers.UpdatePost)
  app.Delete("/post/:id", controllers.DeletePost)

	// start app
	app.Listen(":" + os.Getenv("PORT"))

}
