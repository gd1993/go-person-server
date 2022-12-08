package main

import (
	"fmt"
	"go-person-server/config"
)

func main() {
	conn, err := config.MysqlConn()
	if conn != nil && err != nil {
		fmt.Println("mysql conn failed", err)
	}
	sqlUpdate := "update person set username = ? where user_id = ?"
	update, err := conn.Exec(sqlUpdate, "hzh", 4)
	if err != nil {
		fmt.Println("update failed ", err)
	}

	row, err := update.RowsAffected()
	if err != nil {
		fmt.Println("rows failed", err)
	}
	fmt.Println("update rows:", row)

}
