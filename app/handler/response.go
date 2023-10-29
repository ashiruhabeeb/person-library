package handler

import "time"

// BookResponse contains response body of the book API
type BookResponse struct {
	BooksId     uint    `json:"bookid"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Author      string  `json:"author"`
	Acquisition string  `json:"acquisition"`
	Price       float32 `json:"price"`
	Rating      uint    `json:"rating"`
	Discount    float32	`json:"discount"`
	Quanitity	uint	`json:"quantity"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
