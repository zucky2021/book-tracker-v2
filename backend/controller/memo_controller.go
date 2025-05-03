package controller

import (
	"backend/domain"
	"backend/presenter"
	"backend/usecase"
	"io"
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
	UserID string `form:"userId" binding:"required"`
	BookID string `form:"bookId" binding:"required"`
	Text   string `form:"text" binding:"required,max=1000"`
}

// メモを登録
func (mc *MemoController) CreateMemo(c *gin.Context) {
	var req CreateMemoRequest
	if err := c.ShouldBind(&req); err != nil {
		mc.presenter.OutputError(c, http.StatusBadRequest, err)
		return
	}

	file, fileHeader, err := c.Request.FormFile("imgFile")
	var imgData []byte
	if err == nil && file != nil {
		if fileHeader.Size > domain.ImgMaxSize {
			mc.presenter.OutputError(c, http.StatusBadRequest, err)
			return
		}
		imgData, _ = io.ReadAll(file)
	}

	memo, err := mc.createMemo.Execute(req.UserID, req.BookID, req.Text, imgData, fileHeader)
	if err != nil {
		mc.presenter.OutputError(c, http.StatusInternalServerError, err)
		return
	}
	mc.presenter.Output(c, memo)
}

func (mc *MemoController) GetMemo(c *gin.Context) {
	memoID := c.Param("memoId")
	userID := c.Param("userId")

	intMemoID, err := strconv.ParseUint(memoID, 10, 64)
	if err != nil {
		mc.presenter.OutputError(c, http.StatusBadRequest, err)
		return
	}
	memo, err := mc.getMemo.Execute(uint(intMemoID), userID)
	if err != nil {
		mc.presenter.OutputError(c, http.StatusInternalServerError, err)
		return
	}

	mc.presenter.Output(c, memo)
}

type UpdateMemoRequest = CreateMemoRequest

func (mc *MemoController) UpdateMemo(c *gin.Context) {
	memoID := c.Param("memoId")

	intMemoID, err := strconv.ParseUint(memoID, 10, 64)
	if err != nil {
		mc.presenter.OutputError(c, http.StatusBadRequest, err)
		return
	}

	var req UpdateMemoRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		mc.presenter.OutputError(c, http.StatusBadRequest, err)
		return
	}

	file, fileHeader, err := c.Request.FormFile("imgFile")
	var imgData []byte
	if err == nil && file != nil {
		if fileHeader.Size > domain.ImgMaxSize {
			mc.presenter.OutputError(c, http.StatusBadRequest, err)
			return
		}
		imgData, _ = io.ReadAll(file)
	}

	updatedMemo, err := mc.updateMemo.Execute(
		uint(intMemoID),
		req.UserID,
		req.Text,
		imgData,
		fileHeader,
	)
	if err != nil {
		mc.presenter.OutputError(c, http.StatusInternalServerError, err)
		return
	}
	mc.presenter.Output(c, updatedMemo)
}

func (mc *MemoController) DeleteMemo(c *gin.Context) {
	memoID := c.Param("memoId")
	userID := c.Query("userId")

	intMemoID, err := strconv.ParseUint(memoID, 10, 64)
	if err != nil {
		mc.presenter.OutputError(c, http.StatusBadRequest, err)
		return
	}

	err = mc.deleteMemo.Execute(uint(intMemoID), userID)
	if err != nil {
		mc.presenter.OutputError(c, http.StatusBadRequest, err)
		return
	}
	mc.presenter.OutputError(c, http.StatusNoContent, nil)
}
