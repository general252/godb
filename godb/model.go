package godb

import "time"

type Model struct {
	ID        uint
	Uid       *string `gorm:"column:uid;type:string;uniqueIndex;size:64"` // 唯一索引
	CreatedAt *time.Time
	UpdatedAt *time.Time
}

func (tis *Model) GetID() uint {
	return tis.ID
}

func (tis *Model) GetUID() string {
	if tis == nil || tis.Uid == nil {
		return ""
	}

	return *tis.Uid
}

func (tis *Model) GetCreatedTime() time.Time {
	if tis == nil || tis.Uid == nil {
		return time.Unix(0, 0)
	}

	return *tis.CreatedAt
}

func (tis *Model) GetUpdatedTime() time.Time {
	if tis == nil || tis.Uid == nil {
		return time.Unix(0, 0)
	}

	return *tis.UpdatedAt
}

type InterfaceModel interface {
	GetID() uint
	GetUID() string
	GetCreatedTime() time.Time
	GetUpdatedTime() time.Time
}
