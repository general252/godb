package godb

import "time"

type Model struct {
	ID        uint
	Uid       *string `gorm:"column:uid;type:string;uniqueIndex;size:64;not null"` // 唯一索引
	CreatedAt *time.Time
	UpdatedAt *time.Time
}

func (tis *Model) GetID() uint {
	return tis.ID
}

type InterfaceModel interface {
	GetID() uint
}
