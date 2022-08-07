package domain

import "gorm.io/gorm"

type Tx interface {
	Close() error
	DB() *gorm.DB
	Begin() (tx Tx)
	Rollback() (tx Tx)
	Commit() (tx Tx, err error)
	GetErrors() (errs error)
}
