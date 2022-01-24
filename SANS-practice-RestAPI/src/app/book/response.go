package book

import "time"

type BookResponse struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Price     int       `json:"price"`
	Rating    int       `json:"rating"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
