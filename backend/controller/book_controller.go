package controller

import (
	"backend/domain"
	"backend/usecase"
	"strconv"
)

type BookController struct {
	BookUseCase *usecase.BookUseCase
}

func NewBookController(bookUseCase *usecase.BookUseCase) *BookController {
	return &BookController{
		BookUseCase: bookUseCase,
	}
}

func (b *BookController) GetBooks(queryParams map[string]string) ([]domain.Book, error) {
	// クエリパラメータを処理
	userId := queryParams["userId"]

	shelfId, err := strconv.Atoi(queryParams["shelfId"])
	if err != nil {
		return nil, err
	}

	startIndex, err := strconv.Atoi(queryParams["startIndex"])
	if err != nil {
		startIndex = 0 // デフォルト値
	}

	maxResults, err := strconv.Atoi(queryParams["maxResults"])
	if err != nil {
		maxResults = 40 // デフォルト値
	}

	// ユースケース層を呼び出し
	return b.BookUseCase.Execute(userId, shelfId, startIndex, maxResults)
}
