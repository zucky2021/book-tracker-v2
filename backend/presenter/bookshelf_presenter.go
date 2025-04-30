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

func (bp *BookshelfPresenter) Output(
	c *gin.Context,
	bookshelf *domain.Bookshelf,
) {
	c.JSON(http.StatusOK, gin.H{
		"bookshelf": bookshelf,
	})
}

func (bp *BookshelfPresenter) OutputError(c *gin.Context, status int, err error) {
	c.JSON(status, gin.H{
		"error": err.Error(),
	})
}
