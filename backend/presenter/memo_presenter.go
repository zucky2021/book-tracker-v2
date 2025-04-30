package presenter

import (
	"backend/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

type MemoPresenter struct{}

func NewMemoPresenter() *MemoPresenter {
	return &MemoPresenter{}
}

func (mp *MemoPresenter) Output(c *gin.Context, memo domain.Memo) {
	c.JSON(http.StatusOK, gin.H{
		"memo": memo,
	})
}

func (mp *MemoPresenter) OutputError(c *gin.Context, status int, err error) {
	_ = c.Error(err)
	c.JSON(status, gin.H{
		"error": http.StatusText(status),
	})
}
