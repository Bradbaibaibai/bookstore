package dao

import (
	"fmt"
	"testing"
	"utils"
)

func TestUser(t *testing.T){
	utils.Init()
//	t.Run("Insert",testInsertUserInfor)
//	t.Run("Check",testCheckUserNameAndPassword)
//	t.Run("Get",testGetUserNameAndPassword)
//	t.Run("Check Once",testCheckIfOnlyUsername)
}

func testInsertUserInfor(t *testing.T){
	InsertUserInfor("白云鹏","518518","123456@qq.com")
}

func testCheckUserNameAndPassword(t *testing.T){
	if err := CheckUserNameAndPassword("白云鹏","518518");err != nil{
		fmt.Println("not find")
	}else{
		fmt.Println("find it")
	}
}

func testGetUserNameAndPassword(t *testing.T){
	u,err := GetUserNameAndPassword("白云鹏","518518")
	if err != nil{
		fmt.Println("error")
	}else{
		fmt.Println(u.Username,":",u.Password,":",u.Email)
	}
}

func testCheckIfOnlyUsername(t *testing.T) {
	if err := CheckIfOnlyUsername("白云鹏");err != nil{
		fmt.Println("exist")
	}else{
		fmt.Println("not exist")
	}
}

