package repository

import (
	"backend/domain"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestBookshelfRepositoryImpl_FindAll(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(domain.Bookshelf{
			ID:          2,
			Title:       "To read",
			VolumeCount: 1,
		})
	}))
	defer ts.Close()

	repo := NewBookshelfRepository(ts.URL + "/users")

	bookshelf, err := repo.FindByID("tester", 2)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if bookshelf == nil {
		t.Errorf("expected a bookshelf, got nil")
	}
	if bookshelf.ID != 2 || bookshelf.Title != "To read" || bookshelf.VolumeCount != 1 {
		t.Errorf("unexpected bookshelf: %+v", bookshelf)
	}
}