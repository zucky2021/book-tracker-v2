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

func (uc *GetMemoUseCase) Execute(userID, bookID string) (domain.Memo, error) {
	memo, err := uc.repo.FindByID(uc.uow.Reader(), userID, bookID)
	if err != nil {
		return domain.Memo{}, err
	}
	return memo, nil
}
