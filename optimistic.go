package model

// Optimistic 乐观锁
type Optimistic struct {
	// 版本
	Version int `xorm:"version notnull default(0)" json:"version,omitempty"`
}
