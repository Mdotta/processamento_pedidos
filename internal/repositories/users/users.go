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
	rows, err := u.db.Query("SELECT id, username, email FROM users")
	if err != nil {
		fmt.Println("error querying users:", err)
		return nil
	}
	defer rows.Close()

	users := []models.User{}
	for rows.Next() {
		var idStr, username, email string
		if err := rows.Scan(&idStr, &username, &email); err != nil {
			fmt.Println("error scanning user:", err)
			continue
		}
		id, _ := uuid.Parse(idStr)
		users = append(users, models.User{ID: id, Username: username, Email: email})
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
	_, err := u.db.Exec("INSERT INTO users (id, username, email,password_hash,created_at,updated_at) VALUES ($1, $2, $3,$4,now(),now())",
		newUser.ID, newUser.Username, newUser.Email, newUser.PasswordHash)
	if err != nil {
		fmt.Println("error adding user:", err)
	}
}

func (u Users) GetWithFilters(id *uuid.UUID, email *string) (*models.User, error) {
	query := "SELECT * FROM users"
	hasFiltersAlready := false
	if id != nil {
		query += fmt.Sprintf("WHERE id='%s'", id.String())
		hasFiltersAlready = true
	}
	if email != nil {
		if hasFiltersAlready {
			query += " AND "
		} else {
			query += " WHERE "
		}
		query += fmt.Sprintf("email='%s'", *email)
	}

	rows, err := u.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var user models.User
	if rows.Next() {
		err := rows.Scan(&user.ID, &user.Username, &user.Email)
		if err != nil {
			return nil, err
		}
	}
	return &user, err
}
