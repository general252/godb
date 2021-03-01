package model

// 此文件是更加GoGoBean自动生成

import (
    "github.com/general252/godb/godb"
    "github.com/general252/gout/uerror"
    "gorm.io/gorm"
    "time"
)

type beanBook struct {
    obj *Book
    db  *gorm.DB
}

type BeanBookOption func(*beanBook)

func WithBookDB(db *gorm.DB) BeanBookOption {
    return func(c *beanBook) {
        c.db = db
    }
}

func NewBeanBook(opts ...BeanBookOption) *beanBook {
    c := &beanBook{
        obj: new(Book),
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

func (c *beanBook) Add(m *Book) (*Book, error) {
    if c.db == nil {
        return nil, uerror.WithMessage("[Book] db is nil")
    }

    if m == nil {
        return nil, uerror.WithMessage("[Book] param is nil")
    }

    m.CreatedAt = Time(time.Now().UTC())

    if r := c.db.Model(c.obj).Create(m); r.Error != nil {
        return nil, uerror.WithError(r.Error)
    } else if r.RowsAffected != 1 {
        return nil, uerror.WithMessageF("[Book] add affected row %v", r.RowsAffected)
    }

    return m, nil
}

func (c *beanBook) GetByKeyId(keyId uint) (*Book, error) {
    if c.db == nil {
        return nil, uerror.WithMessage("[Book] db is nil")
    }

    var m Book
    r := c.db.Model(c.obj).Where(&Book{
        Model: godb.Model{
            ID: keyId,
        },
    }).First(&m)
    if r.Error != nil {
        return nil, uerror.WithError(r.Error)
    }
    if m.ID != keyId {
        return nil, uerror.WithMessageF("[Book] get fail. param: %v, get: %v", keyId, m.ID)
    }

    return &m, nil
}

func (c *beanBook) GetUid(uid string) (*Book, error) {
    if c.db == nil {
        return nil, uerror.WithMessage("[Book] db is nil")
    }

    if len(uid) == 0 {
        return nil, uerror.WithMessage("[Book] error param uid")
    }

    var m Book
    r := c.db.Model(c.obj).Where(&Book{
        Model: godb.Model{
            Uid: String(uid),
        },
    }).First(&m)
    if r.Error != nil {
        return nil, uerror.WithError(r.Error)
    }
    if PointString(m.Uid) != uid {
        return nil, uerror.WithMessageF("[Book] get beanBook fail. param: %v, get: %v", uid, m.Uid)
    }

    return &m, nil
}

func (c *beanBook) UpdateByKeyId(obj *Book) error {
    if c.db == nil {
        return uerror.WithMessage("[Book] db is nil")
    }
    if obj == nil || obj.ID == 0 {
        return uerror.WithMessage("[Book] param error")
    }

    obj.UpdatedAt = Time(time.Now().UTC())

    r := c.db.Model(c.obj).Where(&Book{
        Model: godb.Model{
            ID: obj.ID,
        },
    }).Updates(obj)
    if r.Error != nil {
        return uerror.WithError(r.Error)
    }

    return nil
}

func (c *beanBook) UpdateByUId(obj *Book) error {
    if c.db == nil {
        return uerror.WithMessage("[Book] db is nil")
    }
    if obj == nil || len(PointString(obj.Uid)) == 0 {
        return uerror.WithMessage("[Book] param error")
    }

    obj.UpdatedAt = Time(time.Now().UTC())

    r := c.db.Model(c.obj).Where(&Book{
        Model: godb.Model{
            Uid: obj.Uid,
        },
    }).Updates(obj)
    if r.Error != nil {
        return uerror.WithError(r.Error)
    }

    return nil
}

func (c *beanBook) DeleteByKeyId(keyId uint) error {
    if c.db == nil {
        return uerror.WithMessage("[Book] db is nil")
    }

    if keyId == 0 {
        return uerror.WithMessage("[Book] key id is 0")
    }

    var m Book
    r := c.db.Model(c.obj).Where(&Book{
        Model: godb.Model{
            ID: keyId,
        },
    }).Delete(&m)
    if r.Error != nil {
        return uerror.WithError(r.Error)
    }

    if r.RowsAffected != 1 {
        return uerror.WithMessageF("[Book] delete row affected %v", r.RowsAffected)
    }

    return nil
}

func (c *beanBook) DeleteByUId(uid string) error {
    if c.db == nil {
        return uerror.WithMessage("[Book] db is nil")
    }

    if len(uid) == 0 {
        return uerror.WithMessage("[Book] uid error")
    }

    var m Book
    r := c.db.Model(c.obj).Where(&Book{
        Model: godb.Model{
            Uid: String(uid),
        },
    }).Delete(&m)
    if r.Error != nil {
        return uerror.WithError(r.Error)
    }

    if r.RowsAffected != 1 {
        return uerror.WithMessageF("[Book] delete row affected %v", r.RowsAffected)
    }

    return nil
}

type BeanBookDBOption func(r *gorm.DB, field *GoBook) *gorm.DB

func (c *beanBook) Filter(filter *Book, limit int, offset int, customFilters ...BeanBookDBOption) ([]Book, int64, error) {
    if c.db == nil {
        return nil, -1, uerror.WithMessage("[Book] db is nil")
    }

    var funFilter = func() *gorm.DB {
        r := c.db.Model(c.obj)

        if filter != nil {
            r = r.Where(filter)
        }

        if len(customFilters) > 0 {
            var field = newGoBook()
            for _, filter := range customFilters {
                r = filter(r, field)
            }
        }

        return r
    }

    var m []Book
    if r := funFilter().Offset(offset).Limit(limit).Find(&m); r.Error != nil {
        return nil, -1, uerror.WithError(r.Error)
    }

    var totalCount int64
    if r := funFilter().Count(&totalCount); r.Error != nil {
        return nil, -1, uerror.WithError(r.Error)
    }

    return m, totalCount, nil
}

func (c *beanBook) Find(filter *Book, limit int, offset int, customFilters ...BeanBookDBOption) ([]Book, error) {
    if c.db == nil {
        return nil, uerror.WithMessage("[Book] db is nil")
    }

    var funFilter = func() *gorm.DB {
        r := c.db.Model(c.obj)

        if filter != nil {
            r = r.Where(filter)
        }

        if len(customFilters) > 0 {
            var field = newGoBook()
            for _, filter := range customFilters {
                r = filter(r, field)
            }
        }

        return r
    }

    var m []Book
    if r := funFilter().Offset(offset).Limit(limit).Find(&m); r.Error != nil {
        return nil, uerror.WithError(r.Error)
    }

    return m, nil
}
