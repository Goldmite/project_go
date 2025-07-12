package services

import (
	"database/sql"

	"github.com/Goldmite/project_go/internal/models/dto"
)

type StatsService struct {
	database *sql.DB
}

func NewStatsService(db *sql.DB) *StatsService {
	return &StatsService{database: db}
}

func (statsService *StatsService) UpdateBookProgress(req dto.BookProgressRequest) error {
	var (
		query string
		args  []any
	)

	if req.FirstPage != nil {
		query = "UPDATE reading SET pages_read = ?, time_read = ?, first_page = ?, session_created_at = CURRENT_TIMESTAMP WHERE user_id = ? AND book_id = ?"
		args = []any{req.PagesRead, req.TimeRead, *req.FirstPage, req.UserId, req.Isbn}
	} else {
		query = "UPDATE reading SET pages_read = ?, time_read = ?, session_updated_at = CURRENT_TIMESTAMP WHERE user_id = ? AND book_id = ?"
		args = []any{req.PagesRead, req.TimeRead, req.UserId, req.Isbn}
	}
	_, err := statsService.database.Exec(query, args...)
	if err != nil {
		return err
	}
	return nil
}

func (statsService *StatsService) GetBookProgress(userId, isbn string) (*dto.BookProgressResponse, error) {
	query := "SELECT pages_read, time_read, first_page FROM reading WHERE user_id = ? AND book_id = ?"
	row := statsService.database.QueryRow(query, userId, isbn)

	var progress dto.BookProgressResponse
	err := row.Scan(&progress.PagesRead, &progress.TimeRead, &progress.FirstPage)
	if err != nil {
		return nil, err
	}

	return &progress, nil
}
