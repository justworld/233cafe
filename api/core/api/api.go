package api

import (
	"github.com/gin-gonic/gin"
)

type Gin struct {
	C *gin.Context
}

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Result bool      `json:"result"`
	Data interface{} `json:"data"`
}

// Response setting gin.JSON
func (g *Gin) Response(data interface{}) {
	g.C.JSON(200, Response{
		Code: 200,
		Data: data,
		Result: true,
	})
	return
}

func (g *Gin) ErrResponse(errCode int, msg string) {
	g.C.JSON(200, Response{
		Code: errCode,
		Data: nil,
		Msg: msg,
		Result: false,
	})
	return
}