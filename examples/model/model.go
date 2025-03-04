package model

import (
	"log"
	"time"

	"github.com/general252/godb/godb"
)

type Book struct {
	godb.Model
	Author *string `gorm:"column:author;type:string;default:'';size:64" comment:"作者"` // 作者2
	Name   *string `gorm:"column:name;type:string;default:'';size:128" comment:"书名"`  // 书名
}

type User struct {
	godb.Model
	Name      *string    // b
	Age       *int       // c
	Birthday  *time.Time // d
	CompanyID *uint      // e
	ManagerID *uint      `gorm:"column:manager_id;type:int;default:0"` // f
	A         *float64   // a
	B         *float32   // b
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
