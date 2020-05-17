package dao

import (
	"model"
	"utils"
)

func AddSession(sess *model.Session) error {
	sqlStr := "insert into sessions values(?,?,?)"
	_,err := utils.Db.Exec(sqlStr,sess.SessionID,sess.UserName,sess.UserID)
	if err != nil{
		return err
	}
	return nil
}

func DeleteSession(sessID string) error {
	sqlStr := "delete from sessions where session_id = ?"
	_,err := utils.Db.Exec(sqlStr,sessID)
	if err != nil{
		return err
	}else{
		return nil
	}
}

/*
数据库里的session模型:
create table sessions(
session_id varchar(100) primary key,
username varchar(100) not null,
user_id int not null,
foreign key(user_id) references users(id)
)
*/

func GetSession(sessID string)(*model.Session,error){
	s := &model.Session{}
	sql := "select username,user_id from sessions where session_id = ?"
	row := utils.Db.QueryRow(sql,sessID)
	err := row.Scan(&s.UserName,&s.UserID)
	if err != nil{
		return nil,err
	}else{
		s.SessionID = sessID
		return s,nil
	}
}

