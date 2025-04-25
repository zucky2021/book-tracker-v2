package controller

import (
	"backend/presenter"
	"backend/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type BookshelfController struct {
	getBookshelf       *usecase.GetBookshelf
	bookshelfPresenter *presenter.BookshelfPresenter
}

func NewBookshelfController(
	getBookshelf *usecase.GetBookshelf,
	bookshelfPresenter *presenter.BookshelfPresenter,
) *BookshelfController {
	return &BookshelfController{
		getBookshelf:       getBookshelf,
		bookshelfPresenter: bookshelfPresenter,
	}
}

func (bc *BookshelfController) GetBookshelf(c *gin.Context) {
	userId := c.Query("userId")

	shelfId, err := strconv.Atoi(c.Query("shelfId"))
	if err != nil {
		bc.bookshelfPresenter.OutputError(c, http.StatusBadRequest, err)
		return
	}

	bookshelf, err := bc.getBookshelf.Execute(userId, shelfId)
	if err != nil {
		bc.bookshelfPresenter.OutputError(c, http.StatusInternalServerError, err)
		return
	}

	bc.bookshelfPresenter.Output(c, bookshelf)
}
