package infrastructure

import (
	"backend/domain"

	"gorm.io/gorm"
)

type GormUnitOfWork struct {
	writer *gorm.DB
	reader *gorm.DB
}

func NewGormUnitOfWork(writer, reader *gorm.DB) domain.UnitOfWork {
	return &GormUnitOfWork{
		writer: writer,
		reader: reader,
	}
}

func (uow *GormUnitOfWork) ExecuteInTransaction(fn func(tx *gorm.DB) error) error {
	return uow.writer.Transaction(func(tx *gorm.DB) error {
		return fn(tx)
	})
}

func (uow *GormUnitOfWork) Reader() *gorm.DB {
	return uow.reader
}
