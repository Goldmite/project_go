package models

import "github.com/Goldmite/project_go/internal/enums"

type Invitation struct {
	token     string `form:"token"`
	emailTo   string `form:"email_to"`
	groupId   string `form:"group_id"`
	invitedBy string `form:"invited_by"`
	status    enums.Status
	sentAt    string
	expiresAt string
}

type InviteRequest struct {
	token     string `json:"token"`
	groupId   string `json:"group_id"`
	invitedBy string `json:"invited_by"`
}
