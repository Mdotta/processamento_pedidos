package rankingscores

import (
	"fmt"
	"processamento_pedidos/internal/models"
	"processamento_pedidos/internal/repositories"
	"processamento_pedidos/internal/repositories/rankingScores"
	"processamento_pedidos/internal/repositories/users"

	"github.com/google/uuid"
)

type RankingScoreUseCase struct {
	repo     rankingScores.RankingScores
	userRepo users.Users
}

func New(repo *repositories.Repositories) *RankingScoreUseCase {
	return &RankingScoreUseCase{
		repo:     *repo.RankingScore.(*rankingScores.RankingScores),
		userRepo: *repo.User.(*users.Users),
	}
}

func (r RankingScoreUseCase) GetAll() []models.RankingScore {
	rankings := r.repo.GetAll()
	return rankings
}

func (r RankingScoreUseCase) Add(newRanking models.CreateRankingScoreRequest) (uuid.UUID, error) {

	existingRanking, err := r.repo.GetByUserID(newRanking.UserID)
	if err == nil && existingRanking != nil {
		fmt.Println("ranking score for user already exists, updating instead")
		// If the new best time is better (lower) than the existing one, update it
		if newRanking.BestTime < existingRanking.BestTime {
			existingRanking.BestTime = newRanking.BestTime
			rows, err := r.repo.Update(*existingRanking)
			if err != nil {
				return uuid.Nil, err
			}
			if rows == 0 {
				return uuid.Nil, fmt.Errorf("no rows updated")
			} else {
				fmt.Println("ranking score updated successfully")
			}
		}
		return existingRanking.ID, nil
	}

	userExist := r.userRepo.UserIdExists(newRanking.UserID)
	if !userExist {
		fmt.Println("user does not exist, cannot add ranking score")
		return uuid.Nil, fmt.Errorf("user does not exist")
	}

	newRecord := models.RankingScore{
		ID:       uuid.New(),
		UserID:   newRanking.UserID,
		BestTime: newRanking.BestTime,
	}
	// TODO: add error handling
	r.repo.Add(newRecord)
	return newRecord.ID, nil

}
