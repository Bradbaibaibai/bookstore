package main

import (
	"dao"
	"fmt"
	"handler"
	"net/http"
	"utils"
)

func main(){
	utils.Init()	//redis 和 mysql操作句柄的初始化
	err := dao.RsInit()
	if err != nil{
		fmt.Println("error")
	}
	/*
		页面请求
	*/
	http.HandleFunc("/",handler.GetPageBooksByPrice)	//商城首页
	http.HandleFunc("/updatebook",handler.UpdateBook)	 //修改图书页面
	http.HandleFunc("/cartpage",handler.CartPage)		 //购物车页面
	http.HandleFunc("/consiginfor",handler.ConsigInfor)	 //收件人，用户信息页面
	http.HandleFunc("/getPageBooksByPrice",handler.GetPageBooksByPrice)
	/*
		用户操作请求
	*/
	http.HandleFunc("/logout",handler.LogOut)			//注销请求
	http.HandleFunc("/login",handler.LoginHandler)		//登录请求
	http.HandleFunc("/regist",handler.RegistHandler)		//注册请求
	http.HandleFunc("/modifyconsiginfor",handler.ModifyConsigInfor)	//修改用户信息
	http.HandleFunc("/checkUserName",handler.CheckUsernameAjaxHandler)	//判断用户名是否存在的Ajax请求
	/*
		图书管理请求
	*/
	http.HandleFunc("/getpagebooks",handler.GetPageBooks) //获取图书
	http.HandleFunc("/addbook",handler.AddBook)			 //添加图书
	http.HandleFunc("/deletebook",handler.DeletBook)		 //删除图书
	http.HandleFunc("/modifybook",handler.ModifyBook)  	 //修改图书
	/*
		用户购物模块
	*/
	http.HandleFunc("/addBookCart",handler.AddBookCart)	 //添加到购物车
	http.HandleFunc("/cart",handler.CartHandler)			 //结算或删除商品请求
	http.HandleFunc("/cancelpay",handler.CancelPay)		 //取消支付请求
	http.HandleFunc("/checkout",handler.CheckOut)		 //支付请求 -- 下单/使用消息队列
	http.HandleFunc("/mypay",handler.MyPay)				 //查看我的订单
	http.HandleFunc("/cancelmypay",handler.CancelMyPay)   //取消订单
	/*
		静态文件请求
	*/
	http.Handle("/views/static/",http.StripPrefix("/views/static/",http.FileServer(http.Dir("views/static/"))))
	http.Handle("/views/pages/",http.StripPrefix("/views/pages/",http.FileServer(http.Dir("views/pages/"))))
	http.ListenAndServe("localhost:5432",nil)
}
