package handler

import (
	"dao"
	"html/template"
	"model"
	"net/http"
)

//用于处理注册信息验证
func RegistHandler(w http.ResponseWriter,r *http.Request){
	 username := r.PostFormValue("username")
	 //testing : 通过表单粗糙判断用户名是否重复
	 if err := dao.CheckIfOnlyUsername(username);err != nil{
	 	t := template.Must(template.ParseFiles("views/pages/user/regist.html"))
		t.Execute(w,nil)
	 }else{
		 password := r.PostFormValue("password")
		 email := r.PostFormValue("email")
		 dao.InsertUserInfor(username,password,email)
		 t := template.Must(template.ParseFiles("views/pages/user/regist_success.html"))

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

		 t.Execute(w,username)

	 }
}

//用于处理判断用户名称重复的异步请求
func CheckUsernameAjaxHandler(w http.ResponseWriter,r *http.Request){
	username := r.PostFormValue("username")
	if err := dao.CheckIfOnlyUsername(username);err != nil{
			w.Write([]byte("用户名已存在"))
	}else{
			w.Write([]byte("用户名可用"))
	}
}