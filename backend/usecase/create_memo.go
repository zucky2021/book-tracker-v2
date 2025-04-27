package usecase

import "backend/domain"

type CreateMemoUseCase struct {
	repo domain.MemoRepository
}

func NewCreateMemoUseCase(repo domain.MemoRepository) *CreateMemoUseCase {
	return &CreateMemoUseCase{repo: repo}
}

func (uc *CreateMemoUseCase) Execute(userId, bookId, text, imgFileName string) (domain.Memo, error) {
	memo := domain.Memo{
		UserID:      userId,
		BookID:      bookId,
		Text:        text,
		ImgFileName: imgFileName,
	}

	return uc.repo.Create(memo)
}
