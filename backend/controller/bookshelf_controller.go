package controller

import (
	"backend/domain"
	"backend/usecase"
	"strconv"
)

type BookshelfController struct {
	getBookshelf *usecase.GetBookshelf
}

func NewBookshelfController(getBookshelf *usecase.GetBookshelf) *BookshelfController {
	return &BookshelfController{
		getBookshelf: getBookshelf,
	}
}

func (bc *BookshelfController) GetBookshelf(queryParams map[string]string) (*domain.Bookshelf, error) {
	userId := queryParams["userId"]

	shelfId, err := strconv.Atoi(queryParams["shelfId"])
	if err != nil {
		return nil, err
	}

	return bc.getBookshelf.Execute(userId, shelfId)
}
