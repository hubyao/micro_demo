package handler

import (
	"context"
	"github.com/gin-gonic/gin"
	"micro_demo/comm/logging"
	"micro_demo/comm/xhttp"
	"micro_demo/comm/xhttp/errno"
	pbUser "micro_demo/proto/user"
)

// 微信登陆
func WechatLogin(c *gin.Context) {
	req := wechatLoginReq{}
	rsp := wechatLoginRsq{}
	var err error

	// 绑定数据
	if err := c.ShouldBind(&req); err != nil {
		xhttp.SendJsonResponse(c, err, nil)
		return
	}

	rpcUserOauthLogin, err := UserPbClient.UserOauthLogin(context.Background(), &pbUser.UserOauthLoginReq{
		UserOauth: &pbUser.UserOauth{
			Platform:   "wechat",
			OpenId:     req.Openid,
			Unionid:    req.Unionid,
			Sex:        0,
			Name:       req.Name,
			Avatar:     req.Iconurl,
		},
	})

	if err != nil {
		logging.Logger().Error(err)
		xhttp.SendJsonResponse(c, errno.ErrUserLogin, rsp)
		return
	}

	if !rpcUserOauthLogin.BaseResponse.Success {
		logging.Logger().Error(err)
		xhttp.SendJsonResponse(c,
			errno.Errno{int(rpcUserOauthLogin.BaseResponse.Error.Code),
				rpcUserOauthLogin.BaseResponse.Error.Message} , rsp)
	}


	rsp.Token = rpcUserOauthLogin.Token
	rsp.Uid   = rpcUserOauthLogin.Uid
	xhttp.SendJsonResponse(c, err, rsp)
}

type wechatLoginReq struct {
	Openid   string `json:"openid" json:"openid"`
	Unionid  string `json:"unionid" json:"unionid"`
	Gender   string `json:"gender" json:"gender"`
	Name     string `json:"name" json:"name"`
	Iconurl  string `json:"iconurl" json:"iconurl"`
}

type wechatLoginRsq struct {
	Token string `json:"token"`
	Uid   uint64 `json:"uid"`
}