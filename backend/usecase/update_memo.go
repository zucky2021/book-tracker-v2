package usecase

import (
	"backend/domain"
	"fmt"
)

type UpdateMemoUseCase struct {
	repo domain.MemoRepository
}

func NewUpdateMemoUseCase(repo domain.MemoRepository) *UpdateMemoUseCase {
	return &UpdateMemoUseCase{repo: repo}
}

func (uc *UpdateMemoUseCase) Execute(memo domain.Memo) (domain.Memo, error) {
	existingMemo, err := uc.repo.FindByID(memo.ID, memo.UserID)
	if err != nil {
		return domain.Memo{}, err
	}

	if existingMemo.UserID != memo.UserID {
		return domain.Memo{}, fmt.Errorf("unauthorized access: user %s does not own memo %d", memo.UserID, memo.ID)
	}

	return uc.repo.Update(memo)
}
