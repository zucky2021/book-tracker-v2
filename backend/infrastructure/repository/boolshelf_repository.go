package repository

import (
	"backend/domain"
	"fmt"
	"net/http"
	"encoding/json"
	"os"
)

type BookshelfRepositoryImpl struct {}

func NewBookshelfRepository() domain.BookshelfRepository {
	return &BookshelfRepositoryImpl{}
}

func (r *BookshelfRepositoryImpl) FindById(userId string, shelfId int) (*domain.Bookshelf, error) {
	apiKey := os.Getenv("GOOGLE_BOOKS_API_KEY")

	uri := fmt.Sprintf(
		"https://www.googleapis.com/books/v1/users/%s/bookshelves/%d?key=%s",
		userId, shelfId, apiKey,
	)

	resp, err := http.Get(uri)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch bookshelf: status code %d", resp.StatusCode)
	}

	var bookshelf domain.Bookshelf
	if err := json.NewDecoder(resp.Body).Decode(&bookshelf); err != nil {
		return nil, fmt.Errorf("failed to decode response body: %w", err)
	}

	return &bookshelf, nil
}