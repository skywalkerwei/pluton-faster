package model

import (
	"github.com/skywalkerwei/pluton-faster/common/gormV2"
)

type SysRoleDeptModel struct {
	gormV2.BaseDel
	RoleID    int64  `gorm:"column:role_id" json:"roleId"`       // 角色ID
	DeptID    int64  `gorm:"column:dept_id" json:"deptId"`       // 机构ID
	CreatedBy string `gorm:"column:created_by" json:"createdBy"` // 创建人
	UpdatedBy string `gorm:"column:updated_by" json:"updatedBy"` // 更新人
}

func (u *SysRoleDeptModel) TableName() string {
	return "sys_role_dept"
}
