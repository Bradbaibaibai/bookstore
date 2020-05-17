package model
/*
	payid varchar(100) not null,
    bookid int not null,
    bookname varchar(100) not null,
    username varchar(100) not null,
    num int not null,
    price float not null

*/

type WaitPay struct {
	PayID string
	BookID int
	BookName string
	UserName string
	Num int
	Price float64
	SumPrice float64
}