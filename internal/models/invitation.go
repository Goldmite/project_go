package models

import (
	"time"

	"github.com/Goldmite/project_go/internal/enums"
	"github.com/google/uuid"
)

type Invitation struct {
	Token     string `form:"token"`
	EmailTo   string `form:"email_to"`
	GroupId   string `form:"group_id"`
	InvitedBy string `form:"invited_by"`
	Status    enums.Status
	SentAt    string
	ExpiresAt string
}

type InviteRequest struct {
	EmailTo   string `json:"email_to`
	GroupId   string `json:"group_id"`
	InvitedBy string `json:"invited_by"`
}

func NewInvitationFromRequest(r InviteRequest) *Invitation {
	sentAt := time.Now().String()
	expiresAt := time.Now().Add(time.Hour * 24 * 7).String() // expire after a week

	return &Invitation{
		Token:     uuid.New().String(),
		EmailTo:   r.EmailTo,
		GroupId:   r.GroupId,
		InvitedBy: r.InvitedBy,
		Status:    enums.Pending,
		SentAt:    sentAt,
		ExpiresAt: expiresAt,
	}
}
