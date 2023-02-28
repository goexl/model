package model

import (
	"time"
)

// Delete 软删除模型
type Delete struct {
	// 删除时间
	// 用于实现软删除
	Deleted time.Time `xorm:"deleted notnull default('2022-12-23 09:55:52')" json:"deleted"`
}
