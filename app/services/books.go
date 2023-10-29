package services

import (
	"time"

	"github.com/ashiruhabeeb/simple-library/app/model"
	"github.com/ashiruhabeeb/simple-library/app/repository"
)

// service implemets BookRepository  struct
type booksService struct {
	repo repository.BooksRepository
}

// NewService creates a new BookService instance
func NewBooksService(repo repository.BooksRepository) *booksService {
	return &booksService{repo: repo}
}

// FidAll returns all records from the books table
func (bs *booksService) FindAll()([]model.Books, error){
	return bs.repo.FindAll()
}

// FindByBookId returns a single book record from the books table based on the bookid provided
func (bs *booksService) FindByBookId(bkId uint)(model.Books, error){
	return bs.repo.FindByBooksID(bkId)
}

// Delete removes a sing record from the books table based on the provided paramater
func (bs *booksService) Delete(bkId uint) (model.Books, error) {
	book, _ := bs.repo.FindByBooksID(bkId)

	return bs.repo.Delete(book)
}

// Create creates a new book record in the books table based on the paramaeters provided
func (bs *booksService) Create(payload model.BookPayload)(model.Books, error){
	price, _ := payload.Price.Float64()
	rating, _ := payload.Rating.Float64()
	discount, _ := payload.Discount.Int64()
	quantity, _ := payload.Quantity.Int64()

	bk := &model.Books{
		Title:       payload.Title,
		Description: payload.Description,
		Author:      payload.Author,
		Acquisition: payload.Acquisition,
		Price:       float32(price),
		Rating:      float32(rating),
		Discount:    float32(discount),
		Quantity:    uint(quantity),
		CreatedAt:   time.Time{},
	}
	return bs.repo.Create(*bk)
}

// Update changes the existing data in the books table based on the provided parameters
func (bs *booksService) Update(bkId uint, payload model.BookPayload)(model.Books, error){
	bk, _ := bs.repo.FindByBooksID(bkId)

	price, _ := payload.Price.Float64()
	rating, _ := payload.Rating.Float64()
	discount, _ := payload.Discount.Int64()
	quantity,_ := payload.Quantity.Int64()

	bk.Title = payload.Title
	bk.Description = payload.Description
	bk.Author = payload.Author
	bk.Acquisition = payload.Acquisition
	bk.Price = float32(price)
	bk.Rating = float32(rating)
	bk.Discount = float32(discount)
	bk.Quantity = uint(quantity)
	bk.UpdatedAt = time.Now() 

	return bs.repo.Update(bk)
}

// BookService implements CRUD methods for books table from the service struct
type BookService interface {
	FindAll()([]model.Books, error)
	FindByBookId(bkId uint)(model.Books, error)
	Delete(bkId uint) (model.Books, error)
	Create(payload model.BookPayload)(model.Books, error)
	Update(bkId uint, payload model.BookPayload)(model.Books, error)
}
