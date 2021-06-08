package api

import (
	"fmt"
	"time"
	"yasi_audio/model"
	"yasi_audio/model/request"
	"yasi_audio/model/response"
	"yasi_audio/service"
	"yasi_audio/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateTask(c *gin.Context) {
	var T request.Task
	_ = c.ShouldBindJSON(&T)
	fmt.Println(T)
	t, _ := time.ParseInLocation("2006-01-02 15:04:05", T.BattleTime, time.Local)
	if err := utils.Verify(T, utils.CreateTaskVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	task := &model.Task{CreatedBy: T.CreatedBy, CreatedHead: T.CreatedHead, CreatedScore: T.CreatedScore, AcceptedBy: T.AcceptedBy, Accepted: T.Accepted, TargetScore: T.TargetScore, BattleTime: t, RoomId: T.RoomId}
	err, taskReturn := service.CreateTask(*task)
	fmt.Println(*task)
	if err != nil {
		response.FailWithDetailed(response.TaskResponse{Task: taskReturn}, ""+err.Error(), c)
	} else {
		response.OkWithDetailed(response.TaskResponse{Task: *task}, "创建成功", c)
	}
}

func GetTasks(c *gin.Context) {
	err, taskLists, total := service.GetTask()
	fmt.Println(taskLists)
	if err != nil {
		response.FailWithMessage("失败", c)
	} else {
		response.OkWithDetailed(response.GetTasksResponse{Tasks: taskLists, Total: total}, "成功", c)
	}
}

type AcceptRoom struct {
	RoomId     string
	AcceptedBy string
}

func AcceptTask(c *gin.Context) {
	accept := AcceptRoom{}
	c.BindJSON(&accept)
	err, _ := service.AcceptTask(accept.RoomId, accept.AcceptedBy)
	if err == gorm.ErrRecordNotFound {
		response.FailWithMessage("未找到该房间", c)
	}
	if err != nil {
		response.FailWithMessage("未知错误，请联系管理员", c)
	} else {
		response.OkWithMessage("接受成功！", c)
	}
}
