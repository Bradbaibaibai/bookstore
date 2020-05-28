package handler

import (
	"dao"
	"html/template"
	"net/http"
)
func IndexHandler(w http.ResponseWriter,r *http.Request){
	pageNo := r.FormValue("pageNo")
	page,_ := dao.GetPageBooks(pageNo)
	t := template.Must(template.ParseFiles("views/index.html"))

	cookie,err := r.Cookie("Session")
	if err == nil{
		s,err := dao.GetSession(cookie.Value)
		if err == nil{
			page.Show.SetHasLogin()
			page.UserName = s.UserName
			if s.UserName == "admin"{
				page.Show.SetAdmin()
			}
		}
	}

	t.Execute(w,page)
}


