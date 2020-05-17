package handler

import (
	"dao"
	"html/template"
	"net/http"
)

func MyPay(w http.ResponseWriter,r *http.Request){
	cookie,err := r.Cookie("Session")
	if err == nil{
		s,err := dao.GetSession(cookie.Value)
		if err != nil{
			LoginHandler(w,r)
		}else{
			username := s.UserName
			mpays,err := dao.GetMyPay(username)
			if err != nil{
				IndexHandler(w,r)
			}else{
				t := template.Must(template.ParseFiles("views/pages/pay/mypay.html"))
				t.Execute(w,mpays)
			}
		}
	}
}

func CancelMyPay(w http.ResponseWriter,r *http.Request){
	cookie,err := r.Cookie("Session")
	if err == nil{
		_,err := dao.GetSession(cookie.Value)
		if err != nil{
			LoginHandler(w,r)
		}else{
			payID := r.FormValue("payid")
			err := dao.DelMyPay(payID)
			if err != nil{
				MyPay(w,r)
			}else{
				MyPay(w,r)
			}
		}
	}
}