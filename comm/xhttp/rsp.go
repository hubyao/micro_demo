package xhttp

import (
	"log"
	"micro_demo/comm/xhttp/errno"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Response ...
type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// SendJsonResponse ...
func SendJsonResponse(c *gin.Context, err error, data interface{}) {
	code, message := errno.DecodeErr(err)
	SendResponse(c, code, message, data)
	return
}

//  SendResponse ...
func SendResponse(c *gin.Context, code int, message string, data interface{}) {
	resp := Response{
		Code:    code,
		Message: message,
		Data:    data,
	}
	// TODO: log
	log.Printf("uri:%s,SendResponse code:%d,message:%s,data:%+v\n", c.Request.RequestURI, code, message, data)
	c.JSON(http.StatusOK, resp)
}
