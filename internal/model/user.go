package model

type UserAuthInput struct {
	UserName string
	PassWord string
}
type UserAuthOutput struct {
	LoginTime int64
	TimeOut   int64
	Token     string
}
