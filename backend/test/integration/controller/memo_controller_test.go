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
	"encoding/json"
	"io"
	"log"
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
	log.Printf("Starting integration test for MemoController")

	router := setupRouter()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	if err := writer.WriteField("userId", "testUserId1"); err != nil {
		t.Fatalf("failed to write userId field: %v", err)
	}
	if err := writer.WriteField("bookId", "testBookId1"); err != nil {
		t.Fatalf("failed to write bookId field: %v", err)
	}
	if err := writer.WriteField("text", "Dummy text."); err != nil {
		t.Fatalf("failed to write text field: %v", err)
	}

	// テスト用画像ファイルを添付
	imgPath := filepath.Join("testdata", "test.jpg")
	imgFile, err := os.Open(imgPath)
	if err != nil {
		t.Fatalf("テスト用画像ファイルのオープンに失敗: %v", err)
	}
	defer imgFile.Close()

	part, err := writer.CreateFormFile("imgFile", filepath.Base(imgPath))
	if err != nil {
		t.Fatalf("フォームファイルの作成に失敗: %v", err)
	}
	_, err = io.Copy(part, imgFile)
	if err != nil {
		t.Fatalf("ファイルのコピーに失敗: %v", err)
	}
	writer.Close()

	req := httptest.NewRequest(http.MethodPost, "/api/memo", body)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	var resp MemoResponse
	responseBody := w.Body.String()
	if err := json.Unmarshal([]byte(responseBody), &resp); err != nil {
		t.Fatalf("レスポンスのJSONパースに失敗: %v\nレスポンス本文: %s", err, responseBody)
	}

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "testUserId1", resp.Memo.UserID)
	assert.Equal(t, "testBookId1", resp.Memo.BookID)
	assert.Equal(t, "Dummy text.", resp.Memo.Text)

	log.Printf("Finished integration test for MemoController")
}

type MemoDetail struct {
	ID          int    `json:"id"`
	UserID      string `json:"userId"`
	BookID      string `json:"bookId"`
	Text        string `json:"text"`
	ImgFileName string `json:"imgFileName"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`
}

type MemoResponse struct {
	Memo MemoDetail `json:"memo"`
}
