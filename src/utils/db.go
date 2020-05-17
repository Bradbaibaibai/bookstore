package utils

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var (
	Db *sql.DB
	err error
)

func Init(){
	Db,err = sql.Open("mysql","root:518518@tcp(localhost:3306)/bookstore")
	if err != nil{
		panic(err.Error())
	}
}