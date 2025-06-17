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
	Token       string `json:"token"`
	GroupId     string `json:"group_id"`
	GroupName   string `json:"group_name"`
	InviterName string `json:"inviter_name"`
}
