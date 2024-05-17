package model

import (
	"time"
)

// Update 带修改时间模型
type Update struct {
	// 最后更新时间
	// nolint:lll
	Updated time.Time `xorm:"updated notnull default('2022-12-23 09:55:52')  comment('最后更新时间，每次更新时自动修改')" json:"updated,omitempty"`
}
