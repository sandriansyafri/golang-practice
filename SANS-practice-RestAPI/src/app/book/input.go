package book

type BookRequest struct {
	ID     string `json:"id"`
	Title  string `json:"title" binding:"required"`
	Price  int    `json:"price" binding:"required"`
	Rating int    `json:"rating" binding:"required"`
}
