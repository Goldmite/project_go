package dto

type UserBookRequest struct {
	UserId string `json:"userId"`
	Isbn   string `json:"isbn"`
}
