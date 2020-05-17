package dao

import (
	"model"
	"strconv"
	"utils"
)

/*
	payid varchar(100) not null,
    bookid int not null,
    bookname varchar(100) not null,
    username varchar(100) not null,
    num int not null,
    price float not null

*/


//func GetPageWaitPay(pageNo string)(*model.WaitPayHtml,error) {
//	var iPageNo int64Books
//	if pageNo == "" {
//		iPageNo = 1
//	} else {
//		iPageNo, _ = strconv.ParseInt(pageNo, 10, 0)
//	}
//	sqlStr := "select count(*) from waitpay"
//	var totalRecord int64
//	row := utils.Db.QueryRow(sqlStr)
//	row.Scan(&totalRecord)
//	var pageSize int64 = 4
//	var totalPageNo int64
//	if totalRecord%pageSize == 0 {
//		totalPageNo = totalRecord / pageSize
//	} else {
//		totalPageNo = totalRecord/pageSize + 1
//	}
//}


func AddWaitPay(username string,books []string)error{
	sqlStr := "insert into waitpay values(?,?,?,?,?,?)"
	payID := utils.CreateUUID()
	for _,v := range books{
		tmp,_ := strconv.Atoi(v)
		mc,_ := GetCart(username,tmp)
		_,err := utils.Db.Exec(sqlStr,payID,tmp,mc.Bookname,username,mc.Num,mc.Price)
		if err != nil{
			return err
		}
	}
	return nil
}

func DelWaitPay(username string)error{
	sqlStr := "delete from waitpay where username = ?"
	_,err := utils.Db.Exec(sqlStr,username)
	if err != nil{
		return err
	}else{
		return nil
	}
}

func GetWaitPay(username string)([]*model.WaitPay,error){
	sqlStr := "select payid,bookid,bookname,username,num,price from waitpay where username = ?"
	rows,err := utils.Db.Query(sqlStr,username)
	waitPaySlice := make([]*model.WaitPay,0)
	if err != nil{
		return nil,err
	}else{
		for rows.Next(){
			tmp:= &model.WaitPay{}
			err := rows.Scan(&tmp.PayID,&tmp.BookID,&tmp.BookName,&tmp.UserName,&tmp.Num,&tmp.Price)
			tmp.SumPrice = float64(tmp.Num) * tmp.Price
			waitPaySlice = append(waitPaySlice,tmp)
			if err !=nil{
				return nil,err
			}
		}
		return waitPaySlice,nil
	}
}