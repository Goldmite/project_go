package dto

type GetUserResponse struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type FetchBookResponse struct {
	Title   string   `json:"title"`
	Authors []string `json:"authors"`
}

type InviteResponse struct {
	GroupId   string `json:"group_id"`
	GroupName string `json:"group_name"`
	InvitedBy string `json:"invited_by"`
}
