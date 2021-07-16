package model

import (
	uuid "github.com/satori/go.uuid"
)

type SysUser struct {
	UUID      uuid.UUID `json:"uuid" gorm:"comment:用户UUID"`                                                                      // 用户UUID
	Username  string    `json:"userName" gorm:"comment:用户登录名"`                                                                   // 用户登录名
	Password  string    `json:"-"  gorm:"comment:用户登录密码"`                                                                        // 用户登录密码
	OralScore string    `json:"oralscore" gorm:"comment:口语成绩"`                                                                   // 用户口语成绩
	NickName  string    `json:"nickName" gorm:"default:系统用户;comment:用户昵称"`                                                       // 用户昵称
	HeaderImg string    `json:"headerImg" gorm:"default:https://seopic.699pic.com/photo/40011/0709.jpg_wh1200.jpg;comment:用户头像"` // 用户头像
	Identity  int       `json:"identity" gorm:"comment:1为学生，2为已工作"`                                                              // 1为学生，2为已工作
}
