package dto

type MovieResponse struct {
	ID          int    `gorm:"primary_key;not_null" json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	ReleaseDate string `json:"release_date"`
	PosterUrl   string `json:"poster_url"`
	AgeRating   string `json:"age_rating"`
	TicketPrice int    `json:"ticket_price"`
}