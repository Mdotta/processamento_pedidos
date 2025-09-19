package users

import (
	"fmt"
	"processamento_pedidos/cmd/internal/models"
)

type Users struct {
	users []models.User
}

func New() *Users {
	return &Users{users: make([]models.User, 0)}
}

func (u Users) GetAll() []models.User {
	return u.users
}

func (u Users) EmailInUse(email string) bool {
	for _, v := range u.users {
		if v.Email == email {
			return true
		}
	}
	return false
}

func (u *Users) Add(newUser models.User) {
	fmt.Println("adding user", "user", newUser)
	u.users = append(u.users, newUser)
}
