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
