package dao

import (
	"fmt"
	"model"
	"strconv"
	"utils"
	)

func AddCart(username,bookid string) error {
	book,err := GetBook(bookid)
	if err != nil{
		return err
	}else{
		bid,_ := strconv.Atoi(bookid)
		cart,err := GetCart(username,bid)
		if err != nil{
			err = InsertCart(username,book.Title,bid,book.Price,1)
			if err != nil{
				return  err
			}else{
				return nil
			}
		}else{
			err = AddCount(username,bid,cart.Num)
			if err != nil{
				return  err
			}else{
				return nil
			}
		}
	}
}

func GetCarts(username string)([]*model.Cart,error){
	sqlStr := "select username,bookname,bookid,price,num from cart where username = ?"
	rows,err := utils.Db.Query(sqlStr,username)
	if err != nil{
		return nil,err
	}else{
		mdSlice := make([]*model.Cart,0)
		for rows.Next(){
			cartTmp := &model.Cart{}
			err := rows.Scan(&cartTmp.Username,&cartTmp.Bookname,&cartTmp.Bookid,&cartTmp.Price,&cartTmp.Num)
			mdSlice = append(mdSlice,cartTmp)
			if err != nil{
				return nil,err
			}
		}
		if err != nil{
			return nil,err
		}else{
			return mdSlice,nil
		}
	}
}

func GetCart(username string,bookid int)(*model.Cart,error){
	sqlStr := "select username,bookname,bookid,price,num from cart where username = ? and bookid = ?"
	row := utils.Db.QueryRow(sqlStr,username,bookid)
	cartTmp := &model.Cart{}
	err := row.Scan(&cartTmp.Username,&cartTmp.Bookname,&cartTmp.Bookid,&cartTmp.Price,&cartTmp.Num)
	if err != nil{
		return nil,err
	}else{
		return cartTmp,nil
	}
}

func InsertCart(username,bookname string,bookid int,price float64,num int)error{
	sqlStr := "select count(*) from cart where username = ?"
	var countOfCart int
	row := utils.Db.QueryRow(sqlStr,username)
	row.Scan(&countOfCart)
	if countOfCart < 9{
		sqlStr = "insert into cart values(?,?,?,?,?)"
		_,err := utils.Db.Exec(sqlStr,username,bookname,bookid,price,num)
		if err !=nil{
			return err
		}else{
			return nil
		}
	}else{
		return  fmt.Errorf("超范围")
	}
}

func AddCount(username string,bookid,num int)error{
	/*
	var num int
	sqlStr := "select num from cart where username = ?"
	row := utils.Db.QueryRow(sqlStr,username)
	row.Scan(&num)
	*/
	num++
	sqlStr := "update cart set num = ? where username =? and bookid = ? "
	_,err := utils.Db.Exec(sqlStr,num,username,bookid)
	if err !=nil{
		return err
	}else{
		return nil
	}
}

func DelCart(username string,bookid int)error{
	sqlStr := "delete from cart where username =? and bookid =?"
	_,err := utils.Db.Exec(sqlStr,username,bookid)
	if err != nil{
		return err
	}else{
		return nil
	}
}

