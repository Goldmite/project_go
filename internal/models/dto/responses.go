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
