package api

import (
	"fmt"
	"math/rand"
	"time"
	"yasi_audio/global"
	"yasi_audio/model"
	"yasi_audio/model/response"

	"github.com/gin-gonic/gin"
)

type QuestionRequest struct {
	TopicId string `json:"topicId"`
}

type TopicRequest struct {
	Part string `json:"part"`
}

func GetQuestions(c *gin.Context) {
	resData := QuestionRequest{}
	c.ShouldBindJSON(&resData)
	fmt.Println(resData.TopicId)
	QuestionArr := make([]model.Question, 0)

	err := global.GVA_DB.Where("topic_id = ?", resData.TopicId).Find(&QuestionArr).Error
	if err != nil {
		response.FailWithDetailed(err.Error(), "失败", c)
		return
	}
	response.OkWithDetailed(QuestionArr, "成功", c)
}

func GetTopic(c *gin.Context) {
	resData := TopicRequest{}
	c.ShouldBindJSON(&resData)
	TopicArr := make([]model.Topic, 0)
	err := global.GVA_DB.Where("part = ?", resData.Part).Find(&TopicArr).Error
	rand.Seed(time.Now().Unix())
	randIndex := rand.Intn(len(TopicArr))
	// 随机返回一个topic
	fmt.Println(randIndex)
	if err != nil {
		response.FailWithDetailed(err.Error(), "失败", c)
		return
	}
	response.OkWithDetailed(TopicArr[randIndex], "成功", c)
}
