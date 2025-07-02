package dto

type UserBookRequest struct {
	UserId string `json:"userId" binding:"required"`
	Isbn   string `json:"isbn" binding:"required"`
}

type InviteRequest struct {
	EmailTo   []string `form:"email_to" binding:"required"`
	GroupId   string   `form:"group_id" binding:"required"`
	InvitedBy string   `form:"invited_by" binding:"required"`
}

type AcceptRequest struct {
	UserId    string `json:"user_id" binding:"required"`
	UserEmail string `json:"email" binding:"required"`
	GroupId   string `json:"group_id" binding:"required"`
}

type DeclineRequest struct {
	UserEmail string `json:"email" binding:"required"`
	GroupId   string `json:"group_id" binding:"required"`
}
