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
	createMemo *usecase.CreateMemoUseCase
	presenter *presenter.MemoPresenter
}

func NewMemoController(
	getMemo *usecase.GetMemoUseCase,
	presenter *presenter.MemoPresenter,
	createMemo *usecase.CreateMemoUseCase,
) *MemoController {
	return &MemoController{
		getMemo:   getMemo,
		presenter: presenter,
		createMemo: createMemo,
	}
}

func (mc *MemoController) GetMemo(c *gin.Context) {
	memoId := c.Query("memoId")
	userId := c.Query("userId")

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

// DTO for create request
type CreateMemoRequest struct {
	UserId string `json:"userId" binding:"required"`
	BookId string `json:"bookId" binding:"required"`
	Text string `json:"text" binding:"required,max=1000"`
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
