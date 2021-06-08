package api

import (
	"fmt"
	"yasi_audio/global"
	"yasi_audio/model"
	"yasi_audio/model/response"

	"github.com/gin-gonic/gin"
)

var filename string

func CreateAvator(c *gin.Context) {

	file, _ := c.FormFile("file")
	filename = file.Filename
	c.SaveUploadedFile(file, "/home/useravators/"+filename)
	response.OkWithMessage(file.Filename, c)
}

func UpdaeUserHeaderImg(c *gin.Context) {
	Username := c.Param("username")
	Nickname := c.Param("nickname")
	fmt.Println(Nickname)
	user := map[string]interface{}{}
	global.GVA_DB.Model(&model.SysUser{}).Where("username = ?", Username).First(&user)
	global.GVA_DB.Model(&model.SysUser{}).Where("username = ?", Username).Update("header_img", "https://sleepdevil.top/imgs/"+filename)
	global.GVA_DB.Model(&model.Task{}).Where("created_by = ?", Nickname).Update("created_head", "https://sleepdevil.top/imgs/"+filename)
	response.OkWithMessage("成功", c)
}
