package main

import (
	"github.com/Kainansalles/ms-diabetes/app/controllers/api/routes"
	"github.com/Kainansalles/ms-diabetes/config"
	"github.com/gofiber/fiber/v2"
)

type User struct {
	Name string `validate:"required,min=3,max=32"`
	Cpf  int    `validate:"required,lenght=11"`
}

func main() {
	config.InitEnvConfig()
	app := fiber.New()

	routes.Users(app)

	app.Listen(":3000")
}
