package repository

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"testing"

	"backend/domain"
)

// テスト用BookRepositoryImpl（baseURLを差し替え可能にする）
type testBookRepositoryImpl struct {
	baseURL string
}

func (r *testBookRepositoryImpl) FindAll(userId string, shelfId int, startIndex int, maxResults int) ([]domain.Book, error) {
	apiKey := os.Getenv("GOOGLE_BOOKS_API_KEY")
	u, _ := url.Parse(r.baseURL + "/" + userId + "/bookshelves/1/volumes")
	q := u.Query()
	q.Set("startIndex", "0")
	q.Set("maxResults", "10")
	q.Set("key", apiKey)
	u.RawQuery = q.Encode()

	resp, err := http.Get(u.String())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, err
	}

	var result struct {
		Items []domain.Book `json:"items"`
	}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}
	return result.Items, nil
}

func TestBookRepositoryImpl_FindAll(t *testing.T) {
	// モックサーバーを作成
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

	repo := &testBookRepositoryImpl{baseURL: ts.URL + "/users"}

	books, err := repo.FindAll("testuser", 1, 0, 10)
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
