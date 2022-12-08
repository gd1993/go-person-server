package config

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

func MysqlConn() (db2 *sqlx.DB, error error) {
	dsn := "root:Hzh122427@tcp(127.0.0.1:3306)/go_person?charset=utf8mb4&parseTime=True"
	db, error = sqlx.Open("mysql", dsn)
	if error != nil {
		fmt.Println("mysql open failed", error)
		return nil, error
	}
	fmt.Println("mysql open success")
	return db, nil

}
