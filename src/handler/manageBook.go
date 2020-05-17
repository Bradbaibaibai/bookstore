package handler

import (
	"dao"
	"fmt"
	"html/template"
	"model"
	"net/http"
	"strconv"
)

func GetAllTheBooks(w http.ResponseWriter,r *http.Request){
	books,_ := dao.GetAllTheBooks()
	t := template.Must(template.ParseFiles("views/pages/manager/book_manager.html"))
	t.Execute(w,books)
}

func AddBook(w http.ResponseWriter,r *http.Request){
	book := &model.Book{}
	book.Title = r.PostFormValue("title")
	book.Author = r.PostFormValue("author")
	book.Price,_ = strconv.ParseFloat(r.PostFormValue("price"),64)
	tmpSales,_ := strconv.ParseInt(r.PostFormValue("sales"),10,0)
	book.Sales = int(tmpSales)
	tmpStock,_ :=strconv.ParseInt(r.PostFormValue("stock"),10,0)
	book.Stock =  int(tmpStock)
	book.ImgPath = "/static/img/default.jpg"
	dao.AddBook(book)
	GetPageBooks(w,r)
}


func DeletBook(w http.ResponseWriter,r *http.Request){
	bookId := r.FormValue("bookId")
	id,_ := strconv.ParseInt(bookId,10,0)
	dao.DeletBook(int(id))
	GetPageBooks(w,r)
}

func UpdateBook(w http.ResponseWriter,r *http.Request){
	bookId := r.FormValue("bookId")
	book,_ := dao.GetBook(bookId)
	t := template.Must(template.ParseFiles("views/pages/manager/book_modify.html"))
	t.Execute(w,book)
}

func ModifyBook(w http.ResponseWriter,r *http.Request){
	book := &model.Book{}
	tmpID,_ := strconv.ParseInt(r.PostFormValue("bookId"),10,0)
	book.ID = int(tmpID)
	book.Title = r.PostFormValue("title")
	book.Author = r.PostFormValue("author")
	book.Price,_ = strconv.ParseFloat(r.PostFormValue("price"),64)
	tmpSales,_ := strconv.ParseInt(r.PostFormValue("sales"),10,0)
	book.Sales = int(tmpSales)
	tmpStock,_ :=strconv.ParseInt(r.PostFormValue("stock"),10,0)
	book.Stock =  int(tmpStock)
	book.ImgPath = "/static/img/default.jpg"
	dao.UpdateBook(book)
	UpdateBook(w,r)
}


func GetPageBooks(w http.ResponseWriter,r *http.Request){
	pageNo := r.FormValue("pageNo")
	page,_ := dao.GetPageBooks(pageNo)
	t := template.Must(template.ParseFiles("views/pages/manager/book_manager.html"))
	t.Execute(w,page)
}

func GetPageBooksByPrice(w http.ResponseWriter,r  *http.Request){

	pageNo := r.FormValue("pageNo")
	pageMin := r.PostFormValue("min")
	pageMax := r.PostFormValue("max")
	if pageMax == "" && pageMin == ""{
		pageMax = r.FormValue("max")
		pageMin = r.FormValue("min")
	}
	intMin,_ := strconv.Atoi(pageMin)
	intMax,_ := strconv.Atoi(pageMax)
	if pageMin == "" || pageMax == "" || intMin > intMax{
		IndexHandler(w,r)
	}else{
		page,_ := dao.GetPageBooksByPrice(pageNo,pageMin,pageMax)
		page.Max = pageMax
		page.Min = pageMin
		t := template.Must(template.ParseFiles("views/index.html"))
		cookie,err := r.Cookie("Session")
		if err == nil{
			s,err := dao.GetSession(cookie.Value)
			if err == nil{
				fmt.Println(s.UserName)
				page.Show.SetHasLogin()
				page.UserName = s.UserName
				if s.UserName == "admin"{
					page.Show.SetAdmin()
				}
			}
		}
		t.Execute(w,page)
	}
}