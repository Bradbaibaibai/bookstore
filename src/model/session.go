package model

import "utils"

type Session struct {
	SessionID string
	UserName string
	UserID int
}

func (rec *Session)InitSession(username string,userid int){

	rec.UserName = username
	rec.UserID = userid
	rec.SessionID = utils.CreateUUID()
}
