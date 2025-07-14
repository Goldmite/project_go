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
	GroupId     string `json:"group_id"`
	GroupName   string `json:"group_name"`
	InviterName string `json:"inviter_name"`
}

type BookResponse struct {
	OwnedBy []string `json:"owned_by"`
	ISBN    string   `json:"isbn"`
	Title   string   `json:"title"`
	Authors []string `json:"authors"`
	Cover   string   `json:"cover"`
}

type BookInfoResponse struct {
	OwnedBy     []string `json:"owned_by"`
	ISBN        string   `json:"isbn"`
	Title       string   `json:"title"`
	Authors     []string `json:"authors"`
	Pages       uint     `json:"pages"`
	Description string   `json:"description"`
	Publisher   string   `json:"publisher"`
	PublishDate string   `json:"publishedDate"`
	Language    string   `json:"language"`
	Cover       string   `json:"cover"`
}

type BookProgressResponse struct {
	PagesRead   uint `json:"pages_read"`
	TimeRead    uint `json:"time_read"`
	FirstPage   uint `json:"first_page"`
	CurrentPage uint `json:"current_page"`
}
