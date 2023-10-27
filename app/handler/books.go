package handler

import (
	"github.com/ashiruhabeeb/simple-library/app/services"
	"github.com/gofiber/fiber/v2"
)

// bookHandler implements the services.Bookservice struct
type bookHandler struct {
	bs services.BookService
}

// NewBookHandler initializes a new bookhandler instance
func NewBookHandler(bs services.BookService) *bookHandler {
	return &bookHandler{bs}
} 

func (h *bookHandler) CreateBook(c *fiber.Ctx) error {
	return nil
}


func (h *bookHandler) GetAllBooks(c *fiber.Ctx) error {
	return nil
}

func (h *bookHandler) GetBookById(c *fiber.Ctx) error {
	return nil
}

func (h *bookHandler) UpdateBook(c *fiber.Ctx) error {
	return nil
}

func (h *bookHandler) DeleteBook(c *fiber.Ctx) error {
	return nil
}
