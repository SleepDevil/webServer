package request

type Register struct {
	Username        string `json:"userName"`
	Password        string `json:"passWord"`
	NickName        string `json:"nickName" gorm:"default:'QMPlusUser'"`
	HeaderImg       string `json:"headerImg" gorm:"default:'http://www.henrongyi.top/avatar/lufu.jpg'"`
	Invitation_Code string `json:"invitation_code"`
}

// User login structure
type Login struct {
	Username string `json:"username"` // 用户名
	Password string `json:"password"` // 密码
}
