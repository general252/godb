package main

import (
	"github.com/general252/godb/model"
	"log"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	model.Build()
}

var _ = `
func test() {

	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True",
		"root", "sedb@20120711", "192.168.6.59", 9306, "db_test")
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:               dsn,
		DefaultStringSize: 256, // string 类型字段的默认长度
		//DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		//DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		//DontSupportRenameColumn:   true,  // 用 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{
		SkipDefaultTransaction: true,
	})

	if err != nil {
		return
	}

	db.Logger = db.Logger.LogMode(logger.Info)

	_ = db.AutoMigrate(new(model.Book))
	_ = db.AutoMigrate(new(model.User))

	var tony = &model.User{
		Uid:       nil,
		Name:      model.String("tony"),
		Age:       model.Int(20),
		Birthday:  model.Time(time.Now()),
		ManagerID: model.Uint(1230),
	}
	r := db.Create(tony)
	if r.Error != nil {
	}

	var tmp []model.User
	db.Where(&model.User{
		CompanyID: model.Uint(0),
	}).First(&tmp)

	for _, m := range tmp {
		log.Println(m.Birthday.In(time.Local).Format(time.RFC3339))
	}
	strJson, _ := json.MarshalIndent(&tmp, "", "  ")
	log.Printf("%v", string(strJson))
}
`
