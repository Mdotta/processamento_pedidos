package main

import (
	"database/sql"
	"log"
	"processamento_pedidos/internal/handlers"
	"processamento_pedidos/internal/repositories"
	"processamento_pedidos/internal/usecases"
)

// cadastrar e listar usuários
// handlers <- usecases <- repositories
func main() {
	//create connection to db
	connStr := "user=dotta password=safepass dbname=processamento_pedidos sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	//make sure connection will be closed when main function ends
	defer db.Close()

	//create repositories, usecases and handlers
	//pass db connection to repositories
	repos := repositories.New(db)
	useCases := usecases.New(repos)
	h := handlers.New(useCases)

	//start http server listening to port 8080
	h.Listen(8080)
}
