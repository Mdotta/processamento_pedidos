package usecases

import (
	"errors"
	"log/slog"
	"processamento_pedidos/internal/models"
	"processamento_pedidos/internal/repositories"

	"github.com/google/uuid"
)

type UseCases struct {
	repos repositories.Repositories
}

func New(repos *repositories.Repositories) *UseCases {
	return &UseCases{
		repos: *repos,
	}
}

func (u UseCases) GetAll() []models.User {
	users := u.repos.User.GetAll()
	return users
}
func (u UseCases) Add(newUser models.CreateUserRequest) (uuid.UUID, error) {

	exists := u.repos.User.EmailInUse(newUser.Email)
	if exists {
		slog.Error("Email already in use", "email", newUser.Email)
		return uuid.Nil, errors.New("Email already in use")
	}
	repoReq := models.User{
		ID:    uuid.New(),
		Name:  newUser.Name,
		Email: newUser.Email,
	}

	u.repos.User.Add(repoReq)

	return repoReq.ID, nil
}
