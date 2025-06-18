package services

import (
	"database/sql"
	"fmt"

	"github.com/Goldmite/project_go/internal/models"
)

type GroupService struct {
	database *sql.DB
}

func NewGroupService(db *sql.DB) *GroupService {
	return &GroupService{database: db}
}

func (groupService *GroupService) CreateGroup(name string) (*string, error) {
	g := models.NewGroupFromName(name)
	query := "INSERT INTO groups (id, name, created_at) VALUES (?, ?, ?)"
	_, err := groupService.database.Exec(query, g.Id, g.Name, g.CreatedAt)
	if err != nil {
		return nil, fmt.Errorf("failed to create group %w", err)
	}

	return &g.Id, nil
}

func (groupService *GroupService) CreateInvite(inv models.Invitation) error {
	query := "INSERT INTO invitations (token, email_to, group_id, invited_by, status, sent_at, expires_at) VALUES (?, ?, ?, ?, ?, ?, ?)"
	_, err := groupService.database.Exec(query, inv.Token, inv.EmailTo, inv.GroupId, inv.InvitedBy, inv.Status, inv.SentAt, inv.ExpiresAt)
	if err != nil {
		return fmt.Errorf("failed to create invite for %s %w", inv.EmailTo, err)
	}

	return nil
}
