package repository

import (
	"backend/domain"
	"errors"
	"fmt"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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

func (mr *MemoRepositoryImpl) UpsertMemo(db *gorm.DB, memo *domain.Memo) error {
	return db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "user_id"}, {Name: "book_id"}},
		DoUpdates: clause.AssignmentColumns([]string{"text", "img_file_name"}),
	}).Create(memo).Error
}

func (mr *MemoRepositoryImpl) Delete(db *gorm.DB, id uint, userId string) error {
	if err := db.Where("id = ? AND user_id = ?", id, userId).Delete(&domain.Memo{}).Error; err != nil {
		return fmt.Errorf("error occurred while deleting memo: %w", err)
	}
	return nil
}
