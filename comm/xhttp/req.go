package xhttp

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

// GetUid 获取用户id
func  GetUid(c *gin.Context) uint64 {
	xUid, ok := c.Get("x-uid")
	if !ok {
		return 0
	}

	uid, _ := strconv.Atoi(xUid.(string))



	return uint64(uid)
}
