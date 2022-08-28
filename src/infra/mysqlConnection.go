package infra

import (
	"fmt"
	"net/url"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func MustNewMySQLConnection() *gorm.DB {
	db, err := NewMySQLConnection()
	if err != nil {
		panic(fmt.Sprintf("failed to connect database! (%s)", err.Error()))
	}
	return db
}

func NewMySQLConnection() (db *gorm.DB, err error) {
	conn := fmt.Sprintf("%s:%s@tcp(%s)/%s",
		EnvMan.DBUser,
		EnvMan.DBPassword,
		EnvMan.HOST,
		EnvMan.DBName,
	)
	val := url.Values{}
	val.Add("charset", "utf8mb4")
	val.Add("parseTime", "True")
	val.Add("loc", "Asia/Tokyo")
	dsn := fmt.Sprintf("%s?%s", conn, val.Encode())

	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return
	}

	if EnvMan.AppEnv == "development" {
		db.Logger.LogMode(logger.Info)
	} else {
		db.Logger.LogMode(logger.Error)
	}

	mysqlDB, _ := db.DB()
	mysqlDB.SetMaxIdleConns(20)
	mysqlDB.SetMaxOpenConns(20)
	mysqlDB.SetConnMaxLifetime(24 * time.Minute)

	return
}

// MustNewMySQLReadOnlyConnection は、アプリケーション全体で共有するGORMのコネクションをセットします
func MustNewMySQLReadOnlyConnection() *gorm.DB {
	db, err := NewMySQLReadOnlyConnection()
	if err != nil {
		panic(fmt.Sprintf("failed to connect database! (%s)", err.Error()))
	}
	return db
}

// NewMySQLReadOnlyConnection は、アプリケーション全体で共有するGORMのコネクションをセットします
func NewMySQLReadOnlyConnection() (db *gorm.DB, err error) {
	conn := fmt.Sprintf("%s:%s@tcp(%s)/%s",
		EnvMan.DBUser,
		EnvMan.DBPassword,
		EnvMan.HOST,
		EnvMan.DBName,
	)

	val := url.Values{}
	val.Add("charset", "utf8mb4")
	val.Add("parseTime", "True")
	val.Add("loc", "Asia/Tokyo")
	dsn := fmt.Sprintf("%s?%s", conn, val.Encode())

	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return
	}
	// db.Session(&gorm.Session{SkipDefaultTransaction: true})

	if EnvMan.AppEnv == "development" {
		db.Logger.LogMode(logger.Info) // Info
	} else {
		db.Logger.LogMode(logger.Error) // Silent
	}
	// MaxOpenConns が 0 よりも大きく、新しい MaxIdleConns よりも小さい場合
	// 新しい MaxIdleConns は MaxOpenConns の制限値に合わせて減少します。
	mysqlDB, _ := db.DB()
	mysqlDB.SetMaxIdleConns(20)
	mysqlDB.SetMaxOpenConns(20)

	// RDSがスケールアウトしても、sql.DBがアイドル状態のコネクションが残しており
	// 負荷が高いままのインスタンスに接続しに行ってしまうので使ってないコネクションは早めに閉じるようにします
	mysqlDB.SetConnMaxLifetime(24 * time.Minute)

	return
}
