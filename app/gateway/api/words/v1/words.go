package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

type CreateReq struct {
	g.Meta     `path:"words" method:"post" sm:"Create" tags:"Word"`
	Word       string `json:"word" v:"required|length:1,100" dc:"Word"`
	Definition string `json:"definition" v:"required|length:1,300" dc:"Word definition"`
}

type CreateRes struct {
}

type DetailReq struct {
	g.Meta `path:"words/{id}" method:"get" sm:"Details" tags:"Word"`
	Id     uint `json:"id" v:"required"`
}

type DetailRes struct {
	Id                 uint        `json:"id"`
	Word               string      `json:"word"`
	Definition         string      `json:"definition"`
	ExampleSentence    string      `json:"exampleSentence"`
	ChineseTranslation string      `json:"chineseTranslation"`
	Pronunciation      string      `json:"pronunciation"`
	CreatedAt          *gtime.Time `json:"createdAt"`
	UpdatedAt          *gtime.Time `json:"updatedAt"`
}
