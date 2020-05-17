package dao

import (
	"fmt"
	"model"
	"utils"
)

//用于从数据库中获取到用户的信息
func GetUserNameAndPassword(username string , password string) (*model.User,error) {
	sqlStr := "select id,username,password,email from users where username = ? and password = ?"
	row := utils.Db.QueryRow(sqlStr,username,password)
	user := &model.User{}
	err := row.Scan(&user.ID,&user.Username,&user.Password,&user.Email)
	if err != nil{
		return nil,err
	}else{
		return user,nil
	}
}

//从用户数据库中核对用户名是否唯一:
func CheckIfOnlyUsername(username string) error {
	sqlStr := "select id from users where username = ?"
	row := utils.Db.QueryRow(sqlStr,username)
	user := &model.User{}
	err := row.Scan(&user.ID)
	if err != nil{
		return nil
	}else{
		return fmt.Errorf("exist")
	}
}

//从用户数据库中验证用户信息
func CheckUserNameAndPassword(username string, password string) error{
	sqlStr := "select id from users where username = ? and password = ?"
	row := utils.Db.QueryRow(sqlStr,username,password)
	user := &model.User{}
	err := row.Scan(&user.ID)
	if err != nil{
		return err
	}else{
		return nil
	}
}

//添加用户信息
func InsertUserInfor(username string,password string,email string)error{
	sqlStr := "insert into users(username,password,email) values(?,?,?)"
	_,err := utils.Db.Exec(sqlStr,username,password,email)
	return err
}


//修改密码
func ModifyPassword(username string,pwd string)error{
		sqlStr := "update users set password = ? where username = ?"
		_,err := utils.Db.Exec(sqlStr,pwd,username)
		return err
}

//修改邮箱
func ModifyEmail(username string,email string)error{
	sqlStr := "update users set email = ? where username = ?"
	_,err := utils.Db.Exec(sqlStr,email,username)
	return err
}