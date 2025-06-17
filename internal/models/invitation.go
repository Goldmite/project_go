package models

import (
	"time"

	"github.com/Goldmite/project_go/internal/enums"
	"github.com/Goldmite/project_go/internal/models/dto"
	"github.com/google/uuid"
)

type Invitation struct {
	Token     string
	EmailTo   string
	GroupId   string
	InvitedBy string
	Status    enums.Status
	SentAt    string
	ExpiresAt string
}

func NewInvitationFromRequest(r dto.InviteRequest, inviteNr int) *Invitation {
	sentAt := time.Now().String()
	expiresAt := time.Now().Add(time.Hour * 24 * 7).String() // expire after a week

	return &Invitation{
		Token:     uuid.New().String(),
		EmailTo:   r.EmailTo[inviteNr],
		GroupId:   r.GroupId,
		InvitedBy: r.InvitedBy,
		Status:    enums.Pending,
		SentAt:    sentAt,
		ExpiresAt: expiresAt,
	}
}
