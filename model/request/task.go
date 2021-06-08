package request

type Task struct {
	CreatedBy    string `json:"createdby"`    //创建人账号，对应username字段
	CreatedHead  string `json:"createdhead"`  // 创建人头像
	CreatedScore string `json:"createdscore"` //创建人口语成绩
	AcceptedBy   string `json:"acceptedby"`
	Accepted     bool   `json:"accepted"`    // true为接受0为未接受
	TargetScore  string `json:"targetscore"` // 目标分数
	BattleTime   string `json:"battletime"`  //约定时间
	RoomId       string `json:"roomid"`      //约定房间号
}
