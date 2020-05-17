package dao

import (
	"model"
	"utils"
)

func GetConsigHtml(username string)*model.ConsigHtml{
	consigHtml := &model.ConsigHtml{}
	sqlStr := "select email from users where username = ?"
	row := utils.Db.QueryRow(sqlStr,username)
	row.Scan(&consigHtml.Email)
	sqlStr = "select consigname,consigtel,consigadd from consiginfor where username = ?"
	row2 := utils.Db.QueryRow(sqlStr,username)
	row2.Scan(&consigHtml.ConsigName,&consigHtml.ConsigTel,&consigHtml.ConsigAdd)
	consigHtml.Username = username
	return consigHtml
}

func ModifyConsigHtml(consightml *model.ConsigHtml,password string){
	if err := ExistConsig(consightml.Username);err != nil{
		sqlStr := "insert into consiginfor values(?,?,?,?)"
		_,err := utils.Db.Exec(sqlStr,consightml.ConsigName,consightml.ConsigTel,consightml.ConsigAdd,consightml.Username)
		if err == nil{
			if password != ""{
				sqlStr = "update users set password = ? email = ? where username = ?"
				utils.Db.Exec(sqlStr,password,consightml.Email,consightml.Username)
			}else{
				sqlStr = "update users set email = ? where username = ?"
				utils.Db.Exec(sqlStr,consightml.Email,consightml.Username)
			}
		}
	}else{
		lastConsigHtml := GetConsigHtml(consightml.Username)
		if consightml.ConsigTel == ""{
			consightml.ConsigTel = lastConsigHtml.ConsigTel
		}
		if consightml.ConsigAdd == ""{
			consightml.ConsigAdd = lastConsigHtml.ConsigAdd
		}
		if consightml.ConsigName == ""{
			consightml.ConsigName = lastConsigHtml.ConsigName
		}
		if consightml.Email == ""{
			consightml.Email = lastConsigHtml.Email
		}
		sqlStr := "update consiginfor set consigname=?,consigtel=?,consigadd=? where username = ?"
		_,err := utils.Db.Exec(sqlStr,consightml.ConsigName,consightml.ConsigTel,consightml.ConsigAdd,consightml.Username)
		if err == nil{
			if password != ""{
				sqlStr = "update users set password = ? email = ? where username = ?"
				utils.Db.Exec(sqlStr,password,consightml.Email,consightml.Username)
			}else{
				sqlStr = "update users set email = ? where username = ?"
				utils.Db.Exec(sqlStr,consightml.Email,consightml.Username)
			}
		}
	}
}

func ExistConsig(username string)error{
	sqlStr := "select username from consiginfor where username = ?"
	row := utils.Db.QueryRow(sqlStr,username)
	if err := row.Scan(&username);err != nil{
		return err
	}else{
		return nil
	}
}