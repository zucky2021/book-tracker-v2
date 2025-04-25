package presenter

import (
	"backend/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

type BookPresenter struct{}

func NewBookPresenter() *BookPresenter {
	return &BookPresenter{}
}

func (p *BookPresenter) Output(c *gin.Context, books []domain.Book) {
	c.JSON(http.StatusOK, gin.H{
		"books": books,
	})
}

func (bp *BookPresenter) OutputError(c *gin.Context, status int, err error) {
	c.JSON(status, gin.H{
		"error": err.Error(),
	})
}
