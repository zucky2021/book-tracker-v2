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

func (p *BookPresenter) PresentBooks(c *gin.Context, books []domain.Book, err error) {
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"books": books,
	})
}
