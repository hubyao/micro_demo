package xutil

import (
	"bytes"
	"fmt"

	"net/http"
)

const (
	Notify = "https://oapi.dingtalk.com/robot/send?access_token=7718741c249b6a4fa4e582f98e19efb628f8cabbea0e412ac310a5dafd29515d"
)

func AddNotifyDingDing(content string, url string) bool {
	go notifyDingDing(content, url)
	return true
}

func notifyDingDing(content string, url string) {
	formt := `
        {
            "msgtype": "text",
            "text": {
        		"content": "%s"
			},
			 "at": {
        		"isAtAll": false
    		}
        }`
	body := fmt.Sprintf(formt, "err:"+content)

	jsonValue := []byte(body)
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		return
	}
}
