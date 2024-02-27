package entities

import (
	"time"
)

type ITimed interface {
	GetCreatedAt() time.Time
	GetUpdatedAt() time.Time
}

type Timed struct {
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (t Timed) GetCreatedAt() time.Time {
	return t.CreatedAt
}

func (t Timed) GetUpdatedAt() time.Time {
	return t.UpdatedAt
}
