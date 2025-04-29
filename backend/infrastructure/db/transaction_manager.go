package db

import (
	"backend/infrastructure/config"

	"gorm.io/gorm"
)

type TransactionManager struct {
	db *config.DBConnections
}

func NewTransactionManager(db *config.DBConnections) *TransactionManager {
	return &TransactionManager{db: db}
}

func (tm *TransactionManager) BeginTransaction() (*gorm.DB, error) {
	tx := tm.db.Writer.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}
	return tx, nil
}

func (tm *TransactionManager) ExecuteInTransaction(fn func(tx *gorm.DB) error) error {
	tx := tm.db.Writer.Begin()
	if tx.Error != nil {
		return tx.Error
	}
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			panic(r)
		}
	}()

	if err := fn(tx); err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (tm *TransactionManager) CommitTransaction(tx *gorm.DB) error {
	return tx.Commit().Error
}

func (tm *TransactionManager) RollbackTransaction(tx *gorm.DB) error {
	return tx.Rollback().Error
}

