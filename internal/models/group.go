package models

import (
	"time"

	"github.com/google/uuid"
)

type Group struct {
	Id        string
	Name      string
	CreatedAt string
}

func NewGroupFromName(name string) *Group {
	createdAt := time.Now().String()
	return &Group{
		Id:        uuid.New().String(),
		Name:      name,
		CreatedAt: createdAt,
	}
}
