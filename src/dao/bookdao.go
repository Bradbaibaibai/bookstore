package dao

import (
	"encoding/json"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"model"
	"strconv"
	"utils"
)

//从mysql中同步redis数据库
func RsInit() error {
	sqlStr := "select id,title,author,price,sales,stock,img_path from books"
	rows,err := utils.Db.Query(sqlStr)
	if err != nil{
		return err
	}
	for rows.Next(){
		tmpbook := new (model.Book)
		rows.Scan(&tmpbook.ID,&tmpbook.Title,&tmpbook.Author,&tmpbook.Price,&tmpbook.Sales,&tmpbook.Stock,&tmpbook.ImgPath)
		value,err := json.Marshal(*tmpbook)
		if err != nil{
			return err
		}
		_,err = utils.Rs.Do("select","1")
		if err != nil{
			return err
		}
		n,err := utils.Rs.Do("set",tmpbook.ID,value)
		if err != nil{
			return err
		}
		if n == int64(0){
			return fmt.Errorf("error from init redis")
		}
	}
	return nil
}


func GetAllTheBooks() ([]*model.Book,error){
		sqlStr := "select id,title,author,price,sales,stock,img_path from books"
		rows,err := utils.Db.Query(sqlStr)
		if err != nil{
			return nil,err
		}
		var books []*model.Book
		for rows.Next(){
			tmpbook := new (model.Book)
			rows.Scan(&tmpbook.ID,&tmpbook.Title,&tmpbook.Author,&tmpbook.Price,&tmpbook.Sales,&tmpbook.Stock,&tmpbook.ImgPath)
			books = append(books,tmpbook)
		}
		return books,nil
}


func AddBook(b *model.Book) error {
	//添加图书：
	/*
		需要先将书的信息放入到mysql数据库,
		再通过mysql数据库查询得到bookID,
		再加入redis缓存.
	*/
	sqlStr := "insert into books(title,author,price,sales,stock,img_path) values(?,?,?,?,?,?)"
	_,err := utils.Db.Exec(sqlStr,b.Title,b.Author,b.Price,b.Sales,b.Stock,b.ImgPath)
	if err != nil{
		return err
	}else{
		sqlStr = "select id from books where title = ? and author = ?"
		row := utils.Db.QueryRow(sqlStr,b.Title,b.Author)
		err = row.Scan(&b.ID)
		if err != nil{
			return err
		}else{
			value,err := json.Marshal(*b)
			if err != nil{
				return err
			}else{
				_,err = utils.Rs.Do("select","1")
				if err != nil{
					return err
				}
				utils.Rs.Do("SETNX",b.ID,value)
			}
		}
	}
	return nil
}

func DeletBook(id int) error {
	sqlStr := "delete from books where id = ?"
	_,err := utils.Db.Exec(sqlStr,id)
	if err != nil{
		return  err
	}else{
		_,err = utils.Rs.Do("select","1")
		if err != nil{
			return err
		}else{
			_,err = utils.Rs.Do("del",id)
			if err != nil{
				return err
			}
		}
	}
	return nil
}

func GetBook(bookid string) (*model.Book,error){
	/*mysql版本:
		sqlStr := "select id,title,author,price,sales,stock,img_path from books where id = ?"
		row := utils.Db.QueryRow(sqlStr,bookid)
		tmpbook := &model.Book{}
		err := row.Scan(&tmpbook.ID,&tmpbook.Title,&tmpbook.Author,&tmpbook.Price,&tmpbook.Sales,&tmpbook.Stock,&tmpbook.ImgPath)
		if err != nil{
			return nil,err
		}
		return tmpbook,nil
	*/
	_,err := utils.Rs.Do("select","1")
	if err != nil{
		return nil,err
	}
	value,err := redis.Bytes(utils.Rs.Do("get",bookid))
	if err != nil{
		//redis缓存未查询到再去查询mysql
		sqlStr := "select id,title,author,price,sales,stock,img_path from books where id = ?"
		row := utils.Db.QueryRow(sqlStr,bookid)
		tmpbook := &model.Book{}
		err := row.Scan(&tmpbook.ID,&tmpbook.Title,&tmpbook.Author,&tmpbook.Price,&tmpbook.Sales,&tmpbook.Stock,&tmpbook.ImgPath)
		if err != nil{
			return nil,err
		}
		return tmpbook,nil
	}else{
		tmpbook := &model.Book{}
		err  = json.Unmarshal(value,tmpbook)
		if err != nil{
			return nil,err
		}else{
			return tmpbook,err
		}
	}
}

func UpdateBook(newbook *model.Book) error{
	//先更新mysql中的图书信息
	sqlStr := "update books set title=?,author=?,price=?,sales=?,stock=?,img_path=? where id=?"
	_,err := utils.Db.Exec(sqlStr,newbook.Title,newbook.Author,newbook.Price,newbook.Sales,newbook.Stock,newbook.ImgPath,newbook.ID)
	if err != nil{
		return err
	}else{
		//再更新redis中的图书信息
		_,err := utils.Rs.Do("select","1")
		if err != nil{
			return err
		}else{
			value,err := json.Marshal(*newbook)
			if err != nil{
				return err
			}else{
				_,err := utils.Rs.Do("set",newbook.ID,value)
				if err != nil{
					return err
				}
			}
		}
	}
	return nil
}

func GetPageBooks(pageNo string)(*model.Page,error){
	var iPageNo int64
	if pageNo == ""{
		iPageNo = 1
	}else{
		iPageNo,_ = strconv.ParseInt(pageNo,10,0)
	}
	sqlStr := "select count(*) from books"
	var totalRecord int64
	row := utils.Db.QueryRow(sqlStr)
	row.Scan(&totalRecord)
	var pageSize int64 = 4
	var totalPageNo int64
	if totalRecord % pageSize == 0{
		totalPageNo = totalRecord/pageSize
	}else{
		totalPageNo = totalRecord/pageSize + 1
	}
	sqlStr2 := "select id,title,author,price,sales,stock,img_path from books limit ?,?"
	rows,err := utils.Db.Query(sqlStr2,(iPageNo - 1)*pageSize,pageSize)
	if err != nil{
		return nil,err
	}
	var books []*model.Book
	for rows.Next(){
		book := &model.Book{}
		rows.Scan(&book.ID,&book.Title,&book.Author,&book.Price,&book.Sales,&book.Stock,&book.ImgPath)
		books = append(books,book)
	}
	page := &model.Page{}
	page.Books = books
	page.PageNo = iPageNo
	page.PageSize = pageSize
	page.TotalPageNo = totalPageNo
	page.TotalRecord = totalRecord
	return page,nil
}


func GetPageBooksByPrice(pageNo string,min string,max string)(*model.Page,error){
	var iPageNo int64
	if pageNo == ""{
		iPageNo = 1
	}else{
		iPageNo,_ = strconv.ParseInt(pageNo,10,0)
	}
	sqlStr := "select count(*) from books where price < ? and price > ?"
	var totalRecord int64
	row := utils.Db.QueryRow(sqlStr,max,min)
	row.Scan(&totalRecord)
	var pageSize int64 = 4
	var totalPageNo int64
	if totalRecord % pageSize == 0{
		totalPageNo = totalRecord/pageSize
	}else{
		totalPageNo = totalRecord/pageSize + 1
	}
	sqlStr2 := "select id,title,author,price,sales,stock,img_path from books where price < ? and price > ? limit ?,?"
	rows,err := utils.Db.Query(sqlStr2,max,min,(iPageNo - 1)*pageSize,pageSize)
	if err != nil{
		return nil,err
	}
	var books []*model.Book
	for rows.Next(){
		book := &model.Book{}
		rows.Scan(&book.ID,&book.Title,&book.Author,&book.Price,&book.Sales,&book.Stock,&book.ImgPath)
		books = append(books,book)
	}
	page := &model.Page{}
	page.Books = books
	page.PageNo = iPageNo
	page.PageSize = pageSize
	page.TotalPageNo = totalPageNo
	page.TotalRecord = totalRecord
	return page,nil
}

