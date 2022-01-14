package model

import (
	"github.com/google/uuid"
)

type User struct {
	ID       uuid.UUID
	Name     string
	Surname  string
	Position string
	Email    string
}
