package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/general252/godb/examples/model"
	"github.com/general252/godb/godb"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func main() {
	_ = model.User{}
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	var err error

	// 初始化
	err = model.DefaultEngine.Init(&model.InitParam{
		Host:     "192.168.88.80",
		Port:     33306,
		DBName:   "test_go_db",
		Username: "root",
		Password: "123456",
		ShowSQL:  true,
	})
	if err != nil {
		log.Println(err)
		return
	}

	var hBook = model.NewBeanBook()
	_, _ = hBook.Add(&model.Book{
		Model: godb.Model{
			Uid: godb.String("bookUid_1"),
		},
		Author: godb.String("tony"),
		Name:   godb.String("XiYouJi"),
	})

	var helpUser = model.NewBeanUser()
	// 添加
	_, _ = helpUser.Add(&model.User{
		Model: godb.Model{
			Uid: godb.String("userUid_2"),
		},
		Name:      godb.String("tony"),
		Age:       godb.Int(12),
		Birthday:  godb.Time(time.Now()),
		CompanyID: godb.Uint(1234),
	})

	// 修改
	_ = helpUser.UpdateByUId(&model.User{
		Model: godb.Model{
			Uid: godb.String("userUid_1"),
		},
		Age: godb.Int(200),
	})

	helpUser.Filter(&model.UserFilter{
		Limit:  10,
		Offset: 0,
		Cond: &model.UserFilterCond{
			Type:      model.CondTypeLike,
			Number:    nil,
			Between:   nil,
			Container: nil,
			Like: &model.UserLike{
				Name: godb.String("KKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKK"),
			},
		},
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
				Uid: godb.String("userUid"),
			},
		},
	}, func(r *gorm.DB, field *model.GoUser) (*gorm.DB, error) {
		r.Clauses(clause.Like{
			Column: field.FieldNameName(),
			Value:  "na%",
		})
		return r, nil
	})

	log.Println("===============================")
	log.Println(totalCount)
	log.Println(err)
	for _, obj := range objs {
		data, _ := json.MarshalIndent(obj, "", "  ")
		log.Println(string(data))
	}

	{
		objectList, _, err := hBook.Filter(nil, func(r *gorm.DB, field *model.GoBook) (*gorm.DB, error) {
			subQuery, err := helpUser.FindToDB(nil, func(r *gorm.DB, field *model.GoUser) (*gorm.DB, error) {
				r = r.Select(field.FieldNameName())
				return r, nil
			})
			if err != nil {
				return nil, err
			}

			sql := fmt.Sprintf("%v in (?)", field.FieldNameAuthor())
			r = r.Where(sql, subQuery)

			return r, nil
		})
		if err != nil {
			log.Println(err)
			return
		}

		for _, book := range objectList {
			log.Println(book)
		}
	}
}
