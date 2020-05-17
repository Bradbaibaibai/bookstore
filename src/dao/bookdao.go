package dao

import (
	"model"
	"strconv"
	"utils"
)

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
	sqlStr := "insert into books(title,author,price,sales,stock,img_path) values(?,?,?,?,?,?)"
	_,err := utils.Db.Exec(sqlStr,b.Title,b.Author,b.Price,b.Sales,b.Stock,b.ImgPath)
	if err != nil{
		return err
	}
	return nil
}

func DeletBook(id int) error {
	sqlStr := "delete from books where id = ?"
	_,err := utils.Db.Exec(sqlStr,id)
	if err != nil{
		return  err
	}
	return nil
}

func GetBook(bookid string) (*model.Book,error){
	sqlStr := "select id,title,author,price,sales,stock,img_path from books where id = ?"
	row := utils.Db.QueryRow(sqlStr,bookid)
	tmpbook := &model.Book{}
	err := row.Scan(&tmpbook.ID,&tmpbook.Title,&tmpbook.Author,&tmpbook.Price,&tmpbook.Sales,&tmpbook.Stock,&tmpbook.ImgPath)
	if err != nil{
		return nil,err
	}
	return tmpbook,nil
}

func UpdateBook(newbook *model.Book) error{
	sqlStr := "update books set title=?,author=?,price=?,sales=?,stock=?,img_path=? where id=?"
	_,err := utils.Db.Exec(sqlStr,newbook.Title,newbook.Author,newbook.Price,newbook.Sales,newbook.Stock,newbook.ImgPath,newbook.ID)
	if err != nil{
		return err
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

