package app

import (
	"log"

	"github.com/ashiruhabeeb/simple-library/app/handler"
	"github.com/ashiruhabeeb/simple-library/app/repository"
	"github.com/ashiruhabeeb/simple-library/app/services"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func SetUpRoutes(db *gorm.DB) {
	bookRepo := repository.NewBooksRepo(db)
	booksService := services.NewBooksService(bookRepo)
	booksHandler := handler.NewBookHandler(booksService)

	fiber := fiber.New()

	f1 := fiber.Group("/api/v1")
	f1.Post("", booksHandler.CreateBook)
	f1.Get("", booksHandler.GetAllBooks)
	f1.Get("", booksHandler.GetBookById)
	f1.Patch("", booksHandler.UpdateBook)
	f1.Delete("", booksHandler.DeleteBook)

	log.Println("[INIT] App routes sucessfully set up..🎲")
}