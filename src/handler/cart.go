package handler

import (
	"dao"
	"html/template"
	"model"
	"net/http"
	"strconv"
)

//添加购物车
/*
1.没有购物车，则创建
2.有购物车则判断数量是否已满

*/
func AddBookCart(w http.ResponseWriter,r *http.Request){
	cookie,err := r.Cookie("Session")
	if err != nil{
		//返回给ajax异步请求
		w.Write([]byte("请先登录！"))
	}else{
		s,err := dao.GetSession(cookie.Value)
		if err != nil{
			w.Write([]byte("请先登录！"))
		}else{
			bookId := r.PostFormValue("bookId")
			err := dao.AddCart(s.UserName,bookId)
			if err != nil{
				w.Write([]byte("添加购物车失败,余货不足! "))
			}else{
				w.Write([]byte("添加购物车成功! "))
			}
		}
	}
}

func CartPage(w http.ResponseWriter,r *http.Request){
	t := template.Must(template.ParseFiles("views/pages/cart/cart.html"))
	cookie,err := r.Cookie("Session")
	if err != nil{
		LoginHandler(w,r)
	}else{
		s,err := dao.GetSession(cookie.Value)
		if err!=nil{
			LoginHandler(w,r)
		}else{
			carts,err := dao.GetCarts(s.UserName)
			if err !=nil{
				LoginHandler(w,r)
			}else{
				cartHtmlSlice := make([]*model.CartHtml,0)
				for _,v := range carts{
					cartHtmlTmp := &model.CartHtml{}
					cartHtmlTmp.InitCartHtml(v.Bookid,v.Num)
					cartHtmlSlice = append(cartHtmlSlice,cartHtmlTmp)
				}
				t.Execute(w,cartHtmlSlice)
			}
		}
	}
}

func CartHandler(w http.ResponseWriter,r *http.Request){
	cookie,err := r.Cookie("Session")
	if err != nil{
		LoginHandler(w,r)
	}else{
		s,err := dao.GetSession(cookie.Value)
		if err!=nil{
			LoginHandler(w,r)
		}else{
			choose := r.PostFormValue("submit")
			if choose == "结算"{
				username := s.UserName
				r.ParseForm()
				mp := r.PostForm
				books := mp["books"]
				if books != nil{
					err := dao.AddWaitPay(username,books)
					if err != nil{
						CartHandler(w,r)
					}else{
						//去结算页面
						waitPays,err := dao.GetWaitPay(username)
						if err != nil{
							CartHandler(w,r)
						}else{
							waitPayHtml := &model.WaitPayHtml{}
							waitPayHtml.WaitPaySlice = waitPays
							for _,v := range  waitPays{
								waitPayHtml.SumNum += v.Num
								waitPayHtml.SumPrice += v.SumPrice
							}

							t := template.Must(template.ParseFiles("views/pages/pay/waitpay.html"))
							t.Execute(w,waitPayHtml)
						}
					}
				}else{
					CartHandler(w,r)
				}
			}else if choose == "删除"{
				username := s.UserName
				r.ParseForm()
				mp := r.PostForm
				books := mp["books"]
				for _,v := range books{
					tmp,_ := strconv.Atoi(v)
					dao.DelCart(username,tmp)
				}
				CartPage(w,r)
			}else{
				LoginHandler(w,r)
			}
		}
	}

}
