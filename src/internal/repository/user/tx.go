package userstorage

import (
	"gorm.io/gorm"
)

func (r *userStorage) SetTx(tx ...*gorm.DB) *gorm.DB {
	if len(tx) > 0 {
		return tx[0]
	}
	return r.db
}

func (r *userStorage) CreateTx() *gorm.DB {
	return r.db.Begin()
}

func (r *userStorage) CommitTx(tx *gorm.DB) {
	tx.Commit()
}

func (r *userStorage) RollBackTx(tx *gorm.DB) {
	tx.Rollback()
}
