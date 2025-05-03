package usecase

import (
	"backend/domain"
	"context"
	"fmt"
	"mime/multipart"

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
		uow:  uow,
		repo: repo,
	}
}

func (uc *UpdateMemoUseCase) Execute(
	memoID uint,
	userID string,
	text string,
	imgData []byte,
	fileHeader *multipart.FileHeader,
) (domain.Memo, error) {
	var result domain.Memo

	memo, findErr := uc.repo.FindByID(uc.uow.Reader(), memoID, userID)
	if findErr != nil {
		return result, findErr
	}

	if len(text) > domain.TextMaxLength {
		return result, fmt.Errorf("memo text exceeds maximum length of %d characters", domain.TextMaxLength)
	}

	var imgFileName string
	if len(imgData) > 0 {
		ext := fileHeader.Filename[len(fileHeader.Filename)-4:]
		imgFileName = domain.GenerateImgFileName(ext)
		memo.ImgFileName = imgFileName
	}

	memo.Text = text

	updateErr := uc.uow.ExecuteInTransaction(func(tx *gorm.DB) error {
		updated, err := uc.repo.Update(tx, memo)
		if err != nil {
			return err
		}

		if imgFileName != "" && len(imgData) > 0 {
			key := fmt.Sprintf("%s/%s", userID, imgFileName)
			err = uc.storageRepo.Upload(context.TODO(), key, imgData)
			if err != nil {
				return fmt.Errorf("failed to upload image to S3: %w", err)
			}
		}

		result = updated
		return nil
	})
	return result, updateErr
}
