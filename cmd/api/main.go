package main

import (
	"database/sql"
	"log"
	"processamento_pedidos/internal/handlers"
	"processamento_pedidos/internal/repositories"
	"processamento_pedidos/internal/usecases"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

// cadastrar e listar usuários
// handlers <- usecases <- repositories
func main() {
	//create connection to db
	//add host and port
	connStr := "postgresql://dotta:2UqCRIlLqFCcrO46Mi9omsu3gVU1VKrW@dpg-d399qi3e5dus73ao2nl0-a.virginia-postgres.render.com/processamento_pedidos"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	//make sure connection will be closed when main function ends
	defer db.Close()

	initDatabase(db)

	//create repositories, usecases and handlers
	//pass db connection to repositories
	repos := repositories.New(db)
	useCases := usecases.New(repos)
	h := handlers.New(useCases)

	//start http server listening to port 8080
	h.Listen(10000)
}

func initDatabase(db *sql.DB) {
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Fatalf("could not create database driver: %v", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://../../migrations",
		"postgres",
		driver,
	)
	if err != nil {
		log.Fatalf("could not create migrate instance: %v", err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("could not run migrations: %v", err)
	}

	log.Println("Database migrations applied successfully")
}
