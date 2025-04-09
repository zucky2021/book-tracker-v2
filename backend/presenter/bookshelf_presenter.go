package presenter

import (
	"backend/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

type BookshelfPresenter struct{}

func NewBookshelfPresenter() *BookshelfPresenter {
	return &BookshelfPresenter{}
}

func (bp *BookshelfPresenter) PresentBookshelf(c *gin.Context, bookshelf domain.Bookshelf) {
	c.JSON(http.StatusOK, gin.H{
		"bookshelf": bookshelf,
	})
}