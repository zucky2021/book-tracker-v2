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

type BookshelfRepositoryImpl struct {
	baseURL string
	apiKey  string
}

func NewBookshelfRepository(baseURL string) domain.BookshelfRepository {
	if baseURL == "" {
		log.Fatalf("baseURL is required")
	}

	return &BookshelfRepositoryImpl{
		baseURL: baseURL,
		apiKey:  os.Getenv("GOOGLE_BOOKS_API_KEY"),
	}
}

func (br *BookshelfRepositoryImpl) FindByID(userId string, shelfId int) (*domain.Bookshelf, error) {
	u, err := url.Parse(fmt.Sprintf("%s/users/%s/bookshelves/%d", br.baseURL, userId, shelfId))
	if err != nil {
		return nil, fmt.Errorf("failed to parse URL: %w", err)
	}

	q := u.Query()
	q.Set("key", br.apiKey)
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
