package handler

import (
	"context"
	"github.com/gin-gonic/gin"
	"micro_demo/comm/logging"
	"micro_demo/comm/xhttp"
	"micro_demo/comm/xhttp/errno"
	pbUser "micro_demo/proto/user"
)

// DailyTask 每日任务
func DailyTask(c *gin.Context) {

	rsp := make([]dailyTask, 0)
	var err error
	uid := xhttp.GetUid(c)

	rpcData, err := UserPbClient.GetDailyTaskList(context.Background(), &pbUser.GetDailyTaskListReq{
		Uid: uid,
	})
	if err != nil {
		logging.Logger().Error(err)
		xhttp.FailRsp(c, errno.ErrDailyTask, err.Error())
		return
	}
	if !rpcData.BaseResponse.Success {
		logging.Logger().Error(rpcData.BaseResponse.Error)
		xhttp.FailRsp(c, errno.ErrDailyTask, rpcData.BaseResponse.Error.Message)
		return
	}

	for _, v := range rpcData.DailyTaskList {
		rsp = append(rsp, dailyTask{
			DailyTaskId:    v.DailyTaskId,
			Logo:           v.Logo,
			Title:          v.Title,
			ExpNum:         v.ExpNum,
			EnergyNum:      v.EnergyNum,
			CompleteStatus: v.CompleteStatus,
		})
	}

	xhttp.OkRsp(c, rsp)

}

type dailyTaskReq struct {
}

type dailyTaskRsp []dailyTask

type dailyTask struct {
	DailyTaskId    uint64 `json:"daily_task_id"`   // 任务id
	Logo           string `json:"logo"`            // logo
	Title          string `json:"title"`           // 标题
	ExpNum         int64  `json:"exp_num"`         // 经验数
	EnergyNum      int64  `json:"energy_num"`      // 能量数
	CompleteStatus bool   `json:"complete_status"` // 完成状态
}
