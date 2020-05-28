package utils

import (
	"database/sql"
	"fmt"
	"github.com/garyburd/redigo/redis"
	_ "github.com/go-sql-driver/mysql"
)

var (
	Db *sql.DB
	err error
	Rs redis.Conn
)

func Init() error {
	Db,err = sql.Open("mysql","root:518518@tcp(localhost:3306)/bookstore")
	if err != nil{
		fmt.Println("error dbInit mysql")
		return fmt.Errorf("error mysql connerct")
	}
	Rs,err = redis.Dial("tcp","localhost:6379")
	if err != nil{
		fmt.Println("error dbInit redis")
		return fmt.Errorf("error redis connerct")
	}
//	fmt.Println("数据库初始化完成")
	return nil
}