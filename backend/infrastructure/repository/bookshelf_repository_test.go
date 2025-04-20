package repository

import (
	"backend/domain"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestBookshelfRepositoryImpl_FindByID(t *testing.T) {
	os.Setenv("GOOGLE_BOOKS_API_KEY", "dummy-key")
	defer os.Unsetenv("GOOGLE_BOOKS_API_KEY")

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		expectedPath := "/users/tester/bookshelves/2"
		if r.URL.Path != expectedPath {
			t.Errorf("expected path %s, got %s", expectedPath, r.URL.Path)
		}

		if apiKey := r.URL.Query().Get("key"); apiKey == "" {
			t.Error("expected API key in query parameters")
		}

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(domain.Bookshelf{
			ID:          2,
			Title:       "To read",
			VolumeCount: 1,
		})
	}))
	defer ts.Close()

	repo := NewBookshelfRepository(ts.URL)

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

func TestBookshelfRepositoryImpl_FindByID_HTTPError(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
	ts.Close()

	repo := NewBookshelfRepository(ts.URL)

	if _, err := repo.FindByID("tester", 2); err == nil {
		t.Error("expected error, got nil")
	}
}
