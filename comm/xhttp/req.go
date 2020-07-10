package xhttp

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/gin-gonic/gin"
)

func GetQueryJsonObject(c *gin.Context, query interface{}) {
	buf, err := GetPostJsonData(c)
	if err != nil {
		log.Printf("Invalid body[%v] err: %v\n", string(buf), err)
	}

	if len(buf) != 0 {
		err = json.Unmarshal(buf, query)
		if err != nil {
			PostForm := c.Request.PostForm
			Form := c.Request.Form
			log.Printf("Invalid body[%v] PostForm[%v] Form[%v], err: %v\n", string(buf), PostForm, Form, err)

		}
	}
	return
}

func GetPostJsonData(c *gin.Context) ([]byte, error) {
	body, exists := c.Get("body")
	buf := []byte{}
	var err error
	if !exists {
		raw_body := c.Request.Body
		buf, err = ioutil.ReadAll(raw_body)
		c.Set("body", buf)
	} else {
		buf = body.([]byte)
	}
	return buf, err
}
