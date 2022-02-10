package model

import (
	"github.com/skywalkerwei/pluton-faster/common/gormV2"
)

type SysJobModel struct {
	gormV2.BaseDel
	UserName  string `gorm:"column:user_name" json:"userName"`   // 用户名
	Operation string `gorm:"column:operation" json:"operation"`  // 用户操作
	Method    string `gorm:"column:method" json:"method"`        // 请求方法
	Params    string `gorm:"column:params" json:"params"`        // 请求参数
	Time      int64  `gorm:"column:time" json:"time"`            // 执行时长(毫秒)
	IP        string `gorm:"column:ip" json:"ip"`                // IP地址
	CreatedBy string `gorm:"column:created_by" json:"createdBy"` // 创建人
	UpdatedBy string `gorm:"column:updated_by" json:"updatedBy"` // 更新人
}

func (u *SysJobModel) TableName() string {
	return "sys_job"
}
