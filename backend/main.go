package main

import (
	"fmt"
	"log"
	"os"

	"backend/controller"
	"backend/domain"
	"backend/infrastructure"
	"backend/infrastructure/config"
	"backend/infrastructure/repository"
	"backend/presenter"
	"backend/usecase"

	"github.com/gin-gonic/gin"
)

func main() {
	dbConns := config.GetDBConnections()
	modelsToMigrate := []interface{}{
		&domain.Memo{},
	}
	for _, model := range modelsToMigrate {
		if err := dbConns.Writer.AutoMigrate(model); err != nil {
			log.Fatalf("failed to migrate model %T: %v", model, err)
		}
	}
	log.Println("all database migrations completed successfully")

	r := gin.Default()

	config.SetupCORS(r)

	googleBooksURL := os.Getenv("GOOGLE_BOOKS_ENDPOINT")

	bookRepo := repository.NewBookRepository(googleBooksURL)
	bookUsecase := usecase.NewBookUseCase(bookRepo)
	bookController := controller.NewBookController(bookUsecase)
	bookPresenter := presenter.NewBookPresenter()

	bookshelfRepo := repository.NewBookshelfRepository(googleBooksURL)
	getBookshelf := usecase.NewGetBookshelf(bookshelfRepo)
	bookshelfController := controller.NewBookshelfController(getBookshelf)
	bookshelfPresenter := presenter.NewBookshelfPresenter()

	infrastructure.InitRouter(r, bookController, bookPresenter, bookshelfController, bookshelfPresenter)

	fmt.Println("Starting server on :8080")
	if err := r.Run(":8080"); err != nil {
		fmt.Printf("Failed to start server: %v\n", err)
	}
}
