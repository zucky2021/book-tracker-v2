package main

import (
	"context"
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

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/gin-gonic/gin"
)

func main() {
	envVarProvider := infrastructure.NewEnvVarProvider()

	dbConns := config.GetDBConnections(envVarProvider)
	modelsToMigrate := []interface{}{
		&domain.Memo{},
	}
	for _, model := range modelsToMigrate {
		if err := dbConns.Writer.AutoMigrate(model); err != nil {
			log.Fatalf("failed to migrate model %T: %v", model, err)
		}
	}
	log.Println("all database migrations completed successfully")

	uow := infrastructure.NewGormUnitOfWork(dbConns.Writer, dbConns.Reader)

	s3Client := config.NewS3Client(envVarProvider)
	// FIXME: memo機能 and ログ機能
	buckets, err := s3Client.ListBuckets(context.TODO(), &s3.ListBucketsInput{})
	if err != nil {
		log.Fatal("s3 client err: ", err)
	}
	for _, b := range buckets.Buckets {
		log.Print("Bucket name from main.go: ", *b.Name)
	}

	r := gin.Default()

	config.SetupCORS(r)

	googleBooksURL := os.Getenv("GOOGLE_BOOKS_ENDPOINT")

	bookRepo := repository.NewBookRepository(googleBooksURL)
	bookUsecase := usecase.NewBookUseCase(bookRepo)
	bookPresenter := presenter.NewBookPresenter()
	bookController := controller.NewBookController(bookUsecase, bookPresenter)

	bookshelfRepo := repository.NewBookshelfRepository(googleBooksURL)
	getBookshelf := usecase.NewGetBookshelf(bookshelfRepo)
	bookshelfPresenter := presenter.NewBookshelfPresenter()
	bookshelfController := controller.NewBookshelfController(getBookshelf, bookshelfPresenter)

	memoRepo := repository.NewMemoRepository(dbConns)
	memoPresenter := presenter.NewMemoPresenter()
	createMemoUseCase := usecase.NewCreateMemoUseCase(uow, memoRepo)
	getMemoUseCase := usecase.NewGetMemoUseCase(uow, memoRepo)
	updateMemoUseCase := usecase.NewUpdateMemoUseCase(uow, memoRepo)
	deleteMemoUseCase := usecase.NewDeleteMemoUseCase(uow, memoRepo)
	memoController := controller.NewMemoController(
		memoPresenter,
		createMemoUseCase,
		getMemoUseCase,
		updateMemoUseCase,
		deleteMemoUseCase,
	)

	infrastructure.InitRouter(
		r,
		bookController,
		bookshelfController,
		memoController,
	)

	fmt.Println("Starting server on :8080")
	if err := r.Run(":8080"); err != nil {
		fmt.Printf("Failed to start server: %v\n", err)
	}
}
