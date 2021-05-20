package api

import (
	"yasi_audio/model/request"
	"yasi_audio/service"

	"github.com/gin-gonic/gin"
)

func Invite(c *gin.Context) {
	// Creating UUID Version 4
	// u1 := uuid.NewV4()
	// invitation := &model.Uuid{u1, false}
	err, uuid := service.Invite()
	c.JSON(200, gin.H{
		"success":         err,
		"invitation_code": uuid,
	})
}

func Check_Invited(c *gin.Context) {
	var U request.Uuid
	_ = c.ShouldBindJSON(&U)
	if service.CheckUuid(U.Uuid) {
		c.JSON(200, gin.H{
			"msg":     "该邀请码已经使用过",
			"success": "false",
		})
	} else {
		c.JSON(200, gin.H{
			"msg":     "该邀请码可以使用",
			"success": "true",
		})
	}
}
