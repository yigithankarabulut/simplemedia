package poststorage

import "gorm.io/gorm"

func (r *postStorage) SetTx(tx ...*gorm.DB) *gorm.DB {
	if len(tx) > 0 {
		return tx[0]
	}
	return r.db
}

func (r *postStorage) CreateTx() *gorm.DB {
	return r.db.Begin()
}

func (r *postStorage) CommitTx(tx *gorm.DB) {
	tx.Commit()
}

func (r *postStorage) RollBackTx(tx *gorm.DB) {
	tx.Rollback()
}
