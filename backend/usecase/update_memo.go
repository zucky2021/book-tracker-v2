package usecase

import (
	"backend/domain"
)

type UpdateMemoUseCase struct {
	repo domain.MemoRepository
}

func NewUpdateMemoUseCase(repo domain.MemoRepository) *UpdateMemoUseCase {
	return &UpdateMemoUseCase{repo: repo}
}

func (uc *UpdateMemoUseCase) Execute(req domain.Memo) (domain.Memo, error) {
	memo, err := uc.repo.FindByID(req.ID, req.UserID)
	if err != nil {
		return domain.Memo{}, err
	}

	memo.Text = req.Text
	memo.ImgFileName = req.ImgFileName

	return uc.repo.Update(memo)
}
