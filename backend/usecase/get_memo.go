package usecase

import "backend/domain"

type GetMemoUseCase struct {
	uow  domain.UnitOfWork
	repo domain.MemoRepository
}

func NewGetMemoUseCase(uow domain.UnitOfWork, repo domain.MemoRepository) *GetMemoUseCase {
	return &GetMemoUseCase{
		uow:  uow,
		repo: repo,
	}
}

func (gu *GetMemoUseCase) Execute(id uint, userId string) (domain.Memo, error) {
	memo, err := gu.repo.FindByID(gu.uow.Reader(), id, userId)
	if err != nil {
		return domain.Memo{}, err
	}
	return memo, nil
}
