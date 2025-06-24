package models

import (
	"time"

	"github.com/google/uuid"
)

type Group struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
}

func NewGroupFromName(name string) *Group {
	createdAt := time.Now().String()
	return &Group{
		Id:        uuid.New().String(),
		Name:      name,
		CreatedAt: createdAt,
	}
}
