package model

import "gorm.io/gorm"

// Notice 公告
type Notice struct {
	gorm.Model
	Text string `gorm:"type:text"`
}

func (user *Notice) TableName() string {
	return "Notice"
}
