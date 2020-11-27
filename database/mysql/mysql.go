package mysql

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func initDB() (err error) {
	datasource := "liyang:ismeade@tcp(47.95.220.162:3306)/dev"
	db, err = sql.Open("mysql", datasource)
	if err != nil {
		return err
	}
	err = db.Ping()
	if err != nil {
		return err
	}
	db.SetMaxOpenConns(8)
	return nil
}

type Book struct {
	id     int
	title  string
	reader string
}

func queryMultiDemo() {
	var book Book
	sql := "SELECT id, title, reader FROM book where id > ?"
	rows, err := db.Query(sql, 0)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&book.id, &book.title, &book.reader)
		if err != nil {
			fmt.Printf("Scan failed, err:%#v\n", err)
			return
		}
		fmt.Printf("result: %#v\n", book)

	}
}

func queryRowDemo() {
	var book Book
	fmt.Println("connect success.")
	sql := "SELECT id, title, reader FROM book where id = 1"
	err := db.QueryRow(sql).Scan(&book.id, &book.title, &book.reader)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return
	}
	fmt.Printf("result: %#v\n", book)

}
