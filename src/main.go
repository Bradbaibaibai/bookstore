package main

import (
	"handler"
	"net/http"
	"utils"
)

func main(){
	utils.Init()
	http.HandleFunc("/",handler.GetPageBooksByPrice)
	http.HandleFunc("/logout",handler.LogOut)
	http.HandleFunc("/login",handler.LoginHandler)
	http.HandleFunc("/regist",handler.RegistHandler)
	http.HandleFunc("/getpagebooks",handler.GetPageBooks)
	http.HandleFunc("/addbook",handler.AddBook)
	http.HandleFunc("/deletebook",handler.DeletBook)
	http.HandleFunc("/updatebook",handler.UpdateBook)
	http.HandleFunc("/modifybook",handler.ModifyBook)
	http.HandleFunc("/consiginfor",handler.ConsigInfor)
	http.HandleFunc("/addBookCart",handler.AddBookCart)
	http.HandleFunc("/cartpage",handler.CartPage)
	http.HandleFunc("/cart",handler.CartHandler)
	http.HandleFunc("/cancelpay",handler.CancelPay)
	http.HandleFunc("/modifyconsiginfor",handler.ModifyConsigInfor)
	http.HandleFunc("/getPageBooksByPrice",handler.GetPageBooksByPrice)
	http.HandleFunc("/checkUserName",handler.CheckUsernameAjaxHandler)
	http.HandleFunc("/checkout",handler.CheckOut)
	http.HandleFunc("/mypay",handler.MyPay)
	http.HandleFunc("/cancelmypay",handler.CancelMyPay)
	http.Handle("/views/static/",http.StripPrefix("/views/static/",http.FileServer(http.Dir("views/static/"))))
	http.Handle("/views/pages/",http.StripPrefix("/views/pages/",http.FileServer(http.Dir("views/pages/"))))
	http.ListenAndServe("localhost:5432",nil)
}
