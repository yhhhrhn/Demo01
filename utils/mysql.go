package utils

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

const (
//UserName     string = "root"
//Password     string = "1qaz@WSX"
//Addr         string = "localhost"
//Port         int    = 3306
//Database     string = "app_store_crawler"
//MaxLifetime  int    = 10
//MaxOpenConns int    = 10
//MaxIdleConns int    = 10

)

var (
	db            *gorm.DB
	sqlConnection = "root:1qaz@WSX@tcp(127.0.0.1:3306)/app_store_crawler"
)

//var db *sql.DB

func init() {

	//開啟資料庫連接
	var err error
	db, err = gorm.Open("mysql", sqlConnection)
	if err != nil {
		panic("failed to connect database")
	}

	//自動建立
	//db.AutoMigrate(&GormUser{})

}

func GetMysqlDB() *gorm.DB {
	return db
}
