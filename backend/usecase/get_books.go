package usecase

import (
	"backend/domain"
)

type BookUseCase struct {
	BookRepo domain.BookRepository
}

func NewBookUseCase(bookRepo domain.BookRepository) *BookUseCase {
	return &BookUseCase{
		BookRepo: bookRepo,
	}
}

func (u *BookUseCase) GetBooks(userId string, shelfId int, startIndex int, maxResults int) ([]domain.Book, error) {
	books, err := u.BookRepo.FindAll(userId, shelfId, startIndex, maxResults)
	if err != nil {
		return nil, err
	}
	return books, nil
}
