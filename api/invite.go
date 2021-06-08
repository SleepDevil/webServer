package api

import (
	"fmt"
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
	fmt.Println(U.Invitation_Code)
	err := c.PostForm("Invitation_Code")
	fmt.Println(err)
	success, msg := service.CheckUuid(U.Invitation_Code)
	if success {
		c.JSON(200, gin.H{
			"msg":     msg,
			"success": "false",
		})
	} else {
		c.JSON(200, gin.H{
			"msg":     msg,
			"success": "true",
		})
	}
}
