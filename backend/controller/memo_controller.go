package controller

import (
	"backend/presenter"
	"backend/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type MemoController struct {
	getMemo   *usecase.GetMemoUseCase
	presenter *presenter.MemoPresenter
}

func NewMemoController(
	getMemo *usecase.GetMemoUseCase,
	presenter *presenter.MemoPresenter,
) *MemoController {
	return &MemoController{
		getMemo:   getMemo,
		presenter: presenter,
	}
}

func (mc *MemoController) GetMemo(c *gin.Context) {
	memoId := c.Query("memoId")
	userId := c.Query("userId")

	intMemoId, err := strconv.ParseUint(memoId, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	memo, err := mc.getMemo.Execute(uint(intMemoId), userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	mc.presenter.Output(c, memo)
}
