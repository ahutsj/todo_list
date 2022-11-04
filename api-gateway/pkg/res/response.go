package res

import (
	"api-gateway/pkg/e"
	"github.com/gin-gonic/gin"
)

// Response 基础序列化器
type Response struct {
	Status uint        `json:"Status"`
	Data   interface{} `json:"Data"`
	Msg    string      `json:"Msg"`
	Error  string      `json:"Error"`
}

// DataList data struct with total num
type DataList struct {
	Item  interface{} `json:"Item"`
	Total uint        `json:"Total"`
}

// TokenData data struct with token
type TokenData struct {
	User  interface{} `json:"User"`
	Token string      `json:"Token"`
}

func ginH(msgCode int, data interface{}) gin.H {
	return gin.H{
		"code": msgCode,
		"msg":  e.GetMsg(uint(msgCode)),
		"data": data,
	}

}
