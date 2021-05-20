package service

import (
	"fmt"
	"yasi_audio/global"
	"yasi_audio/model"

	uuid "github.com/satori/go.uuid"
)

// return的值为数据库中used的值，false为未使用过
func CheckUuid(uuid string) (success bool) {
	var uuid_instance model.Uuid
	global.GVA_DB.Where("`uuid` = ?", uuid).First(&uuid_instance)
	used := uuid_instance.Used
	fmt.Println("used", used)
	global.GVA_DB.Model(&uuid_instance).Where("`uuid` = ?", uuid).Update("used", true)
	fmt.Println("instanceUsed", uuid_instance.Used)
	return used
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
