package api

import (
	"fmt"
	"time"
	"yasi_audio/global"
	"yasi_audio/model"
	"yasi_audio/model/request"
	"yasi_audio/model/response"
	"yasi_audio/service"
	"yasi_audio/utils"

	"github.com/gin-gonic/gin"
)

func CreateRoom(c *gin.Context) {
	var R request.Audio_time
	_ = c.ShouldBindJSON(&R)
	err, _ := service.AcceptTask(R.RoomId, R.AcceptedBy)
	if err != nil {
		response.FailWithDetailed(err.Error(), "接受失败，请联系管理员", c)
		return
	}
	t, err := time.ParseInLocation("2006-01-02 15:04:05", R.BattleTime, time.Local)
	if err != nil {
		response.FailWithMessage("时间格式不对", c)
		return
	}
	fmt.Println(t, time.Now(), "===========================")
	toekn := utils.GenerateToken(R.RoomId, R.Founder)
	room := &model.Audio_time{Founder: R.Founder, AcceptedBy: R.AcceptedBy, RoomId: R.RoomId, BattleTime: t, Token: toekn}
	err, roomReturn := service.CreateRoom(*room)
	if err != nil {
		response.FailWithDetailed(err.Error(), "创建失败", c)
	} else {
		response.OkWithDetailed(roomReturn, "创建成功", c)
	}
}

type FindToken struct {
	RoomId string `json:"RoomId"`
}

func GetToken(c *gin.Context) {
	token := FindToken{}
	fmt.Println(c.Request.Header)
	err := c.ShouldBindJSON(&token)
	fmt.Println(err)
	var room model.Audio_time
	err = global.GVA_DB.Where("room_id = ?", token.RoomId).First(&room).Error
	if err != nil {
		response.FailWithDetailed(err.Error(), "失败", c)
		return
	}
	response.OkWithDetailed(room.Token, "成功", c)
}
