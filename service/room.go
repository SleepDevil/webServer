package service

import (
	"yasi_audio/global"
	"yasi_audio/model"
)

func CreateRoom(r model.Audio_time) (err error, roomReturn model.Audio_time) {
	var room model.Audio_time
	err = global.GVA_DB.Create(&r).Error
	return err, room
}
