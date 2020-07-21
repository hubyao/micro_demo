/*
 * @Author       : jianyao
 * @Date         : 2020-07-14 09:09:53
 * @LastEditTime : 2020-07-14 10:51:32
 * @Description  : file content
 */ 
 
package handler

import (
	"context"
	"github.com/gin-gonic/gin"
	"micro_demo/comm/logging"
	"micro_demo/comm/xhttp"
	"micro_demo/comm/xhttp/errno"
	pbUser "micro_demo/proto/user"
)

// WechatLogin 微信登陆
func WechatLogin(c *gin.Context) {
	req := wechatLoginReq{}
	rsp := wechatLoginRsq{}
	var err error
	
	// 绑定数据
	if err := c.ShouldBind(&req); err != nil {
		xhttp.FailRsp(c, err, "")
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
		xhttp.FailRsp(c, errno.ErrUserLogin, "")
		return
	}

	if !rpcUserOauthLogin.BaseResponse.Success {
		logging.Logger().Error(err)
		xhttp.FailRsp(c,
			errno.Errno{int(rpcUserOauthLogin.BaseResponse.Error.Code),
				rpcUserOauthLogin.BaseResponse.Error.Message} , 
				"")
				
		return
	}


	rsp.Token = rpcUserOauthLogin.Token
	rsp.Uid   = rpcUserOauthLogin.Uid
	xhttp.OkRsp(c, rsp)
}

type wechatLoginReq struct {
	Openid   string `json:"openid" form:"openid" binding:"required"`
	Unionid  string `json:"unionid" form:"unionid"`
	Gender   string `json:"gender" form:"gender"`
	Name     string `json:"name" form:"name"`
	Iconurl  string `json:"iconurl" form:"iconurl"`
}

type wechatLoginRsq struct {
	Token string `json:"token"`
	Uid   uint64 `json:"uid"`
}