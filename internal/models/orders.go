package models

import "github.com/google/uuid"

type Order struct {
	ID      uuid.UUID
	Name    string
	OwnerID uuid.UUID //user id
}
