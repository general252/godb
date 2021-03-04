package main

import (
	"encoding/json"
	"github.com/general252/godb/examples/model"
	"github.com/general252/godb/godb"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"log"
	"time"
)

func main() {
	_ = model.User{}
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	var err error

	// 初始化
	err = model.DefaultEngine.Init(&model.InitParam{
		Host:     "192.168.6.59",
		Port:     9306,
		DBName:   "test_go_db",
		Username: "root",
		Password: "123456",
		ShowSQL:  true,
	})
	if err != nil {
		log.Println(err)
		return
	}

	var helpUser = model.NewBeanUser()
	// 添加
	_, _ = helpUser.Add(&model.User{
		Model: godb.Model{
			Uid: godb.String("uid2"),
		},
		Name:      godb.String("name"),
		Age:       godb.Int(12),
		Birthday:  godb.Time(time.Now()),
		CompanyID: godb.Uint(1234),
	})

	// 修改
	_ = helpUser.UpdateByUId(&model.User{
		Model: godb.Model{
			Uid: godb.String("uid1"),
		},
		Age: godb.Int(200),
	})

	// 查找
	objs, totalCount, err := helpUser.Filter(&model.UserFilter{
		Limit:  10,
		Offset: 0,
		Object: &model.User{
			Age: godb.Int(200),
		},
		Match: &model.User{
			Model: godb.Model{
				Uid: godb.String("uid1" + "%"),
			},
		},
	}, func(r *gorm.DB, field *model.GoUser) *gorm.DB {
		r.Clauses(clause.Like{
			Column: field.FieldNameName(),
			Value:  "na%",
		})
		return r
	})

	log.Println("===============================")
	log.Println(totalCount)
	log.Println(err)
	for _, obj := range objs {
		data, _ := json.MarshalIndent(obj, "", "  ")
		log.Println(string(data))
	}
}
