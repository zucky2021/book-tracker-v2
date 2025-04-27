package domain

import (
	"fmt"
	"time"
)

const (
	TextMaxLength = 1000
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

func (m Memo) GetImageUrl() string {
	if m.ImgFileName == "" {
		return ""
	}
	return fmt.Sprintf("%s/%s", m.UserID, m.ImgFileName)
}

type MemoRepository interface {
	FindByID(id uint, userID string) (Memo, error)
	Create(memo Memo) (Memo, error)
	// Update(memo Memo) (Memo, error)
	// Delete(id uint, userID string) error
}
