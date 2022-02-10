package model

import (
	"github.com/skywalkerwei/pluton-faster/common/gormV2"
)

type SysUserRoleModel struct {
	gormV2.BaseDel
	UserID    int64  `gorm:"column:user_id" json:"userId"`       // 用户ID
	RoleID    int64  `gorm:"column:role_id" json:"roleId"`       // 角色ID
	CreatedBy string `gorm:"column:created_by" json:"createdBy"` // 创建人
	UpdatedBy string `gorm:"column:updated_by" json:"updatedBy"` // 更新人
}

func (u *SysUserRoleModel) TableName() string {
	return "sys_user_role"
}
