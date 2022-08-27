package domain

import "gorm.io/gorm"

type Tx interface {
	Close() error
	ReadDB() *gorm.DB
	DB() *gorm.DB
	Begin() (tx Tx)
	Rollback() (tx Tx)
	Commit() (tx Tx, err error)
	GetErrors() (errs error)
}
