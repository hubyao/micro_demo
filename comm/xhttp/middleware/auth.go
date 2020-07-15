package middleware

import (
	"github.com/gin-gonic/gin"
	"micro_demo/comm/jwt"
	"micro_demo/comm/logging"
	"micro_demo/comm/xhttp"
	"micro_demo/comm/xhttp/errno"
)

type AuthSrv struct {
}

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		authorization := c.Request.Header.Get("Authorization")

		if authorization == "" {
			xhttp.FailRsp(c, errno.InvalidTokenError, "")
			c.Abort()
			return
		}

		// 截取token
		token := authorization[7:]
		// TODO: 测试环境绕过
		if token[:4] == "uid:"{
			// 存储
			c.Set("x-uid", token[4:])
			logging.Logger().Debugf("debug_info %v", token[4:])
			c.Next()
			return
		}

		// 解析token
		claims, err := jwt.ParseToken(token)
		if err != nil {
			xhttp.FailRsp(c, errno.InvalidTokenError, "")
			c.Abort()
			return
		}

		// 存储
		c.Set("x-uid", claims.UserId)


		c.Next()

	}
}

