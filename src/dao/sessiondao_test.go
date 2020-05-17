package dao

import (
	"fmt"
	"model"
	"testing"
	"utils"
)

func TestSession(t *testing.T){
	utils.Init()
	fmt.Println("测试session")
	//t.Run("AddSession",testAddSession)
	//t.Run("GetSession",testGetSession)
	//t.Run("DelSession",testDeleteSession)
}

func testAddSession(t *testing.T) {
	s := &model.Session{
		SessionID:utils.CreateUUID(),
		UserName:"baibaibai",
		UserID:1,
	}
	err := AddSession(s)
	if err != nil{
		fmt.Println("error")
	}
}


func testGetSession(t *testing.T) {
	s,_ := GetSession("37f96df1-1800-480a-7816-b3c6a83abaad")
	fmt.Println(s.UserID,s.UserName,s.SessionID)
}



func testDeleteSession(t *testing.T) {
	err := DeleteSession("37f96df1-1800-480a-7816-b3c6a83abaad")
	if err != nil{
		fmt.Println("error")
	}
}
