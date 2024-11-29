package utils

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

const (
	UserName     string = "root"
	Password     string = "1qaz@WSX"
	Addr         string = "localhost"
	Port         int    = 3306
	Database     string = "app_store_crawler"
	MaxLifetime  int    = 10
	MaxOpenConns int    = 10
	MaxIdleConns int    = 10
)

var db *sql.DB

func init() {

	conn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", UserName, Password, Addr, Port, Database)
	//連接MySQL
	mysql, err := sql.Open("mysql", conn)
	if err != nil {
		fmt.Println("connection to mysql failed:", err)
		return
	}

	mysql.SetConnMaxLifetime(time.Duration(MaxLifetime) * time.Second)
	mysql.SetMaxOpenConns(MaxOpenConns)
	mysql.SetMaxIdleConns(MaxIdleConns)
	db = mysql
	fmt.Println("Mysql 已連線")
	//db.Exec("insert into task_detail(id,description,status) values (?,?,?)", 12345678, "測試任務", "Pending")
}
func GetMysqlDB() *sql.DB {
	return db
}
