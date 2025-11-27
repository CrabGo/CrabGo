package v1

import "github.com/gogf/gf/v2/frame/g"

type AuthReq struct {
	g.Meta   `path:"/auth" method:"post"`
	UserName string `json:"user_name" type:"string" dc:"用户名" v:"required" `
	PassWord string `json:"pass_word" type:"string" dc:"密码" v:"required"`
}

type AuthRes struct {
	LoginTime int64  `json:"login_time"`
	TimeOut   int64  `json:"time_out"`
	Token     string `json:"token"`
}
