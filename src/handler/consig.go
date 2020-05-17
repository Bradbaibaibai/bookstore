package handler

import (
	"dao"
	"html/template"
	"model"
	"net/http"
)


func ConsigInfor(w http.ResponseWriter,r *http.Request){
	cookie,err := r.Cookie("Session")
	if err == nil{
		s,err := dao.GetSession(cookie.Value)
		if err == nil{
			username := s.UserName
			consigHtml := dao.GetConsigHtml(username)
			t := template.Must(template.ParseFiles("views/pages/user/consig_infor.html"))
			t.Execute(w,consigHtml)
		}
	}else{
		//访问出现错误
	}
}


func ModifyConsigInfor(w http.ResponseWriter,r *http.Request) {
	cookie, err := r.Cookie("Session")
	if err == nil{
		s,err := dao.GetSession(cookie.Value)
		if err == nil {
			if password := r.PostFormValue("password");password == r.PostFormValue("repwd"){
				username := s.UserName
				consigHtml := &model.ConsigHtml{}
				consigHtml.ConsigName = r.PostFormValue("consigname")
				consigHtml.ConsigAdd = r.PostFormValue("consigadd")
				consigHtml.ConsigTel = r.PostFormValue("consigtel")
				consigHtml.Email = r.PostFormValue("email")
				consigHtml.Username = username
				dao.ModifyConsigHtml(consigHtml,password)
				if password != ""{
					LoginHandler(w,r)
				}else{
					ConsigInfor(w,r)
				}
			}
		}else{
			LoginHandler(w,r)
		}
	}else{
		LoginHandler(w,r)
	}
}


