package usecase

import (
	"backend/domain"
	"fmt"

	"gorm.io/gorm"
)

type UpdateMemoUseCase struct {
	uow  domain.UnitOfWork
	repo domain.MemoRepository
}

func NewUpdateMemoUseCase(uow domain.UnitOfWork, repo domain.MemoRepository) *UpdateMemoUseCase {
	return &UpdateMemoUseCase{
		uow:  uow,
		repo: repo,
	}
}

func (uc *UpdateMemoUseCase) Execute(
	memoID uint,
	userID string,
	text string,
	imgFileName string,
) (domain.Memo, error) {
	var result domain.Memo

	memo, findErr := uc.repo.FindByID(uc.uow.Reader(), memoID, userID)
	if findErr != nil {
		return result, findErr
	}

	if len(text) > domain.TextMaxLength {
		return result, fmt.Errorf("memo text exceeds maximum length of %d characters", domain.TextMaxLength)
	}

	memo.Text = text
	memo.ImgFileName = imgFileName

	updateErr := uc.uow.ExecuteInTransaction(func(tx *gorm.DB) error {
		updated, err := uc.repo.Update(tx, memo)
		if err != nil {
			return err
		}

		// FIXME: Add S3 process.

		result = updated
		return nil
	})
	return result, updateErr
}
