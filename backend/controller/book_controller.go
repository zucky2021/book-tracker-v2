package controller

import (
	"backend/presenter"
	"backend/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type BookController struct {
	BookUseCase   *usecase.BookUseCase
	BookPresenter *presenter.BookPresenter
}

func NewBookController(
	bookUseCase *usecase.BookUseCase,
	bookPresenter *presenter.BookPresenter,
) *BookController {
	return &BookController{
		BookUseCase:   bookUseCase,
		BookPresenter: bookPresenter,
	}
}

func (bc *BookController) GetBooks(c *gin.Context) {
	userId := c.Query("userId")

	shelfId, err := strconv.Atoi(c.Query("shelfId"))
	if err != nil {
		bc.BookPresenter.OutputError(c, http.StatusInternalServerError, err)
		return
	}

	startIndex, err := strconv.Atoi(c.Query("startIndex"))
	if err != nil {
		startIndex = 0 // デフォルト値
	}

	maxResults, err := strconv.Atoi(c.Query("maxResults"))
	if err != nil {
		maxResults = 40 // デフォルト値
	}

	books, err := bc.BookUseCase.Execute(userId, shelfId, startIndex, maxResults)
	if err != nil {
		bc.BookPresenter.OutputError(c, http.StatusInternalServerError, err)
		return
	}

	bc.BookPresenter.Output(c, books)
}
