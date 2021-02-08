package template

const GoGoModelsField = `
package model
// 此文件是根据GoGoModelsField自动生成

{{range .Tables}}

    // {{.GoStructName}} database table name is "{{.DBTableName}}"
    type Go{{.GoStructName}} struct {
    {{range .Fields}}
        m{{.GoFieldName}} string // {{.TagString}}
    {{end}}

        cols []string // 字段列表
    }

    func NewGo{{.GoStructName}}() *Go{{.GoStructName}} {
        return &Go{{.GoStructName}} {
            {{range .Fields}}
                m{{.GoFieldName}}: "{{.DBColumnName}}",
            {{end}}
        }
    }

    func (*{{.GoStructName}}) TableName() string {
    return "{{.DBTableName}}"
    }

    // 函数
    {{range .Fields}}
        func (c *Go{{.Parent.GoStructName}}) {{.GoFieldName}}() string {
            return c.m{{.GoFieldName}}
        }
    {{end}}

    // 字段
    {{range .Fields}}
        func (c *Go{{.Parent.GoStructName}}) AddCol{{.GoFieldName}}() *Go{{.Parent.GoStructName}} {
            c.cols = append(c.cols, c.m{{.GoFieldName}})
            return c
        }
    {{end}}

    func (c *Go{{.GoStructName}}) Cols() []string {
        return c.cols
    }

    func (c *Go{{.GoStructName}}) ResetCols() {
        c.cols = []string{}
    }

    func (c *Go{{.GoStructName}}) AddAllCols() {
        c.cols = append(c.cols, "*")
    }
{{end}}

type tableColumn struct {
}

func NewTableColumn() *tableColumn {
    return &tableColumn{}
}

{{range .Tables}}
    func (*tableColumn) {{.GoStructName}}() *Go{{.GoStructName}} {
        return NewGo{{.GoStructName}}()
    }

{{end}}
`

const GoGoTypePoint = `
package model
// 此文件是根据GoGoTypePoint自动生成

import (
	"time"
)

func Time(v time.Time) *time.Time { return &v }

func Uint(v uint) *uint { return &v }

func Bool(v bool) *bool { return &v }

func Int(v int) *int { return &v }

func Int32(v int32) *int32 { return &v }

func Int64(v int64) *int64 { return &v }

func Uint32(v uint32) *uint32 { return &v }

func Uint64(v uint64) *uint64 { return &v }

func Float32(v float32) *float32 { return &v }

func Float64(v float64) *float64 { return &v }

func String(v string) *string { return &v }

func PointTime(v *time.Time) time.Time {
	if v == nil {
		return time.Time{}
	}
	return *v
}

func PointUint(v *uint) uint {
	if v == nil {
		return 0
	}
	return *v
}

func PointBool(v *bool) bool {
	if v == nil {
		return false
	}
	return *v
}

func PointInt(v *int) int {
	if v == nil {
		return 0
	}
	return *v
}

func PointInt32(v *int32) int32 {
	if v == nil {
		return 0
	}
	return *v
}

func PointInt64(v *int64) int64 {
	if v == nil {
		return 0
	}
	return *v
}

func PointUint32(v *uint32) uint32 {
	if v == nil {
		return 0
	}
	return *v
}

func PointUint64(v *uint64) uint64 {
	if v == nil {
		return 0
	}
	return *v
}

func PointFloat32(v *float32) float32 {
	if v == nil {
		return 0
	}
	return *v
}

func PointFloat64(v *float64) float64 {
	if v == nil {
		return 0
	}
	return *v
}

func PointString(v *string) string {
	if v == nil {
		return ""
	}
	return *v
}
`

const GoGoBean = `
package model

// 此文件是更加GoGoBean自动生成

import (
    "github.com/general252/gout/uerror"
    "gorm.io/gorm"
)


type bean{{.GoStructName}} struct {
    obj *{{.GoStructName}}
    db  *gorm.DB
}

type Bean{{.GoStructName}}Option func(*bean{{.GoStructName}})

func With{{.GoStructName}}DB(db *gorm.DB) Bean{{.GoStructName}}Option {
    return func(c *bean{{.GoStructName}}) {
        c.db = db
    }
}

func NewBean{{.GoStructName}}(opts ...Bean{{.GoStructName}}Option) *bean{{.GoStructName}} {
    c := &bean{{.GoStructName}}{
        obj: new({{.GoStructName}}),
    }

    for _, opt := range opts {
        opt(c)
    }

    if c.db == nil {
        // 使用默认的
        c.db = DefaultEngine.GetDB()
    }

    return c
}

func (c *bean{{.GoStructName}}) Add(m *{{.GoStructName}}) (*{{.GoStructName}}, error) {
    if c.db == nil {
        return nil, uerror.WithMessage("db is nil")
    }

    if r := c.db.Model(c.obj).Create(m); r.Error != nil {
        return nil, uerror.WithError(r.Error)
    } else if r.RowsAffected != 1 {
        return nil, uerror.WithMessageF("add bean{{.GoStructName}} affected row %v", r.RowsAffected)
    }

    return m, nil
}

func (c *bean{{.GoStructName}}) GetKeyId(keyId uint) (*{{.GoStructName}}, error) {
    if c.db == nil {
        return nil, uerror.WithMessage("db is nil")
    }

    var m {{.GoStructName}}
    r := c.db.Model(c.obj).Where(&{{.GoStructName}}{ID: keyId}).First(&m)
    if r.Error != nil {
        return nil, uerror.WithError(r.Error)
    }
    if m.ID != keyId {
        return nil, uerror.WithMessageF("get bean{{.GoStructName}} fail. param: %v, get: %v", keyId, m.ID)
    }

    return &m, nil
}

func (c *bean{{.GoStructName}}) GetUid(uid string) (*{{.GoStructName}}, error) {
    if c.db == nil {
        return nil, uerror.WithMessage("db is nil")
    }

    if len(uid) == 0 {
        return nil, uerror.WithMessage("error param uid")
    }

    var m {{.GoStructName}}
    r := c.db.Model(c.obj).Where(&{{.GoStructName}}{Uid: String(uid)}).First(&m)
    if r.Error != nil {
        return nil, uerror.WithError(r.Error)
    }
    if PointString(m.Uid) != uid {
        return nil, uerror.WithMessageF("get beanBook fail. param: %v, get: %v", uid, m.Uid)
    }

    return &m, nil
}

type Bean{{.GoStructName}}DBOption func(*gorm.DB) *gorm.DB

func (c *bean{{.GoStructName}}) WithFilter(filter *{{.GoStructName}}, limit int, offset int, customFilters ...Bean{{.GoStructName}}DBOption) ([]{{.GoStructName}}, int64, error) {
    if c.db == nil {
        return nil, -1, uerror.WithMessage("db is nil")
    }

    var funFilter = func() *gorm.DB {
        r := c.db.Model(c.obj)

        if filter != nil {
            r = r.Where(filter)
        }

        for _, filter := range customFilters {
            r = filter(r)
        }

        return r
    }

    var m []{{.GoStructName}}
    if r := funFilter().Offset(offset).Limit(limit).Find(&m); r.Error != nil {
        return nil, -1, uerror.WithError(r.Error)
    }

    var totalCount int64
    if r := funFilter().Count(&totalCount); r.Error != nil {
        return nil, -1, uerror.WithError(r.Error)
    }

    return m, totalCount, nil
}


`

const GoGoEngine =`
package model

// 此文件是更加GoGoEngine自动生成

import (
	"context"
	"fmt"
	"github.com/general252/gout/uerror"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

// DefaultEngine 默认数据链接
var DefaultEngine DBEngine = NewDBEngine()

type DBEngine interface {
	Init(p *InitParam) error
	Ping(duration time.Duration) error
	GetDB() *gorm.DB
}

func NewDBEngine() *DBEngineImp {
	return &DBEngineImp{}
}

type DBEngineImp struct {
	db *gorm.DB
}

type InitParam struct {
	Host     string
	Port     int
	DBName   string
	Username string
	Password string
	ShowSQL  bool
}

// Init 初始化gorm, 官方文档 https://gorm.io/zh_CN/docs
func (c *DBEngineImp) Init(p *InitParam) error {
	if p == nil {
		return fmt.Errorf("error param")
	}
	if err := c.mMySQLCreateDataBase(p); err != nil {
		return err
	}

	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True",
		p.Username, p.Password, p.Host, p.Port, p.DBName)

	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,   // DSN data source name
		DefaultStringSize:         256,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 change 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{
		SkipDefaultTransaction: true, // 禁用默认事务, 为了确保数据一致性，GORM 会在事务里执行写入操作（创建、更新、删除）。如果没有这方面的要求，您可以在初始化时禁用它，这将获得大约 30%+ 性能提升
	})

	if err != nil {
		return uerror.WithErrorAndMessage(err, "gorm.Open fail")
	}

	{
		// db.Logger = db.Logger.LogMode(logger.Info) // 设置日志打印级别
		// or
		db.Logger = logger.New(log.New(os.Stdout, "", log.LstdFlags), logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Warn,
			Colorful:      false,
		})
	}

	if p.ShowSQL {
		db.Logger = db.Logger.LogMode(logger.Info)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return uerror.WithErrorAndMessage(err, "db.DB() fail")
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxIdleTime(time.Hour)

	if err := c.ping(db, time.Second*10); err != nil {
		return err
	}

	err = db.AutoMigrate(GetModelBeans())
	if err != nil {
		return uerror.WithError(err)
	}

	c.db = db

	return nil
}

func (c *DBEngineImp) Ping(duration time.Duration) error {
	return c.ping(c.db, duration)
}

func (c *DBEngineImp) GetDB() *gorm.DB {
	return c.db
}

func (c *DBEngineImp) ping(db *gorm.DB, duration time.Duration) error {
	sqlDB, err := db.DB()
	if err != nil {
		return uerror.WithError(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), duration)
	defer cancel()
	if err := sqlDB.PingContext(ctx); err != nil {
		return uerror.WithError(err)
	}

	return nil
}

// mMySQLCreateDataBase 创建数据库
func (c *DBEngineImp) mMySQLCreateDataBase(p *InitParam) error {
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/information_schema?charset=utf8mb4&parseTime=True",
		p.Username, p.Password, p.Host, p.Port)

	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,   // DSN data source name
		DefaultStringSize:         256,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 change 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{})
	if err != nil {
		return err
	}

	dbName := p.DBName

	rows, err := db.Raw("SHOW databases;").Rows()
	if err != nil {
		return err
	}
	defer func() {
		_ = rows.Close()
	}()

	var find = false
	for rows.Next() {
		var name string
		_ = rows.Scan(&name)
		if name == dbName {
			find = true
		}

		//log.Printf("database name: %v", name)
	}
	if find {
		return nil
	}

	if r := db.Exec(fmt.Sprintf("CREATE DATABASE %v;", dbName)); r.Error != nil {
		return r.Error
	}

	return nil
}

`