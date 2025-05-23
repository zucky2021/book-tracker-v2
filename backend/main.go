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

	r := gin.Default()

	config.SetupCORS(r)

	googleBooksURL := envVarProvider.GetGoogleBooksEndpoint()

	bookRepo := repository.NewBookRepository(googleBooksURL)
	bookUsecase := usecase.NewBookUseCase(bookRepo)
	bookPresenter := presenter.NewBookPresenter()
	bookController := controller.NewBookController(bookUsecase, bookPresenter)

	bookshelfRepo := repository.NewBookshelfRepository(googleBooksURL)
	getBookshelf := usecase.NewGetBookshelf(bookshelfRepo)
	bookshelfPresenter := presenter.NewBookshelfPresenter()
	bookshelfController := controller.NewBookshelfController(getBookshelf, bookshelfPresenter)

	memoRepo := repository.NewMemoRepository()
	storageRepo := repository.NewS3Repository(s3Client, envVarProvider)
	memoPresenter := presenter.NewMemoPresenter()
	createMemoUseCase := usecase.NewCreateMemoUseCase(uow, memoRepo, storageRepo)
	getMemoUseCase := usecase.NewGetMemoUseCase(uow, memoRepo)
	updateMemoUseCase := usecase.NewUpdateMemoUseCase(uow, memoRepo, storageRepo)
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
