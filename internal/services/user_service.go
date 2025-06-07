package services

import (
	"database/sql"
	"fmt"

	"github.com/Goldmite/project_go/internal/models"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	database *sql.DB
}

func NewUserService(db *sql.DB) *UserService {
	return &UserService{database: db}
}

func CheckHash(hashed, plain string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(plain))
	if err != nil {
		return fmt.Errorf("invalid password")
	}
	return nil
}

func (userService *UserService) CreateUser(r models.CreateUserRequest) (*string, error) {
	u, err := models.NewUserFromRequest(r)
	if err != nil {
		return nil, fmt.Errorf("failed to hash password %w", err)
	}
	query := "INSERT INTO users (id, name, email, pw_hash, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?)"
	_, err = userService.database.Exec(query, u.ID, u.Name, u.Email, u.PasswordHash, u.CreatedAt, u.UpdatedAt)
	if err != nil {
		return nil, fmt.Errorf("failed to create user %w", err)
	}

	return &u.ID, nil
}

func (userService *UserService) GetUserById(id string) (*models.GetUserResponse, error) {
	query := "SELECT id, name, email FROM users WHERE id = ?"

	var res models.GetUserResponse
	err := userService.database.QueryRow(query, id).Scan(&res.ID, &res.Name, &res.Email)
	if err != nil {
		return nil, fmt.Errorf("user not found %w", err)
	}

	return &res, nil
}

func (userService *UserService) GetUserByLogin(email, password string) (*models.GetUserResponse, error) {
	query := "SELECT id, name, email, pw_hash FROM users WHERE email = ?"

	var res models.GetUserResponse
	var hashed string
	err := userService.database.QueryRow(query, email).Scan(&res.ID, &res.Name, &res.Email, &hashed)
	if err != nil {
		return nil, fmt.Errorf("user not found %w", err)
	}

	if err := CheckHash(hashed, password); err != nil {
		return nil, fmt.Errorf("invalid password")
	}

	return &res, nil
}
