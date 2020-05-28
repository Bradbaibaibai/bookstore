package dao



import (
	"encoding/json"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"model"
	"utils"
)

/*
	Redis缓存session会话控制
*/
func AddSession(sess *model.Session) error {
	/*
	mysql版本:
		sqlStr := "insert into sessions values(?,?,?)"
		_,err := utils.Db.Exec(sqlStr,sess.SessionID,sess.UserName,sess.UserID)
		if err != nil{
			return err
		}
		return nil
	*/

	//改用redis:
	value,_ := json.Marshal(*sess)
	_,err := utils.Rs.Do("select","0")
	if err != nil{
		return err
	}
	n,err := utils.Rs.Do("setnx",sess.SessionID,value)
	if err != nil{
		return fmt.Errorf("error set session")
	}else if n == int64(0){
		return fmt.Errorf("error set session")
	}
	return nil
}

func DeleteSession(sessID string) error {
	/*
	mysql版本:
		sqlStr := "delete from sessions where session_id = ?"
		_,err := utils.Db.Exec(sqlStr,sessID)
		if err != nil{
			return err
		}else{
			return nil
		}
	*/
	_,err := utils.Rs.Do("select","0")
	if err != nil{
		return err
	}
	_,err = utils.Rs.Do("del",sessID)
	if err != nil{
		return err
	}
	return nil
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
	/*
	mysql版本:
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
	*/
	s := &model.Session{}
	_,err := utils.Rs.Do("select","0")
	if err != nil{
		return nil,err
	}
	value,err := redis.Bytes(utils.Rs.Do("get",sessID))
	if err != nil{
		return nil,err
	}else{
		err := json.Unmarshal(value,s)
		if err != nil{
			return nil,err
		}else{
			return s,err
		}
	}
}

