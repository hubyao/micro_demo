package middleware

import (
	"micro_demo/comm/logging"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"bytes"
	"strings"
)


// MAX_PRINT_BODY_LEN ...
const MAX_PRINT_BODY_LEN = 512



// DetailLogger ...
type detailLogger struct {
	gin.ResponseWriter
	bodyBuf *bytes.Buffer
}

func (w detailLogger) Write(b []byte) (int, error) {
	//memory copy here!
	w.bodyBuf.Write(b)
	return w.ResponseWriter.Write(b)
}

// DetailLogger 日志详情
func DetailLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		reqBody := "" 

		logging.Logger().Infof("url begin -> %v", c.Request.URL)


		// 解析resp
		respBody := ""
		var blw detailLogger
		blw = detailLogger{bodyBuf: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = blw
		
		c.Next()
	

		ioReqBody, _ := ioutil.ReadAll(c.Request.Body)
        if ioReqBody != nil {
			reqBody = string(ioReqBody)
        }

		data,_ := c.GetRawData()
	
        //把读过的字节流重新放到body
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(data))
		
		respBody = strings.Trim(blw.bodyBuf.String(), "\n")
		if len(respBody) > MAX_PRINT_BODY_LEN {
			respBody = respBody[:(MAX_PRINT_BODY_LEN - 1)]
		}

		logging.Logger().Infof("url -> %v", c.Request.URL)
		logging.Logger().Infof("req body  -> %v", reqBody)
		logging.Logger().Infof("resp data -> %v", respBody)


	}
}