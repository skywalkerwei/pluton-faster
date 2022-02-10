package gormV2

import (
	"github.com/skywalkerwei/pluton-faster/common/gormV2/types"
	"gorm.io/plugin/soft_delete"
)

type Base struct {
	ID        int64      `gorm:"primarykey" json:"id"`
	CreatedAt types.Time `gorm:"created_at" json:"createdAt"`
	UpdatedAt types.Time `gorm:"updated_at" json:"updatedAt"`
}

type BaseDel struct {
	ID        int64                 `gorm:"primarykey" json:"id"`
	CreatedAt types.Time            `gorm:"created_at" json:"createdAt"`
	UpdatedAt types.Time            `gorm:"updated_at" json:"updatedAt"`
	DeletedAt soft_delete.DeletedAt `gorm:"delete_at" json:"-"`
}
