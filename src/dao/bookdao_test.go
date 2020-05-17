package dao

import (
	"fmt"
	"model"
	"testing"
	"utils"
)

func TestBookDao(t *testing.T) {
	utils.Init()
	//t.Run("GetAllThebook",testGetAllTheBooks)
	//t.Run("AddBook",testAddBook)
	//t.Run("DeleteBook",testDeletBook)
	//t.Run("GetBook",testGetbook)
	//t.Run("Update",testUpdateBook)
	t.Run("GetPage",testGetPageBooks)
}

func testGetAllTheBooks(t *testing.T) {
	books,err := GetAllTheBooks()
	if err != nil{
		fmt.Println("error")
	}else{
		for _,v := range books{
			fmt.Println(v.Title,v.Author)
		}
	}
}

func testAddBook(t *testing.T) {
	book := &model.Book{
		Title:   "三国演义",
		Author:  "罗贯中",
		Price:   88.88,
		Sales:   100,
		Stock:   100,
		ImgPath: "/static/img/default.jpg",
	}
	AddBook(book)
}

func testDeletBook(t *testing.T) {
	id := 32
	DeletBook(id)
}

func testGetBook(t *testing.T) {
	book,_ := GetBook("12")
	fmt.Println(book.Title)
}

func testUpdateBook(t *testing.T) {
	book := &model.Book{
		ID: 31,
		Title:   "三国演义",
		Author:  "罗贯中",
		Price:   88.88,
		Sales:   100,
		Stock:   99,
		ImgPath: "/static/img/default.jpg",
	}
	UpdateBook(book)
}

func testGetPageBooks(t *testing.T) {
	books,_ := GetPageBooks("1")
	fmt.Println(*books)
	books,_ = GetPageBooks("2")
	fmt.Println(*books)
}