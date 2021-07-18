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
	int64time := t.Unix() - time.Now().Unix() - 30
	if int64time < 0 {
		fmt.Println("时间已过期")
		fmt.Println(int64time)
		response.FailWithMessage("时间已过期", c)
		return
	}
	time_gap := time.Duration(int64time)
	fmt.Println(time_gap)
	toekn := utils.GenerateToken(R.RoomId, R.Founder)
	room := &model.Audio_time{Founder: R.Founder, AcceptedBy: R.AcceptedBy, RoomId: R.RoomId, BattleTime: t, Token: toekn}
	response.OkWithMessage("创建成功", c)
	time.AfterFunc(time.Second*time_gap, func() {
		fmt.Println("timeafter后执行了！！！")
		service.CreateRoom(*room)
	})
}

type FindToken struct {
	RoomId string `json:"RoomId"`
}

type TokenResponst struct {
	Token   string `json:"token"`
	Teacher string `json:"teacher"`
	Student string `json:"student"`
}

func GetToken(c *gin.Context) {
	token := FindToken{}
	fmt.Println(c.Request.Header)
	fmt.Println(token, "===========")
	err := c.ShouldBindJSON(&token)
	fmt.Println(err)
	fmt.Println(token.RoomId, "===========")
	var room model.Audio_time
	err = global.GVA_DB.Where("room_id = ?", token.RoomId).First(&room).Error
	if err != nil {
		response.FailWithDetailed(err.Error(), "失败", c)
		return
	}

	resData := TokenResponst{Token: room.Token, Teacher: room.Founder, Student: room.AcceptedBy}
	response.OkWithDetailed(resData, "成功", c)
}
