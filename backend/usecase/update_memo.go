package usecase

import (
	"backend/domain"
	"context"
	"fmt"
	"mime/multipart"
	"path/filepath"

	"gorm.io/gorm"
)

type UpdateMemoUseCase struct {
	uow         domain.UnitOfWork
	repo        domain.MemoRepository
	storageRepo domain.StorageRepository
}

func NewUpdateMemoUseCase(
	uow domain.UnitOfWork,
	repo domain.MemoRepository,
	storageRepo domain.StorageRepository,
) *UpdateMemoUseCase {
	return &UpdateMemoUseCase{
		uow:         uow,
		repo:        repo,
		storageRepo: storageRepo,
	}
}

func (uc *UpdateMemoUseCase) Execute(
	ctx context.Context,
	userID string,
	bookID string,
	text string,
	imgData []byte,
	fileHeader *multipart.FileHeader,
) (domain.Memo, error) {
	var result domain.Memo

	memo, findErr := uc.repo.FindByID(uc.uow.Reader(), userID, bookID)
	if findErr != nil {
		return result, findErr
	}

	if len(text) > domain.TextMaxLength {
		return result, fmt.Errorf("memo text exceeds maximum length of %d characters", domain.TextMaxLength)
	}

	var imgFileName string
	if len(imgData) > 0 {
		ext := filepath.Ext(fileHeader.Filename)
		imgFileName = domain.GenerateImgFileName(ext)
		memo.ImgFileName = imgFileName
	}

	memo.Text = text

	updateErr := uc.uow.ExecuteInTransaction(func(tx *gorm.DB) error {
		if err := uc.repo.UpsertMemo(tx, &memo); err != nil {
			return err
		}

		if imgFileName != "" && len(imgData) > 0 {
			key := fmt.Sprintf("%s/%s", userID, imgFileName)
			if err := uc.storageRepo.Upload(context.TODO(), key, imgData); err != nil {
				return fmt.Errorf("failed to upload image to S3: %w", err)
			}
		}

		result = memo
		return nil
	})
	return result, updateErr
}
