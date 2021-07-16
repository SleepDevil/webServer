package api

import (
	"yasi_audio/global"
	"yasi_audio/middleware"
	"yasi_audio/model"
	"yasi_audio/model/response"

	"github.com/gin-gonic/gin"
)

func CheckJwt(c *gin.Context) {
	var user model.SysUser
	json := make(map[string]interface{})
	c.BindJSON(&json)
	userClaim, err := middleware.NewJWT().ParseToken(json["token"].(string))
	error := global.GVA_DB.Where("username=?", userClaim.Username).First(&user).Error
	if error != nil {
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithDetailed(user, "成功", c)
	}
}
