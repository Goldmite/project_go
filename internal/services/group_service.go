package services

import (
	"database/sql"
	"fmt"
	"time"

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
	query := "INSERT INTO invitations (email_to, group_id, invited_by, status, sent_at, expires_at) VALUES (?, ?, ?, ?, ?, ?)"
	_, err := groupService.database.Exec(query, inv.EmailTo, inv.GroupId, inv.InvitedBy, inv.Status, inv.SentAt, inv.ExpiresAt)
	if err != nil {
		return fmt.Errorf("failed to create invite for %s %w", inv.EmailTo, err)
	}

	return nil
}

func (groupService *GroupService) JoinGroup(userId, groupId string) error {
	query := "INSERT INTO members (user_id, group_id, joined_at) VALUES (?, ?, ?)"
	_, err := groupService.database.Exec(query, userId, groupId, time.Now().String())
	if err != nil {
		return fmt.Errorf("failed to create member for group: %w", err)
	}

	return nil
}

func (groupService *GroupService) RemoveInvite(email, groupId string) error {
	query := "DELETE FROM invitations WHERE email_to = ? AND group_id = ?"
	_, err := groupService.database.Exec(query, email, groupId)
	if err != nil {
		return fmt.Errorf("failed to delete invitation: %w", err)
	}

	return nil
}
