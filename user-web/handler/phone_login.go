/*
 * @Author       : jianyao
 * @Date         : 2020-07-14 09:09:53
 * @LastEditTime : 2020-07-14 10:46:08
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

// PhoneLogin 手机号登陆
func PhoneLogin(c *gin.Context) {
	req := phoneLoginReq{}
	rsp := phoneLoginRsp{}
	var err error
	var uid uint64

	// 绑定数据
	if err := c.ShouldBind(&req); err != nil {
		xhttp.FailRsp(c, err,"")
		return
	}

	// TODO：对验证码进行校验



	rpcGetFromPhone, err := UserPbClient.GetFromPhone(context.Background(), &pbUser.GetFromPhoneReq{
		Phone: req.Phone,
	})
	if err != nil {
		logging.Logger().Error(err)
		xhttp.FailRsp(c, err, "")
		return
	}
	if !rpcGetFromPhone.BaseResponse.Success{
		logging.Logger().Error(rpcGetFromPhone.BaseResponse.Error)
		xhttp.FailRsp(c, errno.ErrUserLogin, "")
		return
	}


	if nil != rpcGetFromPhone.UserInfo{
		uid = rpcGetFromPhone.UserInfo.Uid
	}

	// 判断用户是否存在
	if 0 == rpcGetFromPhone.UserInfo.Uid {
		// TODO: 用户不存在进行注册
		rpcAddUser,err  := UserPbClient.AddUser(context.Background(), &pbUser.AddUserReq{
			Phone: req.Phone,
			Nick:  "用户",
		})
		if err !=nil{
			logging.Logger().Error(err)
			xhttp.FailRsp(c, err, "")
			return
		}
		if !rpcAddUser.BaseResponse.Success{
			logging.Logger().Error(rpcAddUser.BaseResponse.Error)
			xhttp.FailRsp(c, errno.ErrUserLogin, "")
			return
		}

		// 新用户
		rsp.IsNew = true
		uid = rpcAddUser.Uid
	}

	// 生成token
	rpcGenerateToken, err  := UserPbClient.GenerateToken(context.Background(), &pbUser.GenerateTokenReq{
		Uid: uid,
	})
	
	if err != nil {
		logging.Logger().Error(err)
		xhttp.FailRsp(c, err, "")
		return
	}
	
	if !rpcGenerateToken.BaseResponse.Success{
		logging.Logger().Error(rpcGenerateToken.BaseResponse.Error)
		xhttp.FailRsp(c, errno.ErrUserLogin, "")
		return
	}


	rsp.Token = rpcGenerateToken.Token
	rsp.Uid = uid
	xhttp.OkRsp(c, rsp)
}

type phoneLoginReq struct {
	Phone string `json:"phone"  binding:"required"` // 手机号
	Code  string `json:"code"  binding:"required"`   // 验证码
}

type phoneLoginRsp struct {
	Token string `json:"token"`
	Uid   uint64 `json:"uid"`
	IsNew bool  `json:"is_new"` // 新用户状态:true=新用户,false=老用户
}




