package usecase

import "backend/domain"

type GetMemoUseCase struct {
	repo domain.MemoRepository
}

func NewGetMemoUseCase(repo domain.MemoRepository) *GetMemoUseCase {
	return &GetMemoUseCase{
		repo: repo,
	}
}

func (gu *GetMemoUseCase) Execute(id uint, userId string) (domain.Memo, error) {
	memo, err := gu.repo.FindByID(id, userId)
	if err != nil {
		return domain.Memo{}, err
	}
	return memo, nil
}