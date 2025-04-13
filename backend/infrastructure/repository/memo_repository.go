package repository

import (
	"backend/domain"

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
	err := mr.DB.Where("id = ? AND user_id = ?", id, userId).First(&memo).Error
	if err != nil {
		return domain.Memo{}, err
	}
	return memo, nil
}