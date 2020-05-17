package model

type IndexHtml struct {
	Admin bool
	HasLogin bool
}

func (rec *IndexHtml)Show(showIt string) bool{
	if showIt == "true"{
		return true
	}else{
		return false
	}
}

func (rec *IndexHtml)IfAdmin() bool {
	return rec.Admin
}

func (rec *IndexHtml)IfHasLogin() bool{
	return rec.HasLogin
}

func (rec *IndexHtml)SetAdmin(){
	rec.Admin = true
}

func (rec *IndexHtml)SetHasLogin(){
	rec.HasLogin = true
}
