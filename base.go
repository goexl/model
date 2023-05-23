package model

import (
	"time"
)

// Base 基础数据库模型
type Base struct {
	// 编号
	Id int64 `xorm:"pk notnull unique index('idx_id') default(0)" json:"id,string,omitempty"`
	// 创建时间
	Created time.Time `xorm:"created notnull default('2022-12-23 09:55:52')" json:"created,omitempty"`
	// 最后更新时间
	Updated time.Time `xorm:"updated notnull default('2022-12-23 09:55:52')" json:"updated,omitempty"`
}
