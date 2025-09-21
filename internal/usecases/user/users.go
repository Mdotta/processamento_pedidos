package users

import (
	"errors"
	"log/slog"
	"processamento_pedidos/internal/models"
	"processamento_pedidos/internal/repositories"
	"processamento_pedidos/internal/repositories/users"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserUseCase struct {
	repo users.Users
}

func New(repo *repositories.Repositories) *UserUseCase {
	return &UserUseCase{repo: *repo.User.(*users.Users)}
}

func (u UserUseCase) GetAll() []models.User {
	users := u.repo.GetAll()
	return users
}

func (u UserUseCase) GetById(userId *uuid.UUID, email *string) (*models.User, error) {
	user, err := u.repo.GetWithFilters(userId, email)
	if err != nil || user == nil {
		slog.Error("error getting user by id", "id", userId, "error", err)
		return nil, errors.New("user not found")
	}
	return user, nil
}

func (u UserUseCase) Add(newUser models.CreateUserRequest) (uuid.UUID, error) {

	exists := u.repo.EmailInUse(newUser.Email)
	if exists {
		slog.Error("email already in use", "email", newUser.Email)
		return uuid.Nil, errors.New("email already in use")
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
	if err != nil {
		slog.Error("error hashing password", "error", err)
		return uuid.Nil, errors.New("failed to hash password")
	}

	repoReq := models.User{
		ID:           uuid.New(),
		Username:     newUser.Username,
		Email:        newUser.Email,
		PasswordHash: string(hashedPassword),
	}

	u.repo.Add(repoReq)

	return repoReq.ID, nil
}
