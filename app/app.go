package app

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
	"time"

	"github.com/ashiruhabeeb/simple-library/app/handler"
	"github.com/ashiruhabeeb/simple-library/app/repository"
	"github.com/ashiruhabeeb/simple-library/app/services"
	"github.com/ashiruhabeeb/simple-library/pkg/config"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"gorm.io/gorm"
)

func SetUpRoutes(db *gorm.DB, cfg *config.AppConfig) {
	bookRepo := repository.NewBooksRepo(db)
	booksService := services.NewBooksService(bookRepo)
	booksHandler := handler.NewBookHandler(booksService)

	fiber := fiber.New()
	fiber.Use(logger.New())

	fiber.Get("/api/healthcheck", handler.HealthCheck)

	// Book grouped route handlers
	f1 := fiber.Group("/api/book")

	f1.Post("/v1/create", booksHandler.CreateBook)
	f1.Get("/v1/allbook", booksHandler.GetAllBooks)
	f1.Get("/v1/getbook/:id", booksHandler.GetBookById)
	f1.Patch("/v1/updatebook/:id", booksHandler.UpdateBook)
	f1.Delete("/v1/delete/:id", booksHandler.DeleteBook)
	
	log.Println("[INIT] App routes sucessfully set up..🎲")

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	var serverShutdown sync.WaitGroup

	ctx, cancel := context.WithTimeout(context.Background(), cfg.Server.Timeout.Server * time.Second)
	defer cancel()

	go func() {
		<-c
		fmt.Println("Gracefully shutting down server...")
		serverShutdown.Add(1)
		defer serverShutdown.Done()
		_ = fiber.ShutdownWithContext(ctx)
	}()

	if err := fiber.Listen(fmt.Sprintf(":" + cfg.Server.Port)); err != nil {
		log.Panic(err)
	}

	serverShutdown.Wait()
}
