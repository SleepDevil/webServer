package service

import (
	"fmt"
	"yasi_audio/global"
	"yasi_audio/model"

	uuid "github.com/satori/go.uuid"
)

// true为失败，false为成功
func CheckUuid(uuid string) (success bool, msg string) {
	var uuid_instance model.Uuid
	err := global.GVA_DB.Where("`uuid` = ?", uuid).First(&uuid_instance).Error
	fmt.Println(err)
	if err != nil {
		return true, "无效邀请码"
	}
	used := uuid_instance.Used
	if used {
		return true, "该验证码已被使用过"
	} else {
		return false, "该邀请码可以使用"
	}
}

func ExpireInvitationCode(uuid string) {
	var uuid_instance model.Uuid
	global.GVA_DB.Model(&uuid_instance).Where("`uuid` = ?", uuid).Update("used", true)
}

func Invite() (success bool, invitation_code uuid.UUID) {
	var uuid_instance model.Uuid
	uuid_instance.UUID = uuid.NewV4()
	uuid_instance.Used = false
	err := global.GVA_DB.Create(&uuid_instance).Error
	if err == nil {
		return true, uuid_instance.UUID
	}
	return false, uuid.Nil
}
