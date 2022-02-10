package model

import (
	"github.com/skywalkerwei/pluton-faster/common/gormV2"
)

type SysConfigModel struct {
	gormV2.BaseDel
	Value       string `gorm:"column:value" json:"value"`             // 数据值
	Label       string `gorm:"column:label" json:"label"`             // 标签名
	Tp          string `gorm:"column:tp" json:"tp"`                   // 类型
	Description string `gorm:"column:description" json:"description"` // 描述
	Sort        int    `gorm:"column:sort" json:"sort"`               // 排序（升序）
	Remarks     string `gorm:"column:remarks" json:"remarks"`         // 备注信息
	CreatedBy   string `gorm:"column:created_by" json:"createdBy"`    // 创建人
	UpdatedBy   string `gorm:"column:updated_by" json:"updatedBy"`    // 更新人
}

func (u *SysConfigModel) TableName() string {
	return "sys_config"
}
