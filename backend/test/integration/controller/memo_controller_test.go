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
	_ = writer.WriteField("userId", "testUserId1")
	_ = writer.WriteField("bookId", "testBookId1")
	_ = writer.WriteField("text", "Dummy text.")

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
		t.Fatalf("レスポンスのJSONパースに失敗: %v", err)
	}

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, resp.Memo.UserID, "testUserId1")
	assert.Equal(t, resp.Memo.BookID, "testBookId1")
	assert.Equal(t, resp.Memo.Text, "Dummy text.")

	log.Printf("Finished integration test for MemoController")
}

type MemoResponse struct {
	Memo struct {
		ID          int    `json:"ID"`
		UserID      string `json:"UserID"`
		BookID      string `json:"BookID"`
		Text        string `json:"Text"`
		ImgFileName string `json:"ImgFileName"`
		CreatedAt   string `json:"CreatedAt"`
		UpdatedAt   string `json:"UpdatedAt"`
	} `json:"memo"`
}
