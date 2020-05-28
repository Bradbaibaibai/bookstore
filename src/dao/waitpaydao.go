package dao

import (
	"encoding/json"
	"github.com/garyburd/redigo/redis"
	"model"
	"strconv"
	"utils"
)

/*
waitpaydao包含对待支付数据库的操作。
*/


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

//redis重写:
type cartRedis struct {
	PayID string
	BookID int
	BookName string
	UserName string
	Num int
	Price float64
}

func AddWaitPay(username string,books []string)error{
	/*
		mysql版本:
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
	*/
		_,err := utils.Rs.Do("select","3")
		if err != nil{
			return err
		}else{
			payID := utils.CreateUUID()
			for _,v := range books{
				bookID,_ := strconv.Atoi(v)
				mc,_ := GetCart(username,bookID)
				tmpStruct := cartRedis{
					PayID:payID,
					BookID:mc.Bookid,
					BookName:mc.Bookname,
					UserName:mc.Username,
					Num:mc.Num,
					Price:mc.Price,
				}
				values,err := json.Marshal(tmpStruct)
				if err != nil{
					return err
				}else{
					_,err := utils.Rs.Do("select","2")
					if err != nil{
						return err
					}else{
						_,err := utils.Rs.Do("rpush",username,values)
						if err != nil{
							return err
						}
					}
				}
			}
		}
		return nil
}

func DelWaitPay(username string)error{
	/*
		sqlStr := "delete from waitpay where username = ?"
		_,err := utils.Db.Exec(sqlStr,username)
		if err != nil{
			return err
		}else{
			return nil
		}
	*/
	_,err := utils.Rs.Do("select","2")
	if err != nil{
		return err
	}else{
		_,err := utils.Rs.Do("del",username)
		if err != nil{
			return err
		}
	}
	return nil
}

func GetWaitPay(username string)([]*model.WaitPay,error){
	/*
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
	*/

	_,err := utils.Rs.Do("select","2")
	if err != nil{
		return nil,err
	}else{
		values,err := redis.Values(utils.Rs.Do("lrange",username,"0","-1"))
		if err != nil{
			return nil,err
		}else{
			waitPaySlice := make([]*model.WaitPay,0)
			for _,v := range values{
				tmp := &model.WaitPay{}
				tmpValue,_ := v.([]byte)
				if err != nil{
					return nil,err
				}else{
					tmpS := &cartRedis{
					}
					json.Unmarshal(tmpValue,tmpS)
					tmp.Price = tmpS.Price
					tmp.Num = tmpS.Num
					tmp.UserName = tmpS.UserName
					tmp.BookName = tmpS.BookName
					tmp.BookID = tmpS.BookID
					tmp.PayID = tmpS.PayID
					tmp.SumPrice = float64(tmp.Num) * tmp.Price
					waitPaySlice = append(waitPaySlice,tmp)
				}
			}
			return waitPaySlice,nil
		}
	}

}