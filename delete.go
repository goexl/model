package model

import (
	"time"
)

// Delete 软删除模型
type Delete struct {
	// 删除时间
	// 用于实现软删除
	// nolint:lll
	Deleted time.Time `xorm:"deleted notnull default('2022-12-23 09:55:52') comment('删除时间，通过本字段实现实现软删除功能，使用删除功能时自动设置本时间')" json:"deleted,omitempty"`
}
