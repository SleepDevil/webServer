package model

import (
	uuid "github.com/satori/go.uuid"
)

type Uuid struct {
	UUID uuid.UUID `json:"uuid" gorm:"comment:用户UUID"` // 用户UUID
	Used bool      `json:"used" gorm:"comment:"使用过"`
}
