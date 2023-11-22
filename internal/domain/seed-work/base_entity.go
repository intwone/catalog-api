package seedwork

import "github.com/google/uuid"

type Entity struct {
	ID uuid.UUID
}

func NewEntity() *Entity {
	return &Entity{
		ID: uuid.New(),
	}
}
