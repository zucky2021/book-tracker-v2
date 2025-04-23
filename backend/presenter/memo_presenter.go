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
