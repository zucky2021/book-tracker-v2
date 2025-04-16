package repository

import (
	"backend/domain"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
)

type BookshelfRepositoryImpl struct{}

func NewBookshelfRepository() domain.BookshelfRepository {
	return &BookshelfRepositoryImpl{}
}

func (r *BookshelfRepositoryImpl) FindByID(userId string, shelfId int) (*domain.Bookshelf, error) {
	apiKey := os.Getenv("GOOGLE_BOOKS_API_KEY")

	baseURL := "https://www.googleapis.com/books/v1/users"
	u, err := url.Parse(fmt.Sprintf("%s/%s/bookshelves/%d", baseURL, userId, shelfId))
	if err != nil {
		return nil, fmt.Errorf("failed to parse URL: %w", err)
	}

	q := u.Query()
	q.Set("key", apiKey)
	u.RawQuery = q.Encode()

	resp, err := http.Get(u.String())
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := resp.Body.Close(); err != nil {
			log.Printf("failed to close response body: %v", err)
		}
	}()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch bookshelf: status code %d", resp.StatusCode)
	}

	var bookshelf domain.Bookshelf
	if err := json.NewDecoder(resp.Body).Decode(&bookshelf); err != nil {
		return nil, fmt.Errorf("failed to decode response body: %w", err)
	}

	return &bookshelf, nil
}
