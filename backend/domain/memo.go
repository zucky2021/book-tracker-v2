package domain

import "time"

type Memo struct {
	ID        uint   `gorm:"primaryKey"`
	UserID    string `gorm:"size:255;not null;uniqueIndex:idx_user_book"`
	BookID    string `gorm:"size:255;not null;uniqueIndex:idx_user_book"`
	Text      string `gorm:"type:text;not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type MemoRepository interface {
	FindByID(id uint, userID string) (Memo, error)
	Create(memo Memo) (Memo, error)
	Update(memo Memo) (Memo, error)
	Delete(id uint, userID string) error
}
