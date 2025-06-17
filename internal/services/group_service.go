package services

import (
	"database/sql"
	"fmt"

	"github.com/Goldmite/project_go/internal/models"
	"github.com/Goldmite/project_go/internal/models/dto"
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
		return fmt.Errorf("failed to create invite %w", err)
	}

	return nil
}

func (groupService *GroupService) GetUserInvites(userId string) ([]dto.InviteResponse, error) {
	query :=
		"SELECT i.group_id, g.name, i.invited_by" +
			"FROM invitations i" +
			"JOIN users u ON u.email = i.email_to" +
			"JOIN groups g ON g.id = i.group_id" +
			"WHERE u.id = ?"
	rows, err := groupService.database.Query(query, userId)
	if err != nil {
		return nil, fmt.Errorf("failed to get user invites %w", err)
	}
	defer rows.Close()

	var groupInvites []dto.InviteResponse
	for rows.Next() {
		var inv dto.InviteResponse
		err := rows.Scan(&inv.GroupId, &inv.GroupName, &inv.InvitedBy)
		if err != nil {
			return nil, err
		}
		//todo query inviter to get name and email for fE
		groupInvites = append(groupInvites, inv)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return groupInvites, nil
}
