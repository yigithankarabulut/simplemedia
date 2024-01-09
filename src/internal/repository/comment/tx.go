package commentstorage

import "gorm.io/gorm"

func (r *commentStorage) SetTx(tx ...*gorm.DB) *gorm.DB {
	if len(tx) > 0 {
		return tx[0]
	}
	return r.db
}

func (r *commentStorage) CreateTx() *gorm.DB {
	return r.db.Begin()
}

func (r *commentStorage) CommitTx(tx *gorm.DB) {
	tx.Commit()
}

func (r *commentStorage) RollBackTx(tx *gorm.DB) {
	tx.Rollback()
}
