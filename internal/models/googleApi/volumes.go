package googleapi

type ImageLinks struct {
	Thumbnail string `json:"thumbnail"`
}

type VolumeInfo struct {
	Title         string     `json:"title"`
	Authors       []string   `json:"authors"`
	Publisher     string     `json:"publisher"`
	PublishedDate string     `json:"publishedDate"`
	Description   string     `json:"description"`
	PageCount     int        `json:"pageCount"`
	ImageLinks    ImageLinks `json:"imageLinks"`
	Language      string     `json:"language"`
}
