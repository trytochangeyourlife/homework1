/*
* @Author: 李澳华
* @Date:   2021/6/6 16:15 
*/

package main

import (
	"database/sql"
	"fmt"
	"github.com/pkg/errors"
	_ "github.com/go-sql-driver/mysql"
)


func Init_sql() *sql.DB {
	db, err := sql.Open("mysql", "root:664720@tcp(127.0.0.1:3306)/go_db")
	if err != nil {
		panic(err)
	}
	fmt.Println("init success")
	return db
}


func Query_info(db *sql.DB) (error) {
	var karma int
	err := db.QueryRow("select id, name from users where id = 1").Scan(&karma)
	//fmt.Println(row)
	if err != nil {
		if err == sql.ErrNoRows {
			return errors.Wrap(err, "the data is nil")
		}else {
			return err
		}
	}
}

func main() {
	db := Init_sql()
	Query_info(db)
}
