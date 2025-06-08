package dto

import (
	"github.com/Goldmite/project_go/internal/models"
)

type UserBookRequest struct {
	UserId string `json:"user_id"`
	Book   models.Book
}
