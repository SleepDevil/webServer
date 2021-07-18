package api

import (
	"context"
	"fmt"
	"math/rand"
	"time"
	"yasi_audio/global"
	"yasi_audio/model"
	"yasi_audio/model/response"

	"github.com/gin-gonic/gin"
)

var ctx = context.Background()

type QuestionRequest struct {
	TopicId string `json:"topicId"`
}

type TopicRequest struct {
	Part   string `json:"part"`
	RoomId string `json:"roomId"`
}

type Part2Request struct {
	RoomId string `json:"roomId"`
}

func GetQuestions(c *gin.Context) {
	resData := QuestionRequest{}
	c.ShouldBindJSON(&resData)
	QuestionArr := make([]model.Question, 0)

	err := global.GVA_DB.Where("topic_id = ?", resData.TopicId).Find(&QuestionArr).Error
	rand.Seed(time.Now().Unix())
	QuestionArrLen := len(QuestionArr)
	randIndex1 := rand.Intn(QuestionArrLen)
	randIndex2 := rand.Intn(QuestionArrLen)
	randIndex3 := rand.Intn(QuestionArrLen)
	if err != nil {
		response.FailWithDetailed(err.Error(), "失败", c)
		return
	}
	resArr := make([]model.Question, 0)
	appendErr := append(resArr, QuestionArr[randIndex1], QuestionArr[randIndex2], QuestionArr[randIndex3])
	fmt.Println(appendErr)
	response.OkWithDetailed(appendErr, "成功", c)
}

func GetTopic(c *gin.Context) {
	resData := TopicRequest{}
	c.ShouldBindJSON(&resData)
	TopicArr := make([]model.Topic, 0)
	err := global.GVA_DB.Where("part = ?", resData.Part).Find(&TopicArr).Error
	if err != nil {
		response.FailWithDetailed(err.Error(), "失败", c)
		return
	}
	rand.Seed(time.Now().Unix())
	randIndex := rand.Intn(len(TopicArr))
	// 随机返回一个topic
	var randomData model.Topic = TopicArr[randIndex]

	_, err = global.GVA_REDIS.Set(ctx, resData.RoomId, randomData.TopicName, time.Duration(30*time.Second)).Result()
	if err != nil {
		response.FailWithDetailed(err.Error(), "失败", c)
		return
	}
	response.OkWithDetailed(TopicArr[randIndex], "成功", c)
}

func GetPart2Question(c *gin.Context) {
	resData := Part2Request{}
	c.ShouldBindJSON(&resData)
	fmt.Println(resData.RoomId)
	val, err := global.GVA_REDIS.Get(ctx, resData.RoomId).Result()
	if err != nil {
		fmt.Println(err.Error())
		response.FailWithDetailed(err.Error(), "获取失败", c)
		return
	}
	fmt.Println(val)
	response.OkWithDetailed(val, "成功", c)
}
