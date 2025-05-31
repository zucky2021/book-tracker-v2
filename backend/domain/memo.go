package domain

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

const (
	TextMaxLength = 1000
	ImgMaxSize    = 5 * 1024 * 1024 // 5MB
)

type Memo struct {
	ID          uint   `gorm:"primaryKey"`
	UserID      string `gorm:"size:255;not null;uniqueIndex:idx_user_book"`
	BookID      string `gorm:"size:255;not null;uniqueIndex:idx_user_book"`
	Text        string `gorm:"type:text;not null"`
	ImgFileName string `gorm:"size:255"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func GenerateImgFileName(ext string) string {
	return fmt.Sprintf("%s%d%s", time.Now().Format("20060102150405"), time.Now().Nanosecond(), ext)
}

func (m Memo) GetImageUrl() string {
	if m.ImgFileName == "" {
		return ""
	}
	return fmt.Sprintf("%s/%s", m.UserID, m.ImgFileName)
}

type MemoRepository interface {
	FindByID(db *gorm.DB, userID, bookID string) (Memo, error)
	Create(db *gorm.DB, memo Memo) (Memo, error)
	Update(db *gorm.DB, memo Memo) (Memo, error)
	Delete(db *gorm.DB, id uint, userID string) error
}
