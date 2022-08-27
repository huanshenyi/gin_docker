package repository

import (
	"gin_docker/src/domain"
	"gin_docker/src/infra"

	"gorm.io/gorm"
)

type Tx struct {
	db      *gorm.DB
	rodb    *gorm.DB
	isBegun bool
}

func NewTx(db *gorm.DB) domain.Tx {
	return &Tx{db, nil, false}
}

func NewTxEmpty() domain.Tx {
	return &Tx{nil, nil, false}
}

func (t *Tx) Close() error {
	if t.db != nil {
		db, _ := t.db.DB()
		if err := db.Close(); err != nil {
			return err
		}
	}
	if t.rodb != nil {
		rodb, _ := t.rodb.DB()
		if err := rodb.Close(); err != nil {
			return err
		}
	}
	return nil
}

// ReadDB -
func (t *Tx) ReadDB() *gorm.DB {
	if t.isBegun {
		return t.DB()
	}

	if t.rodb == nil {
		t.rodb = infra.MustNewMySQLReadOnlyConnection()
	}
	return t.rodb
}

func (t *Tx) DB() *gorm.DB {
	if t.db == nil {
		t.db = infra.MustNewMySQLConnection()
	}
	return t.db
}

func (t *Tx) Begin() (tx domain.Tx) {
	return &Tx{t.db.Begin(), t.rodb, true}
}

func (t *Tx) Rollback() (tx domain.Tx) {
	return &Tx{t.db.Rollback(), t.rodb, false}
}

func (t *Tx) Commit() (tx domain.Tx, err error) {
	db := t.db.Commit()
	tx = &Tx{db, t.rodb, false}
	return tx, db.Error
}

func (t *Tx) GetErrors() (errs error) {
	return t.db.Error
}
