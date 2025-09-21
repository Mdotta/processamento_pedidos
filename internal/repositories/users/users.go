package users

import (
	"database/sql"
	"fmt"
	"processamento_pedidos/internal/models"

	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

type Users struct {
	db *sql.DB
}

func New(db *sql.DB) *Users {
	return &Users{db: db}
}

func (u Users) GetAll() []models.User {
	rows, err := u.db.Query("SELECT id, name, email FROM users")
	if err != nil {
		fmt.Println("error querying users:", err)
		return nil
	}
	defer rows.Close()

	users := []models.User{}
	for rows.Next() {
		var idStr, name, email string
		if err := rows.Scan(&idStr, &name, &email); err != nil {
			fmt.Println("error scanning user:", err)
			continue
		}
		id, _ := uuid.Parse(idStr)
		users = append(users, models.User{ID: id, Name: name, Email: email})
	}
	return users
}

func (u Users) EmailInUse(email string) bool {
	var exists bool
	err := u.db.QueryRow("SELECT EXISTS(SELECT 1 FROM users WHERE email=$1)", email).Scan(&exists)
	if err != nil {
		fmt.Println("error checking email:", err)
		return false
	}
	return exists
}

func (u Users) Add(newUser models.User) {
	_, err := u.db.Exec("INSERT INTO users (id, name, email) VALUES ($1, $2, $3)", newUser.ID.String(), newUser.Name, newUser.Email)
	if err != nil {
		fmt.Println("error adding user:", err)
	}
}

func (u Users) GetByID(id uuid.UUID) (*models.User, error) {
	query := "SELECT * FROM users WHERE id=$1"
	rows, err := u.db.Query(query, id.String())
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var user models.User
	if rows.Next() {
		err := rows.Scan(&user.ID, &user.Name, &user.Email)
		if err != nil {
			return nil, err
		}
	}
	return &user, err
}
