package repository

import (
	"backend/domain"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type MemoRepositoryImpl struct {
	DB *gorm.DB
}

func NewMemoRepository(db *gorm.DB) *MemoRepositoryImpl {
	return &MemoRepositoryImpl{DB: db}
}

func (mr *MemoRepositoryImpl) FindByID(id uint, userId string) (domain.Memo, error) {
	var memo domain.Memo

	if err := mr.DB.Where("id = ? AND user_id = ?", id, userId).First(&memo).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return domain.Memo{}, fmt.Errorf("memo not found: %w", err)
		}
		return domain.Memo{}, fmt.Errorf("error occurred while fetching memo: %w", err)
	}
	return memo, nil
}