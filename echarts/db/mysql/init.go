package mysql

import (
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	//"gorm.io/gorm"
)

var _DB *gorm.DB

func DB() *gorm.DB {
	return _DB
}

func init() {
	_DB = initDB()
}

func initDB() *gorm.DB {
	db, err := gorm.Open("mysql", "root:wxz@tcp(127.0.0.1:3306)/wakaTime?charset=utf8mb4&parseTime=True&loc=Local")


	if err != nil {
		panic(err)
	}
	db.DB().SetMaxOpenConns(100)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetConnMaxLifetime(time.Second * 300)
	if err = db.DB().Ping(); err != nil {
		panic(err)
	}
	return db
}
