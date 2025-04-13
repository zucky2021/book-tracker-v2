package main

import (
	"fmt"
	"log"

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
	db := config.GetDB()

	log.Println("変更が反映されない")
	

	if err := db.AutoMigrate(&domain.Memo{}); err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	} else {
		log.Println("database migrated successfully")
	}

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
