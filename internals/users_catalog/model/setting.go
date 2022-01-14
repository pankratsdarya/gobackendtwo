package model

import "github.com/google/uuid"

type Setting struct {
	ID          uuid.UUID
	Name        string
	Type        string
	Contact     uuid.UUID
	Description string
}
