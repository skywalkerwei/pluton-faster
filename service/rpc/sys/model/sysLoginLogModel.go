package model

import (
	"github.com/skywalkerwei/pluton-faster/common/gormV2"
)

type SysLoginLogModel struct {
	gormV2.Base
	UserName string `gorm:"column:user_name" json:"userName"` // 用户名
	Status   string `gorm:"column:status" json:"status"`      // 登录状态（online:在线，登录初始状态，方便统计在线人数；login:退出登录后将online置为login；logout:退出登录）
	IP       string `gorm:"column:ip" json:"ip"`              // IP地址
}

func (u *SysLoginLogModel) TableName() string {
	return "sys_login_log"
}
