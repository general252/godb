package model

import (
	"github.com/general252/godb/tool"
	"log"
	"time"
)

type Book struct {
	ID     uint    // key id
	Uid    *string `gorm:"column:uid;type:string;default:'';uniqueIndex:udx_book;not null"` // 唯一索引
	Author *string `gorm:"column:author;type:string;default:'';size:64"`                    // 作者
	Name   *string `gorm:"column:name;type:string;default:'';size:128"`                     // 书名
}

type User struct {
	ID        uint       // key id
	Uid       *string    `gorm:"column:uid;type:string;default:'';uniqueIndex:udx_user;not null"` // a
	Name      *string    // b
	Age       *int       // c
	Birthday  *time.Time // d
	CompanyID *uint      // e
	ManagerID *uint      `gorm:"column:manager_id;type:int;default:0"` // f
}

// GetModelBeans 获取需要初始化的模型
func GetModelBeans() []interface{} {
	var beans = []interface{}{
		new(Book),
		new(User),
	}
	return beans
}

func Build() {
	if err := tool.BuildDatabaseTable(GetModelBeans()); err != nil {
		log.Println(err)
	}
}
