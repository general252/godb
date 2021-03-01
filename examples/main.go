package main

import (
	"github.com/general252/godb/examples/model"
	"github.com/general252/godb/godb"
	"gorm.io/gorm"
	"log"
	"time"
)

func main() {
	_ = model.User{}


	var err error
	err = model.DefaultEngine.Init(&model.InitParam{
		Host:     "192.168.6.59",
		Port:     3306,
		DBName:   "test_go_db",
		Username: "root",
		Password: "123456",
		ShowSQL:  false,
	})
	if err != nil {
		log.Println(err)
		return
	}

	var helpUser = model.NewBeanUser()
	_, _ = helpUser.Add(&model.User{
		Model: godb.Model{
			Uid: model.String("uid2"),
		},
		Name:      model.String("name"),
		Age:       model.Int(12),
		Birthday:  model.Time(time.Now()),
		CompanyID: model.Uint(1234),
	})

	_ = helpUser.UpdateByUId(&model.User{
		Model: godb.Model{
			Uid: model.String("uid1"),
		},
		Age: model.Int(200),
	})

	_, _, _ = helpUser.Filter(&model.User{
		Model: godb.Model{
			Uid: model.String("uid"),
		},
	}, 1, 0, func(r *gorm.DB, field *model.GoUser) *gorm.DB {
		r = r.Where("%v=%v", field.FieldNameAge(), 123)
		return r
	})
}
