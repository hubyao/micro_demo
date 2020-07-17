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

// deprecated
func SendJsonResponse(c *gin.Context, err error, data interface{}) {
	code, message := errno.DecodeErr(err)
	sendResponse(c, code, message, data)
	return
}

func FailRsp(c *gin.Context, err error, msg string) {
	code, message := errno.DecodeErr(err)
	if msg != "" {
		message = msg
	}

	sendResponse(c, code, message, nil)
	return
}

func OkRsp(c *gin.Context, data interface{}) {
	sendResponse(c, 0, "", data)
	return
}

//  SendResponse ...
func sendResponse(c *gin.Context, code int, message string, data interface{}) {
	resp := Response{
		Code:    code,
		Message: message,
		Data:    data,
	}
	// TODO: log
	log.Printf("uri:%s,SendResponse code:%d,message:%s,data:%+v\n", c.Request.RequestURI, code, message, data)
	c.JSON(http.StatusOK, resp)
}
