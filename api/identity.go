package api

import (
	"yasi_audio/global"
	"yasi_audio/model"
	"yasi_audio/model/response"

	"github.com/gin-gonic/gin"
)

type Identity struct {
	Username string `json:"username"`
}

func GetIdentity(c *gin.Context) {
	json := Identity{}
	c.BindJSON(&json)
	toFind := json.Username
	user := map[string]interface{}{}
	err := global.GVA_DB.Debug().Model(&model.SysUser{}).Where("nick_name = ?", toFind).First(&user).Error
	if err != nil {
		response.FailWithMessage("失败", c)
		return
	}
	if user["identity"] == 1 {
		response.OkWithDetailed("study", "成功", c)
		return
	}
	if user["identity"] == 2 {
		response.OkWithDetailed("work", "成功", c)
		return
	}
	response.FailWithMessage("失败", c)
}
