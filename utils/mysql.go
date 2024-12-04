package utils

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
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

type Mysql struct {
	UserName string `yaml:"userName"`
	Password string `yaml:"password"`
	Addr     string `yaml:"addr"`
	Port     string `yaml:"port"`
	TaleName string `yaml:"tableName"`
}

var (
	db            *gorm.DB
	sqlConnection string
)

//var db *sql.DB

func init() {
	file, e := ioutil.ReadFile("env.yaml")
	if e != nil {
		log.Fatal(e)
	}
	var mysql []Mysql
	err2 := yaml.Unmarshal(file, &mysql)
	if err2 != nil {
		log.Fatal(err2)
	}
	sqlConnection = mysql[0].UserName + ":" + mysql[0].Password + "@tcp(" + mysql[0].Addr + ":" + mysql[0].Port + ")/" + mysql[0].TaleName
	//fmt.Println(sqlConnection)
	//開啟資料庫連接
	var err error
	db, err = gorm.Open("mysql", sqlConnection)
	if err != nil {
		panic("failed to connect database")
	}

}

func GetMysqlDB() *gorm.DB {
	return db
}
