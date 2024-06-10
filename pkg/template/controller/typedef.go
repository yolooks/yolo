package controller

var TYPEDEF_TPL = `package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type MyResponse struct {
	RequestID string      ` + "`" + `json:"requestid"` + "`" + ` // 请求ID
	Code      int         ` + "`" + `json:"code"` + "`" + `      // 自定义错误码
	Message   string      ` + "`" + `json:"message"` + "`" + `   // 响应消息
	Error     string      ` + "`" + `json:"error"` + "`" + `     // 异常错误
	ErrorMsg  string      ` + "`" + `json:"errormsg"` + "`" + `  // 异常错误说明
	Timestamp int         ` + "`" + `json:"timestamp"` + "`" + ` // 响应时间戳
	Data      interface{} ` + "`" + `json:"data"` + "`" + `      // 响应主题数据
}

func ResponseAPI(c *gin.Context, myResponse MyResponse) {
	c.JSON(http.StatusOK, myResponse)
}

func Response(c *gin.Context, code int, data interface{}) {
	c.JSON(code, data)
}
`
