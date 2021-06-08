package model

import (
	"time"
	"yasi_audio/global"
)

type Task struct {
	global.GVA_MODEL
	CreatedBy    string    `json:"CreatedBy" gorm:"comment:创建人账号"`                                                                  //创建人账号，对应username字段
	CreatedHead  string    `json:"headerImg" gorm:"default:https://seopic.699pic.com/photo/40011/0709.jpg_wh1200.jpg;comment:用户头像"` // 创建人头像
	CreatedScore string    `json:"OralScore" gorm:"comment:口语成绩"`                                                                   //创建人口语成绩
	AcceptedBy   string    `json:"AcceptedBy" gorm:"comment:接受人账号"`                                                                 // 空字符串为未接受
	Accepted     bool      `json:"accepted" gorm:"comment:是否接受"`                                                                    // true为接受0为未接受
	TargetScore  string    `json:"TargetScore" gorm:"comment:目标分数"`                                                                 // 目标分数
	BattleTime   time.Time `json:"BattleTime" gorm:"comment:约定时间"`                                                                  //约定时间
	RoomId       string    `json:"RoomId" gorm:"comment:房间号;uniqueIndex"`                                                           //约定房间号
}
