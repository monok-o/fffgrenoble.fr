package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/robjullian/fffgrenoble.fr/controllers"
	"github.com/robjullian/fffgrenoble.fr/middlewares"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	api.Use(logger.New()) // Add logger

	// Public endpoints
	api.Get("/events", controllers.GetEvents)
	api.Get("/links", controllers.GetLinks)
	api.Get("/socials", controllers.GetSocials)

	// Event api
	api.Post("/event", middlewares.CheckAuth, controllers.AddEvent)
	api.Put("/event/:id", middlewares.CheckAuth, controllers.UpdateEvent)
	api.Delete("/event/:id", middlewares.CheckAuth, controllers.DeleteEvent)

	// Link api
	api.Post("/link", middlewares.CheckAuth, controllers.AddLink)
	api.Put("/link/:id", middlewares.CheckAuth, controllers.UpdateLink)
	api.Delete("/link/:id", middlewares.CheckAuth, controllers.DeleteLink)

	// Social api
	api.Post("/social", middlewares.CheckAuth, controllers.AddSocial)
	api.Put("/social/:id", middlewares.CheckAuth, controllers.UpdateSocial)
	api.Delete("/social/:id", middlewares.CheckAuth, controllers.DeleteSocial)
}
