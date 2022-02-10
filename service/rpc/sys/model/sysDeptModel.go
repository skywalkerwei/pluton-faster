package model

import (
	"github.com/skywalkerwei/pluton-faster/common/gormV2"
)

type SysDeptModel struct {
	gormV2.BaseDel
	Name      string `gorm:"column:name" json:"name"`            // 机构名称
	ParentID  int64  `gorm:"column:parent_id" json:"parentId"`   // 上级机构ID，一级机构为0
	OrderNum  int    `gorm:"column:order_num" json:"orderNum"`   // 排序
	CreatedBy string `gorm:"column:created_by" json:"createdBy"` // 创建人
	UpdatedBy string `gorm:"column:updated_by" json:"updatedBy"` // 更新人
}

func (u *SysDeptModel) TableName() string {
	return "sys_dept"
}
