package model

import (
	"time"
)

// Update 带修改时间模型
type Update struct {
	// 编号
	Id int64 `xorm:"pk notnull unique index('idx_id') default(0)" json:"id,string"`
	// 最后更新时间
	Updated time.Time `xorm:"updated default('2022-12-23 09:55:52')" json:"updated"`
}
