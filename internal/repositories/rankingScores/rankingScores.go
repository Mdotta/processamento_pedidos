package rankingScores

import (
	"database/sql"
	"fmt"
	"processamento_pedidos/internal/models"

	"github.com/google/uuid"
)

type RankingScores struct {
	db *sql.DB
}

func New(db *sql.DB) *RankingScores {
	return &RankingScores{db: db}
}

func (u RankingScores) GetAll() []models.RankingScore {
	rows, err := u.db.Query("SELECT id, user_id, best_time FROM ranking_scores")
	if err != nil {
		fmt.Println("error querying users:", err)
		return nil
	}
	defer rows.Close()

	users := []models.RankingScore{}
	for rows.Next() {
		var idStr, userIdStr string
		var bestTime int
		if err := rows.Scan(&idStr, &userIdStr, &bestTime); err != nil {
			fmt.Println("error scanning user:", err)
			continue
		}
		id, _ := uuid.Parse(idStr)
		userId, _ := uuid.Parse(userIdStr)
		users = append(users, models.RankingScore{ID: id, UserID: userId, BestTime: bestTime})
	}
	return users
}

func (r RankingScores) Add(newRanking models.RankingScore) {

	//if there's a data with the same UserID, update it instead of adding a new one
	_, err := r.db.Exec("INSERT INTO ranking_scores (id, user_id, best_time,created_at,updated_at) VALUES ($1, $2, $3,now(),now())",
		newRanking.ID, newRanking.UserID.String(), newRanking.BestTime)
	if err != nil {
		fmt.Println("error adding ranking score:", err)
	}
}

func (r RankingScores) Update(updatedRanking models.RankingScore) (int64, error) {

	result, err := r.db.Exec("UPDATE ranking_scores SET best_time=$1, updated_at=now() WHERE user_id=$2",
		updatedRanking.BestTime, updatedRanking.UserID)
	if err != nil {
		fmt.Println("error updating ranking score:", err)
		return 0, err
	}
	rows, _ := result.RowsAffected()
	return rows, nil
}

func (r RankingScores) UserExistInRanking(userID uuid.UUID) bool {
	var exists bool
	err := r.db.QueryRow("SELECT EXISTS(SELECT 1 FROM ranking_scores WHERE user_id=$1)", userID).Scan(&exists)
	if err != nil {
		fmt.Println("error checking user in ranking:", err)
		return false
	}
	return exists
}

func (r RankingScores) GetByUserID(userID uuid.UUID) (*models.RankingScore, error) {
	row := r.db.QueryRow("SELECT id, user_id, best_time FROM ranking_scores WHERE user_id=$1", userID)
	var idStr, userIdStr string
	var bestTime int
	if err := row.Scan(&idStr, &userIdStr, &bestTime); err != nil {
		return nil, err
	}
	id, _ := uuid.Parse(idStr)
	userId, _ := uuid.Parse(userIdStr)
	return &models.RankingScore{ID: id, UserID: userId, BestTime: bestTime}, nil
}
