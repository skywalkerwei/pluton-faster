package model

import (
	"github.com/skywalkerwei/pluton-faster/common/gormV2"
)

type SysMenuModel struct {
	gormV2.BaseDel
	Name          string `gorm:"column:name" json:"name"`          // 菜单名称
	ParentID      int64  `gorm:"column:parent_id" json:"parentId"` // 父菜单ID，一级菜单为0
	URL           string `gorm:"column:url" json:"url"`
	Perms         string `gorm:"column:perms" json:"perms"`                  // 授权(多个用逗号分隔，如：sys:user:add,sys:user:edit)
	Tp            int    `gorm:"column:tp" json:"tp"`                        // 类型   0：目录   1：菜单   2：按钮
	Icon          string `gorm:"column:icon" json:"icon"`                    // 菜单图标
	OrderNum      int    `gorm:"column:order_num" json:"orderNum"`           // 排序
	VuePath       string `gorm:"column:vue_path" json:"vuePath"`             // vue系统的path
	VueComponent  string `gorm:"column:vue_component" json:"vueComponent"`   // vue的页面
	VueIcon       string `gorm:"column:vue_icon" json:"vueIcon"`             // vue的图标
	VueRedirect   string `gorm:"column:vue_redirect" json:"vueRedirect"`     // vue的路由重定向
	BackgroundURL string `gorm:"column:background_url" json:"backgroundUrl"` // 后台地址
	CreatedBy     string `gorm:"column:created_by" json:"createdBy"`         // 创建人
	UpdateBy      string `gorm:"column:update_by" json:"updateBy"`           // 更新人
}

func (u *SysMenuModel) TableName() string {
	return "sys_menu"
}
