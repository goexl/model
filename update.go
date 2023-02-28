package model

import (
	"time"
)

// Update 带修改时间模型
type Update struct {
	// 最后更新时间
	Updated time.Time `xorm:"updated notnull default('2022-12-23 09:55:52')" json:"updated"`
}
