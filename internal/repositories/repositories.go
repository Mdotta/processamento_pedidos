package repositories

import (
	"processamento_pedidos/cmd/internal/models"
	"processamento_pedidos/cmd/internal/repositories/users"
)

type Repositories struct {
	User interface {
		GetAll() []models.User
		Add(newUser models.User)
		EmailInUse(email string) bool
	}
}

func New() *Repositories {
	return &Repositories{
		User: users.New(),
	}
}
