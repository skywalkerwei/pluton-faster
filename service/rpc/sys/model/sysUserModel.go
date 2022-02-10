package model

import (
	"github.com/skywalkerwei/pluton-faster/common/gormV2"
)

type SysUserModel struct {
	gormV2.BaseDel
	Name      string `gorm:"column:name" json:"name"`            // 用户名
	NickName  string `gorm:"column:nick_name" json:"nickName"`   // 昵称
	Avatar    string `gorm:"column:avatar" json:"avatar"`        // 头像
	Password  string `gorm:"column:password" json:"password"`    // 密码
	Salt      string `gorm:"column:salt" json:"salt"`            // 加密盐
	Email     string `gorm:"column:email" json:"email"`          // 邮箱
	Mobile    string `gorm:"column:mobile" json:"mobile"`        // 手机号
	DeptID    int64  `gorm:"column:dept_id" json:"deptId"`       // 机构ID
	JobID     int    `gorm:"column:job_id" json:"jobId"`         // 岗位Id
	Status    int8   `gorm:"column:status" json:"status"`        // 状态  0：禁用   1：正常
	CreatedBy string `gorm:"column:created_by" json:"createdBy"` // 创建人
	UpdatedBy string `gorm:"column:updated_by" json:"updatedBy"` // 更新人
}

func (u *SysUserModel) TableName() string {
	return "sys_user"
}
