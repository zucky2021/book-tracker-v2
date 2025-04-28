package controller

import (
	"backend/domain"
	"backend/presenter"
	"backend/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type MemoController struct {
	presenter  *presenter.MemoPresenter
	createMemo *usecase.CreateMemoUseCase
	getMemo    *usecase.GetMemoUseCase
	updateMemo *usecase.UpdateMemoUseCase
	deleteMemo *usecase.DeleteMemoUseCase
}

func NewMemoController(
	presenter *presenter.MemoPresenter,
	createMemo *usecase.CreateMemoUseCase,
	getMemo *usecase.GetMemoUseCase,
	updateMemo *usecase.UpdateMemoUseCase,
	deleteMemo *usecase.DeleteMemoUseCase,
) *MemoController {
	return &MemoController{
		presenter:  presenter,
		createMemo: createMemo,
		getMemo:    getMemo,
		updateMemo: updateMemo,
		deleteMemo: deleteMemo,
	}
}

// DTO for create request
type CreateMemoRequest struct {
	UserId      string `json:"userId" binding:"required"`
	BookId      string `json:"bookId" binding:"required"`
	Text        string `json:"text" binding:"required,max=1000"`
	ImgFileName string `json:"imgFileName"`
}

// メモを登録
func (mc *MemoController) CreateMemo(c *gin.Context) {
	var req CreateMemoRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		mc.presenter.OutputError(c, http.StatusBadRequest, err)
		return
	}

	// TODO: 画像ファイルの保存処理を追加する
	memo, err := mc.createMemo.Execute(req.UserId, req.BookId, req.Text, req.ImgFileName)
	if err != nil {
		mc.presenter.OutputError(c, http.StatusInternalServerError, err)
		return
	}
	mc.presenter.Output(c, memo)
}

func (mc *MemoController) GetMemo(c *gin.Context) {
	memoId := c.Param("memoId")
	userId := c.Param("userId")

	intMemoId, err := strconv.ParseUint(memoId, 10, 64)
	if err != nil {
		mc.presenter.OutputError(c, http.StatusBadRequest, err)
		return
	}
	memo, err := mc.getMemo.Execute(uint(intMemoId), userId)
	if err != nil {
		mc.presenter.OutputError(c, http.StatusInsufficientStorage, err)
		return
	}

	mc.presenter.Output(c, memo)
}

type UpdateMemoRequest struct {
	BookId      string `json:"bookId" binding:"required"`
	Text        string `json:"text" binding:"required,max=1000"`
	ImgFileName string `json:"imgFileName"`
}

func (mc *MemoController) UpdateMemo(c *gin.Context) {
	memoId := c.Param("memoId")
	userId := c.Query("userId")

	intMemoId, err := strconv.ParseUint(memoId, 10, 64)
	if err != nil {
		mc.presenter.OutputError(c, http.StatusBadRequest, err)
		return
	}

	var req UpdateMemoRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		mc.presenter.OutputError(c, http.StatusBadRequest, err)
		return
	}

	memo := domain.Memo{
		ID:          uint(intMemoId),
		UserID:      userId,
		BookID:      req.BookId,
		Text:        req.Text,
		ImgFileName: req.ImgFileName,
	}
	updatedMemo, err := mc.updateMemo.Execute(memo)
	if err != nil {
		mc.presenter.OutputError(c, http.StatusInternalServerError, err)
		return
	}
	mc.presenter.Output(c, updatedMemo)
}

func (mc *MemoController) DeleteMemo(c *gin.Context) {
	memoId := c.Param("memoId")
	userId := c.Query("userId")

	intMemoId, err := strconv.ParseUint(memoId, 10, 64)
	if err != nil {
		mc.presenter.OutputError(c, http.StatusBadRequest, err)
		return
	}

	err = mc.deleteMemo.Execute(uint(intMemoId), userId)
	if err != nil {
		mc.presenter.OutputError(c, http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusNoContent, nil)
}
