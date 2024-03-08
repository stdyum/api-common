package models

import (
	"github.com/google/uuid"
)

type TypesIds struct {
	GroupsIds   []uuid.UUID
	RoomsIds    []uuid.UUID
	SubjectsIds []uuid.UUID
	TeachersIds []uuid.UUID
}

type TypesModels struct {
	GroupsIds   map[uuid.UUID]Group
	RoomsIds    map[uuid.UUID]Room
	SubjectsIds map[uuid.UUID]Subject
	TeachersIds map[uuid.UUID]Teacher
}

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
