//go:build integration

package controller_test

import (
	"backend/controller"
	"backend/infrastructure"
	"backend/infrastructure/config"
	"backend/infrastructure/repository"
	"backend/presenter"
	"backend/usecase"
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
)

func setupRouter() *gin.Engine {
	r := gin.Default()

	envVarProvider := infrastructure.NewEnvVarProvider()

	dbConns := config.GetDBConnections(envVarProvider)
	uow := infrastructure.NewGormUnitOfWork(dbConns.Writer, dbConns.Reader)

	s3Client := config.NewS3Client(envVarProvider)

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

	r.POST("/api/memo", memoController.CreateMemo)
	return r
}

func TestCreateMemo(t *testing.T) {
	router := setupRouter()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	_ = writer.WriteField("userId", "test_user_id")
	_ = writer.WriteField("bookId", "test_book_id")
	_ = writer.WriteField("text", "テストメモ")

	// テスト用画像ファイルを添付
	imgPath := filepath.Join("testdata", "test.jpg")
	imgFile, _ := os.Open(imgPath)
	defer imgFile.Close()
	part, _ := writer.CreateFormFile("imgFile", filepath.Base(imgPath))
	_, _ = io.Copy(part, imgFile)
	writer.Close()

	req := httptest.NewRequest(http.MethodPost, "/api/memo", body)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, w.Body.String(), "user1")
	assert.Equal(t, w.Body.String(), "book1")
	assert.Equal(t, w.Body.String(), "テストメモ")
}
