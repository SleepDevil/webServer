package request

type Audio_time struct {
	Founder    string `json:"founder" gorm:"comment:语音发起人"`
	AcceptedBy string `json:"acceptedby" gorm:"comment:语音接受人"`
	RoomId     string `json:"roomid" gorm:"comment:语音房间号;uniqueIndex"`
	BattleTime string `json:"battletime" gorm:"comment:约定时间"`
}
