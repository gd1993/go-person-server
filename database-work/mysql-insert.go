package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"go-person-server/config"
)

var Db *sqlx.DB

func main() {
	// 调用输出化数据库的函数
	conn, err := config.MysqlConn()
	if err != nil && conn != nil {
		fmt.Printf("init db failed,err:%v\n", err)
	}
	Db = conn
	insertPersonDemo()
	insertPlaceDemo()
	conn.Close()
}

func insertPlaceDemo() {
	sqlStr := `insert into place(country,city,telcode) values ("china","beijing","1");`
	_, err := Db.Exec(sqlStr)
	if err != nil {
		fmt.Println("exec insert palce failed, ", err)
		return
	}
	fmt.Println("exec insert place success")

}

func insertPersonDemo() {
	sqlStr := `insert into person(username, sex, email)values("stu001", "man", "stu01@qq.com");`
	r, err := Db.Exec(sqlStr)
	if err != nil {
		fmt.Println("exec failed, ", err)
		return
	}
	theID, err := r.LastInsertId() // 新插入数据的id
	if err != nil {
		fmt.Printf("get lastinsert ID failed, err:%v\n", err)
		return
	}
	fmt.Printf("runoob_id is %d.\n", theID)
}
