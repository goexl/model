package model

// Optimistic 乐观锁
type Optimistic struct {
	// 版本
	// nolint:lll
	Version int `xorm:"version notnull default(0) comment(版本，用于实现乐观锁，当更新时本字段自动增一)" json:"version,omitempty"`
}
