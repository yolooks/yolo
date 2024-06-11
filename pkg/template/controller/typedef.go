package controller

var TYPEDEF_TPL = `package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type MyResponse struct {
	Code      int         ` + "`" + `json:"code"` + "`" + `      // error code
	Message   string      ` + "`" + `json:"message"` + "`" + `   // error message
	Data      interface{} ` + "`" + `json:"data"` + "`" + `      // data
}

func ResponseAPI(c *gin.Context, myResponse MyResponse) {
	c.JSON(http.StatusOK, myResponse)
}

func Response(c *gin.Context, code int, data interface{}) {
	c.JSON(code, data)
}
`
