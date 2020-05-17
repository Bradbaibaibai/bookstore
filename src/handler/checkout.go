package handler

import (
	"dao"
	"html/template"
	"model"
	"net/http"
)

func CheckOut(w http.ResponseWriter,r *http.Request){
	//先判断是否填写了默认收货地址，收件人，收件人电话
	//根据订单判断库存是否足够，然后加锁，等待全部加锁后再减去库存，添加到订单表
	cookie,err := r.Cookie("Session")
	if err != nil{
		LoginHandler(w,r)
	}else{
		s,err := dao.GetSession(cookie.Value)
		if err != nil{
			LoginHandler(w,r)
		}else{
			existConsigInfor := dao.GetConsigHtml(s.UserName)
			if existConsigInfor.ConsigName != "" && existConsigInfor.ConsigAdd != "" && existConsigInfor.ConsigTel != ""{

				wpay,_ := dao.GetWaitPay(s.UserName)
				mpays := make([]*model.MyPay,0)
				for _,v := range wpay{
					tmp := &model.MyPay{}
					tmp.UserName = s.UserName
					tmp.ConsigTel = existConsigInfor.ConsigTel
					tmp.ConsigAdd = existConsigInfor.ConsigAdd
					tmp.ConsigName = existConsigInfor.ConsigName
					tmp.BookName = v.BookName
					tmp.Num = v.Num
					tmp.Price = v.Price
					tmp.BookID = v.BookID
					tmp.PayID = v.PayID
					mpays = append(mpays,tmp)
				}
				err = dao.BookStockJn(mpays)
				if err != nil{
					t := template.Must(template.ParseFiles("views/pages/pay/checkoutfail.html"))
					t.Execute(w,r)
				}else{
					err = dao.Addpayed(mpays)
					if err != nil{
						t := template.Must(template.ParseFiles("views/pages/pay/checkoutfail.html"))
						t.Execute(w,r)
					}else{
						dao.DelWaitPay(s.UserName)
						t := template.Must(template.ParseFiles("views/pages/pay/checkout.html"))
						t.Execute(w,r)
					}
				}
			}else{
				dao.DelWaitPay(s.UserName)
				ConsigInfor(w,r)
			}

		}
	}
}
