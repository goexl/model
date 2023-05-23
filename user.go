package model

import (
	"time"
)

// User 基础用户数据
type User struct {
	// 编号
	Id int64 `xorm:"pk notnull unique index('idx_id') default(0)" json:"id,string,omitempty"`
	// 创建时间
	Created time.Time `xorm:"created notnull default('2022-12-23 09:55:52')" json:"created,omitempty"`
	// 最后更新时间
	Updated time.Time `xorm:"updated notnull default('2022-12-23 09:55:52')" json:"updated,omitempty"`

	// 用户名
	// nolint: lll
	Username string `xorm:"varchar(32) notnull default('') unique(uidx_name)" json:"username,omitempty" validate:"omitempty,min=1,max=32,email"`
	// 手机号
	// 类似于+86-17089792784
	// nolint: lll
	Phone string `xorm:"varchar(15) notnull default('') unique(uidx_phone)" json:"phone,omitempty" validate:"omitempty,mobile"`
	// 密码
	Passwd string `xorm:"varchar(512) notnull default('')" json:"-"`
}
