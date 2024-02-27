package models

import (
	"github.com/google/uuid"
)

type Group struct {
	ID   uuid.UUID
	Name string
}

type Room struct {
	ID   uuid.UUID
	Name string
}

type Subject struct {
	ID   uuid.UUID
	Name string
}

type Teacher struct {
	ID   uuid.UUID
	Name string
}
