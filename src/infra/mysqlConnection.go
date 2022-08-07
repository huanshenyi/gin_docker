package infra

import (
	"fmt"
	"net/url"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
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
		db.Logger.LogMode(4) // Info
	} else {
		db.Logger.LogMode(1) // Silent
	}

	mysqlDB, _ := db.DB()
	mysqlDB.SetMaxIdleConns(20)
	mysqlDB.SetMaxOpenConns(20)
	mysqlDB.SetConnMaxLifetime(24 * time.Minute)

	return
}
