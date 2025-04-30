package domain

import "gorm.io/gorm"

type UnitOfWork interface {
	ExecuteInTransaction(fn func(tx *gorm.DB) error) error
	Reader() *gorm.DB
}
