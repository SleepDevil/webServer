package model

type Room_Id struct {
	RoomId string `json:"RoomId" gorm:"comment:房间号"`
	Token  string `json:"Token" gorm:"comment:token"`
}
