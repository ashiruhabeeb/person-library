package model

import (
	"encoding/json"
	"time"
)

// Books struct for books table
type Books struct {
	BooksId     uint    	`gorm:"primaryKey;autoIncrement"`
	Title       string 		`gorm:"type:varchar(100);not null"`
	Description string  	`gorm:"type:text;not null"`
	Author      string  	`gorm:"type:text;not null"`
	Acquisition string  	`gorm:"type:text;not null"`
	Price       float32 	`gorm:"not null"`
	Rating      float32 	`gorm:"not null"`
	Discount    float32		`gorm:"null"`
	Quantity	uint		`gorm:"not null"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// BookPayload struct represents book request body for the book api
type BookPayload struct {
	Title       string			`json:"title" binding:"required"`
	Description string 			`json:"description" binding:"required"`
	Author      string			`json:"author" binding:"required"`
	Acquisition string			`json:"acquisition" binding:"required"`
	Price       json.Number		`json:"price" binding:"required,number"`
	Rating      json.Number		`json:"rating" binding:"required,number"`
	Discount	json.Number		`json:"discount" binding:"required,number"`
	Quantity	json.Number		`json:"quantity" binding:"required,number"`
}
