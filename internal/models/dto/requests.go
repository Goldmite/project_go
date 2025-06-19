package dto

type UserBookRequest struct {
	UserId string `json:"userId"`
	Isbn   string `json:"isbn"`
}

type InviteRequest struct {
	EmailTo   []string `form:"email_to"`
	GroupId   string   `form:"group_id"`
	InvitedBy string   `form:"invited_by"`
}

type AcceptRequest struct {
	UserId    string `json:"user_id"`
	UserEmail string `json:"email"`
	GroupId   string `json:"group_id"`
}

type DeclineRequest struct {
	UserEmail string `json:"email"`
	GroupId   string `json:"group_id"`
}
