package model

import (
	"utils"

)

type CartHtml struct{
	BookName string
	BookID int
	Num int
	Price float64
	SumPrice float64
	Status string
}

func(rec *CartHtml)InitCartHtml(BookID,Num int){
	rec.BookID = BookID
	rec.Num = Num

	sqlStr := "select id,title,author,price,sales,stock,img_path from books where id = ?"
	row := utils.Db.QueryRow(sqlStr,BookID)
	bk := &Book{}
	row.Scan(&bk.ID,&bk.Title,&bk.Author,&bk.Price,&bk.Sales,&bk.Stock,&bk.ImgPath)

	rec.Price = bk.Price
	rec.BookName = bk.Title
	rec.SumPrice = float64(Num)*rec.Price
	if Num <= bk.Stock{
		rec.Status = "余额充足"
	}else{
		rec.Status = "余额不足"
	}
}