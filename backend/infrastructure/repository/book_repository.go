package repository

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"

	"backend/domain"
)

type BookRepositoryImpl struct{}

func NewBookRepository() domain.BookRepository {
	return &BookRepositoryImpl{}
}

func (r *BookRepositoryImpl) FindAll(userId string, shelfId int, startIndex int, maxResults int) ([]domain.Book, error) {
	apiKey := os.Getenv("GOOGLE_BOOKS_API_KEY")

	baseURL := "https://www.googleapis.com/books/v1/users"
	u, err := url.Parse(fmt.Sprintf("%s/%s/bookshelves/%d/volumes", baseURL, userId, shelfId))
	if err != nil {
		return nil, fmt.Errorf("failed to parse URL: %w", err)
	}

	q := u.Query()
	q.Set("startIndex", fmt.Sprintf("%d", startIndex))
	q.Set("maxResults", fmt.Sprintf("%d", maxResults))
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
		return nil, fmt.Errorf("failed to fetch books: %s", resp.Status)
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
