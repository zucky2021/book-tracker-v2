package repository

import (
	"backend/domain"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type MemoRepositoryImpl struct{}

func NewMemoRepository() domain.MemoRepository {
	return &MemoRepositoryImpl{}
}

func (mr *MemoRepositoryImpl) FindByID(db *gorm.DB, userID, bookID string) (domain.Memo, error) {
	var memo domain.Memo

	if err := db.Where("user_id = ? AND book_id = ?", userID, bookID).First(&memo).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return domain.Memo{}, fmt.Errorf("memo not found: %w", err)
		}
		return domain.Memo{}, fmt.Errorf("error occurred while fetching memo: %w", err)
	}
	return memo, nil
}

func (mr *MemoRepositoryImpl) Create(db *gorm.DB, memo domain.Memo) (domain.Memo, error) {
	if err := db.Create(&memo).Error; err != nil {
		return domain.Memo{}, fmt.Errorf("error occurred while creating memo: %w", err)
	}
	return memo, nil
}

func (mr *MemoRepositoryImpl) Update(db *gorm.DB, memo domain.Memo) (domain.Memo, error) {
	if err := db.Model(&memo).Updates(domain.Memo{
		Text:        memo.Text,
		ImgFileName: memo.ImgFileName,
	}).Error; err != nil {
		return domain.Memo{}, fmt.Errorf("error occurred while updating memo: %w", err)
	}
	return memo, nil
}

func (mr *MemoRepositoryImpl) Delete(db *gorm.DB, id uint, userId string) error {
	if err := db.Where("id = ? AND user_id = ?", id, userId).Delete(&domain.Memo{}).Error; err != nil {
		return fmt.Errorf("error occurred while deleting memo: %w", err)
	}
	return nil
}
