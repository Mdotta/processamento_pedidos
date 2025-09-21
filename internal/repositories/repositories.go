package repositories

import (
	"database/sql"
	"processamento_pedidos/internal/models"
	"processamento_pedidos/internal/repositories/users"

	"github.com/google/uuid"
)

type Repositories struct {
	User interface {
		GetAll() []models.User
		Add(newUser models.User)
		EmailInUse(email string) bool
		GetWithFilters(id *uuid.UUID, email *string) (*models.User, error)
	}
}

func New(db *sql.DB) *Repositories {
	return &Repositories{
		User: users.New(db),
	}
}
