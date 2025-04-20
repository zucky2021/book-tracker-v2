package usecase

import (
	"backend/domain"
	"errors"
	"reflect"
	"testing"
)

type MockBookshelfRepository struct {
	MockFindByID func(userId string, shelfId int) (*domain.Bookshelf, error)
}

func (m *MockBookshelfRepository) FindByID(userId string, shelfId int) (*domain.Bookshelf, error) {
	return m.MockFindByID(userId, shelfId)
}

func TestGetBookshelf(t *testing.T) {
	mockRepo := &MockBookshelfRepository{}
	useCase := NewGetBookshelf(mockRepo)

	t.Run("successfully retrieves bookshelf", func(t *testing.T) {
		expected := &domain.Bookshelf{
			ID:          2,
			Title:       "To read",
			VolumeCount: 1,
		}
		mockRepo.MockFindByID = func(userId string, shelfId int) (*domain.Bookshelf, error) {
			return expected, nil
		}

		result, err := useCase.Execute("user1", 2)
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("expected %v, got %v", expected, result)
		}
	})

	t.Run("returns error when repository fails", func(t *testing.T) {
		mockRepo.MockFindByID = func(userId string, shelfId int) (*domain.Bookshelf, error) {
			return nil, errors.New("repository error")
		}

		_, err := useCase.Execute("user1", 2)
		if err == nil {
			t.Fatal("expected error, got nil")
		}
		if err.Error() != "repository error" {
			t.Errorf("expected error 'repository error', got %v", err)
		}
	})
}
