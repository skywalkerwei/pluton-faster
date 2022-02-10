package model

import (
	"github.com/skywalkerwei/pluton-faster/common/gormV2"
)

type SysRoleMenuModel struct {
	gormV2.Base
	RoleID    int64  `gorm:"column:role_id" json:"roleId"`       // 角色ID
	MenuID    int64  `gorm:"column:menu_id" json:"menuId"`       // 菜单ID
	CreatedBy string `gorm:"column:created_by" json:"createdBy"` // 创建人
	UpdatedBy string `gorm:"column:updated_by" json:"updatedBy"` // 更新人
}

func (u *SysRoleMenuModel) TableName() string {
	return "sys_role_menu"
}
