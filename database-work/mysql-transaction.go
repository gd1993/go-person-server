package main

import (
	"fmt"
	"go-person-server/config"
)

func main() {
	conn, err := config.MysqlConn()
	if err != nil && conn != nil {
		fmt.Println("mysql conn failed", err)
	}

	insertSql := `insert into place (country,city,telcode) values ("chain","shanghai","2");`
	dbconn, err := conn.Begin()
	if err != nil {
		fmt.Println("open transcation failed")
	}
	result, err := dbconn.Exec(insertSql)
	if err != nil {
		fmt.Println("insert failed", err)
		dbconn.Rollback()
		return
	}
	id, err := result.LastInsertId()
	if err != nil {
		fmt.Println("insert success , un commit", err)
	}
	fmt.Println("person   user_id:", id)
	dbconn.Commit()
}
