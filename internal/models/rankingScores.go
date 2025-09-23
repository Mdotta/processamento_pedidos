package models

import "github.com/google/uuid"

type RankingScore struct {
	ID       uuid.UUID
	UserID   uuid.UUID //user id
	BestTime int
}

type CreateRankingScoreRequest struct {
	UserID   uuid.UUID `json:"user_id"`
	BestTime int       `json:"best_time"`
}

type CreateRankingScoreResponse struct {
	RankingId uuid.UUID `json:"ranking_id"`
}
