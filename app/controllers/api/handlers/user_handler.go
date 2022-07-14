package handlers

import (
	"github.com/Kainansalles/ms-diabetes/app/controllers/api/requests"
	"github.com/Kainansalles/ms-diabetes/app/infra/gorm/mysql"
	"github.com/Kainansalles/ms-diabetes/pkg/structs"
	"github.com/gofiber/fiber/v2"
)

// PaymentHandler type.
type UserHandler struct {
}

func (handler UserHandler) SaveUser(ctx *fiber.Ctx) error {
	user := new(requests.User)

	if err := ctx.BodyParser(user); err != nil {
		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(&fiber.Map{
			"status":  "error",
			"message": err.Error(),
		})
	}

	if err := structs.ValidateStruct(user); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(err)
	}

	mysql.Connect()

	// mysql.Database.Create(user) // pass pointer of data to Create

	isSaved := mysql.Database.Create(user)

	if isSaved.Error != nil {
		return ctx.JSON(&fiber.Map{
			"status":  false,
			"message": "Erro ao salvar usuário",
		})
	}

	// regra para salvar

	return ctx.JSON(&fiber.Map{
		"status":  "success",
		"message": "Usuário salvo com sucesso",
		"data":    nil,
	})
}
