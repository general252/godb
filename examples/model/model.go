package model

import (
	"github.com/general252/godb/godb"
	"log"
	"time"
)

type Model struct {
	ID        uint
	Uid       *string `gorm:"column:uid;type:string;uniqueIndex;not null"` // 唯一索引
	CreatedAt *time.Time
	UpdatedAt *time.Time
}

type Book struct {
	Model
	Author *string `gorm:"column:author;type:string;default:'';size:64"`                    // 作者
	Name   *string `gorm:"column:name;type:string;default:'';size:128"`                     // 书名
}

type User struct {
	Model
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
	if err := godb.BuildDatabaseTable(GetModelBeans()); err != nil {
		log.Println(err)
	}
}

func init() {
	Build()
}
