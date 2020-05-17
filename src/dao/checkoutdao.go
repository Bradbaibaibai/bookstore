package dao

import (
	"fmt"
	"model"
	"utils"
)

/*
	payid varchar(100) not null,
    bookid int not null,
    bookname varchar(100) not null,
    username varchar(100) not null,
    num int not null,
    price float not null,
    consigadd varchar(100) not null,
    consigtel varchar(100) not null,
    consigname varchar(100) not null
*/

func Addpayed(mpays []*model.MyPay) error {
	sqlStr := "insert into payed values(?,?,?,?,?,?,?,?,?)"
	for _,v := range mpays{
		_,err := utils.Db.Exec(sqlStr,v.PayID,v.BookID,v.BookName,v.UserName,v.Num,v.Price,v.ConsigAdd,v.ConsigTel,v.ConsigName)
		if err != nil{
			return err
		}
	}
	return nil
}

func BookStockJn(mpays []*model.MyPay)error{
	//先对书的库存加锁再减,对mpays也要加锁
	var nums []int
	for _,v := range mpays{
		sqlStr := "select stock from books where id = ?"
		row := utils.Db.QueryRow(sqlStr,v.BookID)
		var num int
		err := row.Scan(&num)
		if num < v.Num{
			return fmt.Errorf("error")
		}
		if err != nil{
			return err
		}
		nums = append(nums,num)
	}
	i := 0
	for _,v := range mpays{
		sqlStr := "update books set stock=? where id=?"
		n := nums[i]
		i++
		n -= v.Num
		_,err := utils.Db.Exec(sqlStr,n,v.BookID)
		if err != nil{
			return err
		}
	}
	return nil
}

func GetMyPay(username string) ([]*model.MyPay,error){
	sqlStr := "select payid,bookid,bookname,username,num,price,consigadd,consigtel,consigname from payed where username = ?"
	rows,err := utils.Db.Query(sqlStr,username)
	if err != nil{
		return nil,err
	}else{
		var mpays []*model.MyPay
		for rows.Next(){
			tmp := &model.MyPay{}
			err := rows.Scan(&tmp.PayID,&tmp.BookID,&tmp.BookName,&tmp.UserName,&tmp.Num,&tmp.Price,&tmp.ConsigAdd,&tmp.ConsigTel,&tmp.ConsigName)
			mpays = append(mpays,tmp)
			if err != nil{
				return nil,err
			}
		}
		return mpays,nil
	}
}


func DelMyPay(payID string)error{
	sqlStr := "delete from payed where payid = ?"
	_,err := utils.Db.Exec(sqlStr,payID)
	if err != nil{
		return err
	}else{
		return nil
	}
}