package v1

import "github.com/gogf/gf/v2/frame/g"

type LoginReq struct {
	g.Meta   `path:"users/login" method:"post" sm:"Login" tags:"User"`
	Username string `json:"username" v:"required|length:3,12"`
	Password string `json:"password" v:"required|length:6,16"`
}

type LoginRes struct {
	Token string `json:"token" dc:"Add Authorization: token in header for authenticated endpoints"`
}
