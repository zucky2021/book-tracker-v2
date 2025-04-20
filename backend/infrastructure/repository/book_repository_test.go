package repository

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"backend/domain"

	"github.com/stretchr/testify/assert"
)

func TestBookRepositoryImpl_FindAll(t *testing.T) {
	os.Setenv("GOOGLE_BOOKS_API_KEY", "dummy-key")
	defer os.Unsetenv("GOOGLE_BOOKS_API_KEY")

	t.Run("正常系", func(t *testing.T) {
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			expectedPath := "/users/tester/bookshelves/1/volumes"
			if r.URL.Path != expectedPath {
				t.Errorf("expected path %s, got %s", expectedPath, r.URL.Path)
			}

			if startIndex := r.URL.Query().Get("startIndex"); startIndex != "0" {
				t.Errorf("expected startIndex=0, got %s", startIndex)
			}
			if maxResults := r.URL.Query().Get("maxResults"); maxResults != "10" {
				t.Errorf("expected maxResults=10, got %s", maxResults)
			}
			if apiKey := r.URL.Query().Get("key"); apiKey == "" {
				t.Error("expected API key in query parameters")
			}

			w.WriteHeader(http.StatusOK)
			w.Header().Set("Content-Type", "application/json")
			_ = json.NewEncoder(w).Encode(map[string]interface{}{
				"items": []domain.Book{
					{
						ID: "1",
						VolumeInfo: domain.VolumeInfo{
							Title:       "Book 1",
							Authors:     []string{"Author 1"},
							Description: "Description 1",
						},
					},
					{
						ID: "2",
						VolumeInfo: domain.VolumeInfo{
							Title:       "Book 2",
							Authors:     []string{"Author 2"},
							Description: "Description 2",
						},
					},
				},
			})
		}))
		defer ts.Close()

		repo := NewBookRepository(ts.URL)

		books, err := repo.FindAll("tester", 1, 0, 10)
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}
		assert.Equal(t, 2, len(books), "expected 2 books")

		assert.Equal(t, "1", books[0].ID, "expected book ID '1'")

		assert.Equal(t, "Book 1", books[0].VolumeInfo.Title, "expected title 'Book 1'")
	})
}

func TestBookRepositoryImpl_FindAll_HTTPError(t *testing.T) {
	// サーバーをすぐに閉じて接続エラーを発生させる
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
	ts.Close()

	repo := NewBookRepository(ts.URL)
	_, err := repo.FindAll("tester", 1, 0, 10)
	if err == nil {
		t.Fatalf("expected an error, got nil")
	}
}
