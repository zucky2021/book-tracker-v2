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
	ID          uint      `gorm:"primaryKey" json:"id"`
	UserID      string    `gorm:"size:255;not null;uniqueIndex:idx_user_book" json:"userId"`
	BookID      string    `gorm:"size:255;not null;uniqueIndex:idx_user_book" json:"bookId"`
	Text        string    `gorm:"type:text;not null" json:"text"`
	ImgFileName string    `gorm:"size:255" json:"imgFileName,omitempty"` // Optional, can be empty if no image is uploaded
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
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
	UpsertMemo(db *gorm.DB, memo *Memo) error
	Delete(db *gorm.DB, id uint, userID string) error
}
