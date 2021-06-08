package service

import (
	"fmt"
	"yasi_audio/global"
	"yasi_audio/model"
)

func CreateTask(t model.Task) (err error, taskReturn model.Task) {
	var task model.Task
	err = global.GVA_DB.Create(&t).Error
	return err, task
}

func GetTask() (err error, tasksList interface{}, total int64) {
	db := global.GVA_DB.Model(&model.Task{})
	var taskList []model.Task
	result := db.Find(&taskList)
	fmt.Println(result)
	return result.Error, taskList, result.RowsAffected
}

func AcceptTask(room_id string, accepted_by string) (err error, taskReturn interface{}) {
	task := model.Task{}
	err = global.GVA_DB.Model(&model.Task{}).Where("room_id = ?", room_id).First(&task).Error
	global.GVA_DB.Model(&model.Task{}).Where("room_id = ?", room_id).Update("accepted", true).Update("accepted_by", accepted_by)
	fmt.Println(task)
	return err, task
}
