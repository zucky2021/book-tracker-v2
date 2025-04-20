package usecase

import (
	"backend/domain"
	"errors"
	"reflect"
	"testing"
)

type MockBookRepository struct {
	MockFindAll func(userId string, shelfId int, startIndex int, maxResult int) ([]domain.Book, error)
}

func (mr *MockBookRepository) FindAll(userId string, shelfId int, startIndex int, maxResult int) ([]domain.Book, error) {
	return mr.MockFindAll(userId, shelfId, startIndex, maxResult)
}

func TestGetBooks(t *testing.T) {
	mockRepo := &MockBookRepository{}
	useCase := NewBookUseCase(mockRepo)

	t.Run("successfully retrieves books", func(t *testing.T) {
		expectedBooks := []domain.Book{
			{
				ID:         "1",
				SaleInfo:   domain.SaleInfo{},
				VolumeInfo: domain.VolumeInfo{},
			},
			{
				ID:         "2",
				SaleInfo:   domain.SaleInfo{},
				VolumeInfo: domain.VolumeInfo{},
			},
		}

		mockRepo.MockFindAll = func(userId string, shelfId int, startIndex int, maxResults int) ([]domain.Book, error) {
			return expectedBooks, nil
		}

		books, err := useCase.Execute("user1", 1, 0, 10)
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}

		if !reflect.DeepEqual(books, expectedBooks) {
			t.Errorf("expected %v, got %v", expectedBooks, books)
		}
	})

	t.Run("handles empty results", func(t *testing.T) {
		mockRepo.MockFindAll = func(userId string, shelfId int, startIndex int, maxResults int) ([]domain.Book, error) {
			return []domain.Book{}, nil
		}

		books, err := useCase.Execute("user1", 1, 0, 10)
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}

		if len(books) != 0 {
			t.Errorf("expected empty result, got %v", books)
		}
	})

	t.Run("returns error when repository fails", func(t *testing.T) {
		mockRepo.MockFindAll = func(userId string, shelfId int, startIndex int, maxResults int) ([]domain.Book, error) {
			return nil, errors.New("repository error")
		}

		_, err := useCase.Execute("user1", 1, 0, 10)
		if err == nil {
			t.Fatal("expected error, got nil")
		}

		expectedError := "repository error"
		if err.Error() != expectedError {
			t.Errorf("expected error %q, got %q", expectedError, err.Error())
		}
	})
}
