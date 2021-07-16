package api

import (
	"fmt"
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
	global.GVA_DB.Model(&model.SysUser{}).Where("username = ?", toFind).First(&user)
	fmt.Println(user)
	if user["identity"] == 1 {
		response.OkWithDetailed("学生", "成功", c)
		return
	}
	if user["identity"] == 2 {
		response.OkWithDetailed("已工作", "成功", c)
		return
	}

}
