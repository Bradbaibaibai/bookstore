package model

//联系数据库的模型
type Book struct {
	ID int
	Title string
	Author string
	Price float64
	Sales int
	Stock int
	ImgPath string
}


//用于向html模板传的模型
//既可以传数据，又可以传方法
type Page struct{
	Books []*Book
	PageNo int64
	PageSize int64
	TotalPageNo int64
	TotalRecord int64
	Min string
	Max string
	Show IndexHtml
	UserName string
}

func (rec *Page)ifExistNext() bool{
	return rec.PageNo < rec.TotalPageNo
}

func (rec *Page)ifExistLast() bool{
	return rec.PageNo > 1
}

func (rec *Page)GetNextPageNo() int64{
	if rec.ifExistNext(){
		return rec.PageNo + 1
	}else{
		return 1
	}
}

func (rec *Page)GetLastPageNo() int64{
	if rec.ifExistLast(){
		return rec.PageNo - 1
	}else{
		return rec.TotalPageNo
	}
}

