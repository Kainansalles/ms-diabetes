package requests

type User struct {
	Name string `json:"name" validate:"required" example:"Kainan Salles"`
	CPF  int    `json:"cpf" validate:"required" example:"99999999999"`
}
