package usecases

import (
	"processamento_pedidos/internal/repositories"
	users "processamento_pedidos/internal/usecases/user"
)

type UseCases struct {
	User users.UserUseCase //Upercased: exported | lowercased: unexported
}

func New(repos *repositories.Repositories) *UseCases {
	return &UseCases{
		User: *users.New(repos),
	}
}
