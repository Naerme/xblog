package log_api

import (
	"blogx_server/common"
	"blogx_server/common/res"
	"blogx_server/global"
	"blogx_server/models"
	"blogx_server/models/enum"
	"blogx_server/service/log_service"
	"fmt"
	"github.com/gin-gonic/gin"
)

type LogApi struct {
}

type LogListRequest struct {
	common.PageInfo
	LogType     enum.LogType      `form:"logType"` // 日志类型 1 2 3
	Level       enum.LogLevelType `form:"level"`   // 日志级别 1 2 3
	UserID      uint              `form:"userID"`  // 用户id
	IP          string            `form:"ip"`
	LoginStatus bool              `form:"loginStatus"` //登录
	ServiceName string            `form:"serviceName"`
}

type LogListResponse struct { //自定义的响应
	models.LogModel
	UserNickname string `json:"userNickname"`
	UserAvatar   string `json:"userAvatar"`
}

func (LogApi) LogListView(c *gin.Context) {
	var cr LogListRequest
	fmt.Println(cr)
	err := c.ShouldBindQuery(&cr)
	if err != nil {
		res.FailWithError(err, c)
		return
	}

	list, count, err := common.ListQuery(models.LogModel{
		LogType:     cr.LogType,
		Level:       cr.Level,
		UserID:      cr.UserID,
		IP:          cr.IP,
		LoginStatus: cr.LoginStatus,
		ServiceName: cr.ServiceName,
	}, common.Options{
		PageInfo:     cr.PageInfo,
		Likes:        []string{"title"},
		Preloads:     []string{"UserModel"},
		DefaultOrder: "created_at desc",
	})

	var _list = make([]LogListResponse, 0)

	for _, logModel := range list {
		_list = append(_list, LogListResponse{
			LogModel:     logModel,
			UserNickname: logModel.UserModel.Nickname,
			UserAvatar:   logModel.UserModel.Avatar,
		})
	}

	res.OkWithList(_list, int(count), c)
	return
}

func (LogApi) LogReadView(c *gin.Context) {
	var cr models.IDRequest
	err := c.ShouldBindUri(&cr)
	if err != nil {
		res.FailWithError(err, c)
		return
	}
	var log models.LogModel
	err = global.DB.Take(&log, cr.ID).Error
	if err != nil {
		res.FailWithMsg("不存在的日志", c)
		return
	}
	if !log.IsRead {
		global.DB.Model(&log).Update("is_read", true)

	}
	res.OkWithMsg("日志读取成功", c)

}

func (LogApi) LogRemoveView(c *gin.Context) {
	var cr models.RemoveRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, c)
		return
	}
	log := log_service.GetLog(c)
	log.ShowRequest()
	log.ShowResponse()

	var logList []models.LogModel
	global.DB.Find(&logList, "id in ?", cr.IDList)

	if len(logList) > 0 {
		global.DB.Delete(&logList)
	}

	msg := fmt.Sprintf("日志删除成功，共删除%d条", len(logList))

	res.OkWithMsg(msg, c)

	//var cr models.RemoveRequest
	//err := c.ShouldBindJSON(&cr)
	//if err != nil {
	//	res.FailWithError(err, c)
	//}
	//
	//log := log_service.GetLog(c)
	//log.ShowRequest()
	//log.ShowResponse()
	//
	//var logList []models.LogModel
	//global.DB.Find(&logList, "id in ?", cr.IDList)
	//if len(logList) > 0 {
	//	global.DB.Delete(&logList)
	//}
	//msg := fmt.Sprintf("日志删除成功，共删除%d条", len(logList))
	//res.OkWithMsg(msg, c)
}
