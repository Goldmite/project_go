package services

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/Goldmite/project_shelf/internal/models"
	"github.com/Goldmite/project_shelf/internal/models/dto"
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

func (groupService *GroupService) GetAllUserGroups(userId string) ([]models.Group, error) {
	query := "SELECT g.id, g.name, g.created_at FROM groups g JOIN members m ON g.id = m.group_id WHERE m.user_id = ?"
	rows, err := groupService.database.Query(query, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var userGroups []models.Group
	for rows.Next() {
		var g models.Group
		err := rows.Scan(&g.Id, &g.Name, &g.CreatedAt)
		if err != nil {
			return nil, err
		}

		userGroups = append(userGroups, g)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return userGroups, nil
}

func (groupService *GroupService) CreateInvite(inv models.Invitation) error {
	query := "INSERT INTO invitations (email_to, group_id, invited_by, status, sent_at, expires_at) VALUES (?, ?, ?, ?, ?, ?)"
	_, err := groupService.database.Exec(query, inv.EmailTo, inv.GroupId, inv.InvitedBy, inv.Status, inv.SentAt, inv.ExpiresAt)
	if err != nil {
		if strings.Contains(err.Error(), "UNIQUE constraint failed") {
			return nil
		}
		return fmt.Errorf("failed to create invite for %s %w", inv.EmailTo, err)
	}

	return nil
}

func (groupService *GroupService) GetInvites(groupId string) ([]dto.GetUserResponse, error) {
	query := "SELECT id, name, email FROM users JOIN invitations ON email = email_to WHERE group_id = ? ORDER BY sent_at ASC"
	rows, err := groupService.database.Query(query, groupId)
	if err != nil {
		return nil, fmt.Errorf("failed to get invites to group %w", err)
	}
	defer rows.Close()

	var invitedUsers []dto.GetUserResponse
	for rows.Next() {
		var user dto.GetUserResponse
		err := rows.Scan(&user.ID, &user.Name, &user.Email)
		if err != nil {
			return nil, err
		}

		invitedUsers = append(invitedUsers, user)
	}

	return invitedUsers, nil
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
