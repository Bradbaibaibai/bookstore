package handler

import (
	"dao"
	"net/http"
)

func CancelPay(w http.ResponseWriter,r *http.Request){
	cookie,err := r.Cookie("Session")
	if err == nil{
		s,err := dao.GetSession(cookie.Value)
		if err != nil{
			LoginHandler(w,r)
		}else{
			username := s.UserName
			dao.DelWaitPay(username)
			CartPage(w,r)
		}
	}
}