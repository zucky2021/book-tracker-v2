package usecase

import "backend/domain"

type DeleteMemoUseCase struct {
	repo domain.MemoRepository
}

func NewDeleteMemoUseCase(repo domain.MemoRepository) *DeleteMemoUseCase {
	return &DeleteMemoUseCase{repo: repo}
}

func (uc *DeleteMemoUseCase) Execute(memoId uint, userId string) error {
	return uc.repo.Delete(memoId, userId)
}
