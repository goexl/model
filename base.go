package model

import (
	"strconv"
	"time"
)

// Base 基础数据库模型
type Base struct {
	// 编号
	Id int64 `xorm:"pk notnull default(0)" json:"id,string,omitempty"`
	// 创建时间
	Created time.Time `xorm:"created notnull default('2022-12-23 09:55:52')" json:"created,omitempty"`
	// 最后更新时间
	Updated time.Time `xorm:"updated notnull default('2022-12-23 09:55:52')" json:"updated,omitempty"`
}

func (b *Base) Identify() string {
	return strconv.FormatInt(b.Id, 10)
}
