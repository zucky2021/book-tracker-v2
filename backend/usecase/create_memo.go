package usecase

import (
	"backend/domain"

	"gorm.io/gorm"
)

type CreateMemoUseCase struct {
	uow  domain.UnitOfWork
	repo domain.MemoRepository
}

func NewCreateMemoUseCase(
	uow domain.UnitOfWork,
	repo domain.MemoRepository,
) *CreateMemoUseCase {
	return &CreateMemoUseCase{
		uow:  uow,
		repo: repo,
	}
}

func (uc *CreateMemoUseCase) Execute(userId, bookId, text, imgFileName string) (domain.Memo, error) {
	var result domain.Memo
	err := uc.uow.ExecuteInTransaction(func(tx *gorm.DB) error {
		memo := domain.Memo{
			UserID:      userId,
			BookID:      bookId,
			Text:        text,
			ImgFileName: imgFileName,
		}

		created, err := uc.repo.Create(tx, memo)
		if err != nil {
			return err
		}

		// FIXME:S3に画像を保存する処理を追加(失敗したらロールバック)

		result = created
		return nil
	})
	return result, err
}
