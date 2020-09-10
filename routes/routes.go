package routes

import (
	"server/handlers"

	"github.com/gofiber/fiber"
)

// CreateRoutes adds all the routes to the fiber app
func CreateRoutes(app *fiber.App) {
	app.Get("/posts", handlers.GetPosts)
	app.Post("/post", handlers.CreatePost)
	app.Get("post/:id", handlers.GetPostByID)

	app.Post("/signup", handlers.CreateUser)
	app.Post("/login", handlers.Login)
}
