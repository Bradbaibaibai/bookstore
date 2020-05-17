package model

/*
create table payed(
	payid varchar(100) not null,
    bookid int not null,
    bookname varchar(100) not null,
    username varchar(100) not null,
    num int not null,
    price float not null,
    consigadd varchar(100) not null,
    consigtel varchar(100) not null,
    consigname varchar(100) not null
);
*/

type MyPay struct {
	PayID string
	BookID int
	BookName string
	UserName string
	Num int
	Price float64
	ConsigAdd string
	ConsigTel string
	ConsigName string
}