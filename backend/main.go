package main

import (
	"fmt"
	"log"

	"backend/config"
	"backend/controller"
	"backend/infrastructure"
	"backend/infrastructure/repository"
	"backend/presenter"
	"backend/usecase"

	"github.com/gin-gonic/gin"
)

func main() {
	db := infrastructure.ConnectDB()
	defer func () {
		if err := db.Close(); err != nil {
			log.Printf("Failed to close database connection: %v", err)
		}
	}()

	r := gin.Default()

	config.SetupCORS(r)

	bookRepo := repository.NewBookRepository()
	bookUsecase := usecase.NewBookUseCase(bookRepo)
	bookController := controller.NewBookController(bookUsecase)
	bookPresenter := presenter.NewBookPresenter()

	bookshelfRepo := repository.NewBookshelfRepository()
	getBookshelf := usecase.NewGetBookshelf(bookshelfRepo)
	bookshelfController := controller.NewBookshelfController(getBookshelf)
	bookshelfPresenter := presenter.NewBookshelfPresenter()

	infrastructure.InitRouter(r, bookController, bookPresenter, bookshelfController, bookshelfPresenter)

	fmt.Println("Starting server on :8080")
	if err := r.Run(":8080"); err != nil {
		fmt.Printf("Failed to start server: %v\n", err)
	}
}
