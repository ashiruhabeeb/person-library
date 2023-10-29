package handler

import (
	"fmt"
	"strconv"
	"time"

	"github.com/ashiruhabeeb/simple-library/app/model"
	"github.com/ashiruhabeeb/simple-library/app/services"
	"github.com/go-playground/validator/v10"
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
	var payload model.BookPayload

	if err := c.BodyParser(&payload); err != nil {
		// handlerError(err)
		errorMessages := []string{}

		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("struct field error %v, condition: %v", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}

		c.Status(400).JSON(fiber.Map{"error": errorMessages})
	}

	book, err := h.bs.Create(payload)
	if err != nil {
		handlerError(err)
	}

	bkRes := convertToBookResponse(book)
	
	return c.Status(201).JSON(fiber.Map{"data": bkRes})
}

func (h *bookHandler) GetAllBooks(c *fiber.Ctx) error {
	books, err := h.bs.FindAll()
	if err != nil {
		c.Status(500).JSON(fiber.Map{"error": err})
	}
	
	var bkRes []BookResponse

	for _, b := range books {
		bookRes := convertToBookResponse(b)
		bkRes = append(bkRes, bookRes)
	}

	return c.Status(200).JSON(fiber.Map{"data": bkRes})
}

func (h *bookHandler) GetBookById(c *fiber.Ctx) error {
	idString := c.Params("books_id")

	id, err := strconv.Atoi(idString)
	if err != nil {
		c.Status(500).JSON(fiber.Map{"error": err})
	}

	book, err := h.bs.FindByBookId(uint(id))
	if err != nil {
		c.Status(500).JSON(fiber.Map{"error": err})
	}
	
	bkRes := convertToBookResponse(book)

	return c.Status(200).JSON(fiber.Map{"data": bkRes})
}

func (h *bookHandler) UpdateBook(c *fiber.Ctx) error {
	var payload model.BookPayload

	if err := c.BodyParser(&payload); err != nil {
		c.Status(400).JSON(fiber.Map{"error": err})
	}

	idString := c.Params("books_id")

	id, err := strconv.Atoi(idString)
	if err != nil {
		c.Status(500).JSON(fiber.Map{"error": err})
	}

	book, err := h.bs.Update(uint(id), payload)
	handlerError(err)

	bkRes := convertToBookResponse(book)

	return c.Status(200).JSON(fiber.Map{"data": bkRes})
}

func (h *bookHandler) DeleteBook(c *fiber.Ctx) error {
	idString := c.Params("books_id")

	id, err := strconv.Atoi(idString)
	if err != nil {
		c.Status(400).JSON(fiber.Map{"error": err})
	}

	b, err := h.bs.Delete(uint(id))
	if err != nil {
		c.Status(400).JSON(fiber.Map{"error": err})
	}

	bkRes := convertToBookResponse(b)
	
	return c.Status(200).JSON(fiber.Map{"data": bkRes})
}

// convertToBookResponse converts data from model.Books struct to BookResponse struct
func convertToBookResponse(bk model.Books) BookResponse {
	return BookResponse{
		BooksId:     bk.BooksId,
		Title:       bk.Title,
		Description: bk.Description,
		Author:      bk.Author,
		Acquisition: bk.Acquisition,
		Price:       bk.Price,
		Rating:      uint(bk.Rating),
		Discount:    bk.Discount,
		Quanitity:   bk.Quantity,
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
	}
}
