package api

import (
	"context"
	"encoding/json"
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
	randIndex1 := rand.Intn(len(TopicArr))
	randIndex2 := rand.Intn(len(TopicArr))
	// 随机返回一个topic数组，长度为2
	resArr := make([]model.Topic, 0)
	resArr = append(resArr, TopicArr[randIndex1], TopicArr[randIndex2])
	fmt.Println(resArr)
	var FirstQuestion model.Topic = TopicArr[randIndex1]
	var SecondQuestion model.Topic = TopicArr[randIndex2]
	fmt.Println(FirstQuestion.TopicName)
	fmt.Println(SecondQuestion.TopicName)
	marshalData, err := json.Marshal(resArr)
	fmt.Println(err)

	_, err = global.GVA_REDIS.Set(ctx, resData.RoomId, marshalData, 60*60*24*time.Second).Result()
	if err != nil {
		response.FailWithDetailed(err.Error(), "失败", c)
		return
	}
	response.OkWithDetailed(resArr, "成功", c)
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
	// tempData := json.Unmarshal()
	response.OkWithDetailed(val, "成功", c)
}
