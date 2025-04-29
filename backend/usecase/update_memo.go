package usecase

import (
	"backend/domain"

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

func (uc *UpdateMemoUseCase) Execute(req domain.Memo) (domain.Memo, error) {
	var result domain.Memo

	memo, findErr := uc.repo.FindByID(uc.uow.Reader(),req.ID, req.UserID)
	if findErr != nil {
		return result, findErr
	}

	memo.Text = req.Text
	memo.ImgFileName = req.ImgFileName

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
