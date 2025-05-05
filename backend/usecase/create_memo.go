package usecase

import (
	"backend/domain"
	"context"
	"fmt"
	"mime/multipart"
	"path/filepath"

	"gorm.io/gorm"
)

type CreateMemoUseCase struct {
	uow         domain.UnitOfWork
	memoRepo    domain.MemoRepository
	storageRepo domain.StorageRepository
}

func NewCreateMemoUseCase(
	uow domain.UnitOfWork,
	memoRepo domain.MemoRepository,
	storageRepo domain.StorageRepository,
) *CreateMemoUseCase {
	return &CreateMemoUseCase{
		uow:         uow,
		memoRepo:    memoRepo,
		storageRepo: storageRepo,
	}
}

func (uc *CreateMemoUseCase) Execute(
	ctx context.Context,
	userID string,
	bookID string,
	text string,
	imgData []byte,
	header *multipart.FileHeader,
) (domain.Memo, error) {
	var result domain.Memo

	var imgFileName string
	if len(imgData) > 0 {
		ext := filepath.Ext(header.Filename)
		imgFileName = domain.GenerateImgFileName(ext)
	}
	err := uc.uow.ExecuteInTransaction(func(tx *gorm.DB) error {
		memo := domain.Memo{
			UserID:      userID,
			BookID:      bookID,
			Text:        text,
			ImgFileName: imgFileName,
		}

		created, err := uc.memoRepo.Create(tx, memo)
		if err != nil {
			return err
		}

		if imgFileName != "" && len(imgData) > 0 {
			key := fmt.Sprintf("%s/%s", userID, imgFileName)
			err = uc.storageRepo.Upload(ctx, key, imgData)
			if err != nil {
				return fmt.Errorf("failed to upload image to S3: %w", err)
			}
		}

		result = created
		return nil
	})
	return result, err
}
