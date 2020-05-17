package model

type ConsigHtml struct {
	ConsigName string
	ConsigTel string
	ConsigAdd string
	Email string
	Username string
}

func (rec *ConsigHtml)ExistConsigName()bool{
	return rec.ConsigName != ""
}

func (rec *ConsigHtml)ExistConsigTel()bool{
	return rec.ConsigTel != ""
}

func (rec *ConsigHtml)ExistConsigAdd()bool{
	return rec.ConsigAdd != ""
}

func (rec *ConsigHtml)ExistEmail()bool{
	return rec.Email != ""
}


