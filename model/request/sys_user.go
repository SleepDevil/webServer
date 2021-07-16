package request

type Register struct {
	Username        string `json:"userName"`
	Password        string `json:"passWord"`
	NickName        string `json:"nickName" gorm:"default:'QMPlusUser'"`
	HeaderImg       string `json:"headerImg" gorm:"default:'https://seopic.699pic.com/photo/40011/0709.jpg_wh1200.jpg'"`
	Invitation_Code string `json:"invitation_code"`
	OralScore       string `json:"oralscore"`
	Identity        int    `json:"identity" gorm:"comment:1为学生，2为已工作"` // 1为学生，2为已工作
}

// User login structure
type Login struct {
	Username string `json:"username"` // 用户名
	Password string `json:"password"` // 密码
}
