package repositories

import (
	"database/sql"
	"processamento_pedidos/internal/models"
	"processamento_pedidos/internal/repositories/rankingScores"
	"processamento_pedidos/internal/repositories/users"

	"github.com/google/uuid"
)

type Repositories struct {
	User interface {
		GetAll() []models.User
		Add(newUser models.User)
		EmailInUse(email string) bool
		GetWithFilters(id *uuid.UUID, email *string) (*models.User, error)
		UserIdExists(ID uuid.UUID) bool
	}
	RankingScore interface {
		GetAll() []models.RankingScore
		Add(newRanking models.RankingScore)
		Update(updatedRanking models.RankingScore) (int64, error)
		GetByUserID(userID uuid.UUID) (*models.RankingScore, error)
		UserExistInRanking(userID uuid.UUID) bool
	}
}

func New(db *sql.DB) *Repositories {
	return &Repositories{
		User:         users.New(db),
		RankingScore: rankingScores.New(db),
	}
}
