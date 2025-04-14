package repository

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"backend/domain"
)

type BookRepositoryImpl struct{}

func NewBookRepository() domain.BookRepository {
	return &BookRepositoryImpl{}
}

func (r *BookRepositoryImpl) FindAll(userId string, shelfId int, startIndex int, maxResults int) ([]domain.Book, error) {
	apiKey := os.Getenv("GOOGLE_BOOKS_API_KEY")

	url := fmt.Sprintf(
		"https://www.googleapis.com/books/v1/users/%s/bookshelves/%d/volumes?startIndex=%d&maxResults=%d&key=%s",
		userId, shelfId, startIndex, maxResults, apiKey,
	)

	resp, err := http.Get(url)
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
