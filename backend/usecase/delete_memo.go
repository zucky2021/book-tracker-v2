package usecase

import (
	"backend/domain"

	"gorm.io/gorm"
)

type DeleteMemoUseCase struct {
	uow  domain.UnitOfWork
	repo domain.MemoRepository
}

func NewDeleteMemoUseCase(uow domain.UnitOfWork, repo domain.MemoRepository) *DeleteMemoUseCase {
	return &DeleteMemoUseCase{
		uow:  uow,
		repo: repo,
	}
}

func (uc *DeleteMemoUseCase) Execute(memoId uint, userId string) error {
	err := uc.uow.ExecuteInTransaction(func(tx *gorm.DB) error {
		if err := uc.repo.Delete(tx, memoId, userId); err != nil {
			return err
		}
		return nil
	})

	return err
}
