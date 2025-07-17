package services

import (
	"database/sql"
	"time"

	"github.com/Goldmite/project_shelf/internal/models/dto"
)

type StatsService struct {
	database *sql.DB
}

func NewStatsService(db *sql.DB) *StatsService {
	return &StatsService{database: db}
}

func (statsService *StatsService) UpdateBookProgress(req dto.BookProgressRequest) error {
	tx, err := statsService.database.Begin()
	if err != nil {
		return err
	}
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p)
		} else if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()
	var (
		query string
		args  []any
	)

	if req.FirstPage != nil {
		query = "UPDATE reading SET pages_read = ?, time_read = ?, first_page = ?, current_page = ?, session_created_at = CURRENT_TIMESTAMP WHERE user_id = ? AND book_id = ?"
		args = []any{req.PagesRead, req.TimeRead, *req.FirstPage, req.CurrentPage, req.UserId, req.Isbn}
	} else {
		query = "UPDATE reading SET pages_read = ?, time_read = ?, current_page = ?, session_updated_at = CURRENT_TIMESTAMP WHERE user_id = ? AND book_id = ?"
		args = []any{req.PagesRead, req.TimeRead, req.CurrentPage, req.UserId, req.Isbn}
	}
	_, err = statsService.database.Exec(query, args...)
	if err != nil {
		return err
	}
	// Save reading stats:
	err = statsService.UpdateUserTotalStats(req.UserId, req.PagesRead, req.TimeRead)
	if err != nil {
		return err
	}
	err = statsService.AddSessionEntry(req.UserId, req.TimeRead)
	if err != nil {
		return err
	}

	return nil
}

func (statsService *StatsService) GetBookProgress(userId, isbn string) (*dto.BookProgressResponse, error) {
	query := "SELECT pages_read, time_read, first_page, current_page FROM reading WHERE user_id = ? AND book_id = ?"
	row := statsService.database.QueryRow(query, userId, isbn)

	var progress dto.BookProgressResponse
	err := row.Scan(&progress.PagesRead, &progress.TimeRead, &progress.FirstPage, &progress.CurrentPage)
	if err != nil {
		return nil, err
	}

	return &progress, nil
}

func (statsService *StatsService) GetUserStats(userId string) (*dto.TotalProgressResponse, error) {
	query := "SELECT total_pages, total_time FROM stats WHERE user_id = ?"
	row := statsService.database.QueryRow(query, userId)

	var stats dto.TotalProgressResponse
	err := row.Scan(&stats.TotalPagesRead, &stats.TotalTimeRead)
	if err != nil {
		return nil, err
	}

	return &stats, nil
}

func (statsService *StatsService) GetUserSessions(userId, fromDate string) ([]dto.ReadingSession, error) {
	query := "SELECT date, time_read FROM sessions WHERE user_id = ? AND date >= ?"
	rows, err := statsService.database.Query(query, userId, fromDate)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var entries []dto.ReadingSession
	for rows.Next() {
		var e dto.ReadingSession
		err := rows.Scan(&e.Date, &e.TimeRead)
		if err != nil {
			return nil, err
		}
		entries = append(entries, e)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return entries, nil
}

func (statsService *StatsService) UpdateUserTotalStats(userId string, pages, time uint) error {
	query := "UPDATE stats SET total_pages = total_pages + ?, total_time = total_time + ?, updated_at = CURRENT_TIMESTAMP WHERE user_id = ?"
	result, err := statsService.database.Exec(query, pages, time, userId)
	if err != nil {
		return err
	}
	if row, err := result.RowsAffected(); row == 0 {
		if err != nil {
			return err
		}
		query = "INSERT INTO stats (user_id, total_pages, total_time, created_at, updated_at) VALUES(?, ?, ?, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)"
		_, err := statsService.database.Exec(query, userId, pages, time)
		if err != nil {
			return err
		}
	}
	return nil
}

func (statsService *StatsService) AddSessionEntry(userId string, timeRead uint) error {
	query := "INSERT INTO sessions (user_id, date, time_read) VALUES(?, ?, ?) ON DUPLICATE KEY UPDATE time_read = time_read + VALUES(time_read)"
	dateTime := time.Now().Format("2006-01-02")
	_, err := statsService.database.Exec(query, userId, dateTime, timeRead)
	if err != nil {
		return err
	}
	return nil
}
