package usecases

import (
	"processamento_pedidos/internal/repositories"
	rankingscores "processamento_pedidos/internal/usecases/rankingScores"
	users "processamento_pedidos/internal/usecases/user"
)

type UseCases struct {
	User         users.UserUseCase
	RankingScore rankingscores.RankingScoreUseCase //Upercased: exported | lowercased: unexported
}

func New(repos *repositories.Repositories) *UseCases {
	return &UseCases{
		User:         *users.New(repos),
		RankingScore: *rankingscores.New(repos),
	}
}
