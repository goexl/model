package model

import (
	"strconv"
	"time"
)

// Base 基础数据库模型
type Base struct {
	// 编号
	Id uint64 `xorm:"pk notnull default(0) comment(编号，用来唯一标识数据)" json:"id,string,omitempty"`
	// 创建时间
	// nolint:lll
	Created time.Time `xorm:"created notnull default(CURRENT_TIMESTAMP) comment(创建时间，创建时自动设置)" json:"created,omitempty"`
	// 最后更新时间
	// nolint:lll
	Updated time.Time `xorm:"updated notnull default(CURRENT_TIMESTAMP) comment(最后更新时间，每次更新时自动修改)" json:"updated,omitempty"`
}

func (b *Base) Identify() string {
	return strconv.FormatUint(b.Id, 10)
}
