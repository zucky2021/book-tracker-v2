package repository

import (
	"backend/domain"
	"backend/infrastructure/config"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type MemoRepositoryImpl struct {
	DB *config.DBConnections
}

func NewMemoRepository(db *config.DBConnections) domain.MemoRepository {
	return &MemoRepositoryImpl{DB: db}
}

func (mr *MemoRepositoryImpl) FindByID(id uint, userId string) (domain.Memo, error) {
	var memo domain.Memo

	if err := mr.DB.Reader.Where("id = ? AND user_id = ?", id, userId).First(&memo).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return domain.Memo{}, fmt.Errorf("memo not found: %w", err)
		}
		return domain.Memo{}, fmt.Errorf("error occurred while fetching memo: %w", err)
	}
	return memo, nil
}

func (mr *MemoRepositoryImpl) Create(memo domain.Memo) (domain.Memo, error) {
	if err := mr.DB.Writer.Create(&memo).Error; err != nil {
		return domain.Memo{}, fmt.Errorf("error occurred while creating memo: %w", err)
	}
	return memo, nil
}

func (mr *MemoRepositoryImpl) Update(memo domain.Memo) (domain.Memo, error) {
	if err := mr.DB.Writer.Save(&memo).Error; err != nil {
		return domain.Memo{}, fmt.Errorf("error occurred while updating memo: %w", err)
	}
	return memo, nil
}

func (mr *MemoRepositoryImpl) Delete(id uint, userId string) error {
	if err := mr.DB.Writer.Where("id = ? AND user_id = ?", id, userId).Delete(&domain.Memo{}).Error; err != nil {
		return fmt.Errorf("error occurred while deleting memo: %w", err)
	}
	return nil
}
