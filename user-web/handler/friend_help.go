package handler

import (
	"context"
	"github.com/gin-gonic/gin"
	"micro_demo/comm/logging"
	"micro_demo/comm/xhttp"
	"micro_demo/comm/xhttp/errno"
	pbUser "micro_demo/proto/user"
)

// FriendHelp 好友助力
func FriendHelp(c *gin.Context) {
	rsp := friendHelpRsp{}
	rsp.UserInfoList = []UserInfo{}

	uid := xhttp.GetUid(c)
	logging.Logger().Debugf("debug_info uid %v",uid)

	rpcData, err := UserPbClient.GetFriendHelpListByUser(context.Background(), &pbUser.GetFriendHelpListByUserReq{
		Uid: uid,
	})

	if err != nil {
		logging.Logger().Error(err)
		xhttp.FailRsp(c, errno.ErrGetFriendHelp, err.Error())
		return
	}
	if !rpcData.BaseResponse.Success {
		logging.Logger().Error(rpcData.BaseResponse.Error)
		xhttp.FailRsp(c, errno.ErrGetFriendHelp,"")
		return
	}

	for _, v := range rpcData.UserInfoList {
		rsp.UserInfoList = append(rsp.UserInfoList, UserInfo{
			Uid:    v.Uid,
			Avatar: v.Avatar,
			Nick:   v.Nick,
		})
	}

	xhttp.OkRsp(c,rsp)
	return
}

type friendHelpReq struct {

}

type friendHelpRsp struct {
	UserInfoList []UserInfo  `json:"user_info_list"`
}

type UserInfo struct {
	Uid    uint64 `json:"uid"`    // 用户id
	Avatar string `json:"avatar"` // 头像
	Nick   string `json:"nick"`   // 昵称
}
