package main

import (
	"processamento_pedidos/internal/handlers"
	"processamento_pedidos/internal/repositories"
	"processamento_pedidos/internal/usecases"
)

// cadastrar e listar usuários
// handlers <- usecases <- repositories
func main() {
	repos := repositories.New()
	useCases := usecases.New(repos)
	h := handlers.New(useCases)

	h.Listen(8080)
}
