package main

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"go-person-server/config"
)

type Person struct {
	UserId   int    `db:"user_id"`
	Username string `db:"username"`
	Sex      string `db:"sex"`
	Email    string `db:"email"`
}

type Place struct {
	id      int    `db:"id"`
	Country string `db:"country"`
	City    string `db:"city"`
	TelCode int    `db:"telcode"`
}

var db *sqlx.DB

func main() {
	conn, err := config.MysqlConn()
	if conn != nil && err != nil {
		fmt.Println("mysql conn failed", err)
	}
	db = conn

	var person Person

	sqlStr := "select user_id, username, sex, email from person where user_id > ?"

	rows, err := db.Query(sqlStr, 1)
	if err != nil {
		fmt.Println("query  failed", err)
	}
	defer func() {
		rows.Close()
	}()

	for rows.Next() {
		err := rows.Scan(&person.UserId, &person.Username, &person.Sex, &person.Email)
		if err != nil {
			fmt.Println("scn failed ", err)
		}
		fmt.Printf("person:%v\n", person)
	}

}
