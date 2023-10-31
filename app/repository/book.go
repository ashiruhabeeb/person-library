package repository

import (
	"github.com/ashiruhabeeb/simple-library/app/model"
	"gorm.io/gorm"
)

// BooksRepo implements gorm.DB
type BooksRepo struct {
	db *gorm.DB
}

// NewRepository creates a new BooksRepo instance
func NewBooksRepo(db *gorm.DB) *BooksRepo {
	return &BooksRepo{db: db}
}

// Create saves a new Books data to the bool table
func (b *BooksRepo) Create(bk model.Books) (model.Books, error) {
	err := b.db.Create(&bk).Error
	return bk, err
}

// FindAll fetch all records from the Books table
func (b *BooksRepo) FindAll() ([]model.Books, error) {
	var bks []model.Books
	err := b.db.Find(&bks).Error
	return bks, err
}

// FindByBooksID returns a single record from Books table based on the the provided BooksId
func (b *BooksRepo) FindByBooksID(bkId uint) (model.Books, error) {
	var bk model.Books
	err := b.db.First(&bk, bkId).Error
	return bk, err
}

// Update changes the exxisting data in the Books table
func (b *BooksRepo) Update(bk model.Books) (model.Books, error) {
	err := b.db.Save(&bk).Error
	return bk, err
}

// Delete removes a single record of from Books table
func (b *BooksRepo) Delete(bk model.Books) (model.Books, error){
	err := b.db.Delete(&bk).Error
	return bk, err
}

// BooksRepository stores the CRUD methods for Bookss table
type BooksRepository interface {
	Create(bk model.Books) (model.Books, error)
	FindAll() ([]model.Books, error)
	FindByBooksID(bkId uint) (model.Books, error)
	Update(bk model.Books) (model.Books, error)
	Delete(bk model.Books) (model.Books, error)
}
