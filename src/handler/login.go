package handler

import (
	"dao"
	"html/template"
	"model"
	"net/http"
)

func LoginHandler(w http.ResponseWriter,r *http.Request){
	username := r.PostFormValue("username")
	password := r.PostFormValue("password")
	if err := dao.CheckUserNameAndPassword(username,password);err != nil{
		//用户信息与数据库验证出错
		t := template.Must(template.ParseFiles("views/pages/user/login.html"))
		t.Execute(w,"error")
	}else{
		u,_ := dao.GetUserNameAndPassword(username,password)
		s := &model.Session{}
		s.InitSession(username,u.ID)
		dao.AddSession(s)

		cookie := &http.Cookie{
			Name:       "Session",
			Value:      s.SessionID,
			HttpOnly: true,
		}

		http.SetCookie(w,cookie)

		t := template.Must(template.ParseFiles("views/pages/user/login_success.html"))
		t.Execute(w,username)

	}
}

func LogOut(w http.ResponseWriter,r *http.Request){
	c,_ := r.Cookie("Session")
	dao.DeleteSession(c.Value)
	GetPageBooksByPrice(w,r)
}