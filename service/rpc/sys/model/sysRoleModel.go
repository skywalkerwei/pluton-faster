package model

import (
	"github.com/skywalkerwei/pluton-faster/common/gormV2"
)

type SysRoleModel struct {
	gormV2.BaseDel
	Name      string `gorm:"column:name" json:"name"`            // 角色名称
	Remark    string `gorm:"column:remark" json:"remark"`        // 备注
	Status    int8   `gorm:"column:status" json:"status"`        // 状态  1:启用,0:禁用
	CreatedBy string `gorm:"column:created_by" json:"createdBy"` // 创建人
	UpdatedBy string `gorm:"column:updated_by" json:"updatedBy"` // 更新人
}

func (u *SysRoleModel) TableName() string {
	return "sys_role"
}
