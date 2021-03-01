package godb

import "time"

type Model struct {
	ID        uint
	Uid       *string `gorm:"column:uid;type:string;uniqueIndex;not null"` // 唯一索引
	CreatedAt *time.Time
	UpdatedAt *time.Time
}
