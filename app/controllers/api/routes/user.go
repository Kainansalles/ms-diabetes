package routes

import (
	"github.com/Kainansalles/ms-diabetes/app/controllers/api/handlers"
	"github.com/gofiber/fiber/v2"
)

func Users(app *fiber.App) {
	UserHandler := handlers.UserHandler{}

	app.Post("/user", UserHandler.SaveUser)
}
