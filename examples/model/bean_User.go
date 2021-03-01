package model

// 此文件是更加GoGoBean自动生成

import (
    "github.com/general252/godb/godb"
    "github.com/general252/gout/uerror"
    "gorm.io/gorm"
    "time"
)

type beanUser struct {
    obj *User
    db  *gorm.DB
}

type BeanUserOption func(*beanUser)

func WithUserDB(db *gorm.DB) BeanUserOption {
    return func(c *beanUser) {
        c.db = db
    }
}

func NewBeanUser(opts ...BeanUserOption) *beanUser {
    c := &beanUser{
        obj: new(User),
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

func (c *beanUser) Add(m *User) (*User, error) {
    if c.db == nil {
        return nil, uerror.WithMessage("[User] db is nil")
    }

    if m == nil {
        return nil, uerror.WithMessage("[User] param is nil")
    }

    m.CreatedAt = Time(time.Now().UTC())

    if r := c.db.Model(c.obj).Create(m); r.Error != nil {
        return nil, uerror.WithError(r.Error)
    } else if r.RowsAffected != 1 {
        return nil, uerror.WithMessageF("[User] add affected row %v", r.RowsAffected)
    }

    return m, nil
}

func (c *beanUser) GetByKeyId(keyId uint) (*User, error) {
    if c.db == nil {
        return nil, uerror.WithMessage("[User] db is nil")
    }

    var m User
    r := c.db.Model(c.obj).Where(&User{
        Model: godb.Model{
            ID: keyId,
        },
    }).First(&m)
    if r.Error != nil {
        return nil, uerror.WithError(r.Error)
    }
    if m.ID != keyId {
        return nil, uerror.WithMessageF("[User] get fail. param: %v, get: %v", keyId, m.ID)
    }

    return &m, nil
}

func (c *beanUser) GetUid(uid string) (*User, error) {
    if c.db == nil {
        return nil, uerror.WithMessage("[User] db is nil")
    }

    if len(uid) == 0 {
        return nil, uerror.WithMessage("[User] error param uid")
    }

    var m User
    r := c.db.Model(c.obj).Where(&User{
        Model: godb.Model{
            Uid: String(uid),
        },
    }).First(&m)
    if r.Error != nil {
        return nil, uerror.WithError(r.Error)
    }
    if PointString(m.Uid) != uid {
        return nil, uerror.WithMessageF("[User] get beanBook fail. param: %v, get: %v", uid, m.Uid)
    }

    return &m, nil
}

func (c *beanUser) UpdateByKeyId(obj *User) error {
    if c.db == nil {
        return uerror.WithMessage("[User] db is nil")
    }
    if obj == nil || obj.ID == 0 {
        return uerror.WithMessage("[User] param error")
    }

    obj.UpdatedAt = Time(time.Now().UTC())

    r := c.db.Model(c.obj).Where(&User{
        Model: godb.Model{
            ID: obj.ID,
        },
    }).Updates(obj)
    if r.Error != nil {
        return uerror.WithError(r.Error)
    }

    return nil
}

func (c *beanUser) UpdateByUId(obj *User) error {
    if c.db == nil {
        return uerror.WithMessage("[User] db is nil")
    }
    if obj == nil || len(PointString(obj.Uid)) == 0 {
        return uerror.WithMessage("[User] param error")
    }

    obj.UpdatedAt = Time(time.Now().UTC())

    r := c.db.Model(c.obj).Where(&User{
        Model: godb.Model{
            Uid: obj.Uid,
        },
    }).Updates(obj)
    if r.Error != nil {
        return uerror.WithError(r.Error)
    }

    return nil
}

func (c *beanUser) DeleteByKeyId(keyId uint) error {
    if c.db == nil {
        return uerror.WithMessage("[User] db is nil")
    }

    if keyId == 0 {
        return uerror.WithMessage("[User] key id is 0")
    }

    var m User
    r := c.db.Model(c.obj).Where(&User{
        Model: godb.Model{
            ID: keyId,
        },
    }).Delete(&m)
    if r.Error != nil {
        return uerror.WithError(r.Error)
    }

    if r.RowsAffected != 1 {
        return uerror.WithMessageF("[User] delete row affected %v", r.RowsAffected)
    }

    return nil
}

func (c *beanUser) DeleteByUId(uid string) error {
    if c.db == nil {
        return uerror.WithMessage("[User] db is nil")
    }

    if len(uid) == 0 {
        return uerror.WithMessage("[User] uid error")
    }

    var m User
    r := c.db.Model(c.obj).Where(&User{
        Model: godb.Model{
            Uid: String(uid),
        },
    }).Delete(&m)
    if r.Error != nil {
        return uerror.WithError(r.Error)
    }

    if r.RowsAffected != 1 {
        return uerror.WithMessageF("[User] delete row affected %v", r.RowsAffected)
    }

    return nil
}

type BeanUserDBOption func(r *gorm.DB, field *GoUser) *gorm.DB

func (c *beanUser) Filter(filter *User, limit int, offset int, customFilters ...BeanUserDBOption) ([]User, int64, error) {
    if c.db == nil {
        return nil, -1, uerror.WithMessage("[User] db is nil")
    }

    var funFilter = func() *gorm.DB {
        r := c.db.Model(c.obj)

        if filter != nil {
            r = r.Where(filter)
        }

        if len(customFilters) > 0 {
            var field = newGoUser()
            for _, filter := range customFilters {
                r = filter(r, field)
            }
        }

        return r
    }

    var m []User
    if r := funFilter().Offset(offset).Limit(limit).Find(&m); r.Error != nil {
        return nil, -1, uerror.WithError(r.Error)
    }

    var totalCount int64
    if r := funFilter().Count(&totalCount); r.Error != nil {
        return nil, -1, uerror.WithError(r.Error)
    }

    return m, totalCount, nil
}

func (c *beanUser) Find(filter *User, limit int, offset int, customFilters ...BeanUserDBOption) ([]User, error) {
    if c.db == nil {
        return nil, uerror.WithMessage("[User] db is nil")
    }

    var funFilter = func() *gorm.DB {
        r := c.db.Model(c.obj)

        if filter != nil {
            r = r.Where(filter)
        }

        if len(customFilters) > 0 {
            var field = newGoUser()
            for _, filter := range customFilters {
                r = filter(r, field)
            }
        }

        return r
    }

    var m []User
    if r := funFilter().Offset(offset).Limit(limit).Find(&m); r.Error != nil {
        return nil, uerror.WithError(r.Error)
    }

    return m, nil
}
