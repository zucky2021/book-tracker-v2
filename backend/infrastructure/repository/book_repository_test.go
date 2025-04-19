package repository

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"backend/domain"
)

func TestBookRepositoryImpl_FindAll(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(map[string]interface{}{
			"items": []domain.Book{
				{ID: "1", VolumeInfo: domain.VolumeInfo{Title: "Book 1"}},
				{ID: "2", VolumeInfo: domain.VolumeInfo{Title: "Book 2"}},
			},
		})
	}))
	defer ts.Close()

	repo := NewBookRepository(ts.URL + "/users")

	books, err := repo.FindAll("tester", 1, 0, 10)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if len(books) != 2 {
		t.Errorf("expected 2 books, got %d", len(books))
	}
	if books[0].ID != "1" || books[0].VolumeInfo.Title != "Book 1" {
		t.Errorf("unexpected book[0]: %+v", books[0])
	}
	if books[1].ID != "2" || books[1].VolumeInfo.Title != "Book 2" {
		t.Errorf("unexpected book[1]: %+v", books[1])
	}
}
