package model

// 此文件是更加GoGoBean自动生成

import (
	"time"

	"github.com/general252/godb/godb"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type beanBook struct {
	obj *Book
	db  *gorm.DB
}

type BeanBookOption func(*beanBook)

// WithBookDB 设置DB
func WithBookDB(db *gorm.DB) BeanBookOption {
	return func(c *beanBook) {
		c.db = db
	}
}

// NewBeanBook 创建Book辅助
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

// Add 添加一条Book记录
func (c *beanBook) Add(m *Book) (*Book, error) {
	if c.db == nil {
		return nil, errors.New("[Book] db is nil")
	}

	if m == nil {
		return nil, errors.New("[Book] param is nil")
	}

	m.CreatedAt = godb.Time(time.Now().UTC())

	if r := c.db.Model(c.obj).Create(m); r.Error != nil {
		return nil, errors.WithStack(r.Error)
	} else if r.RowsAffected != 1 {
		return nil, errors.Errorf("[Book] add affected row %v", r.RowsAffected)
	}

	return m, nil
}

// GetByKeyId 根据keyId获取Book
func (c *beanBook) GetByKeyId(keyId uint) (*Book, error) {
	if c.db == nil {
		return nil, errors.New("[Book] db is nil")
	}

	var m Book
	r := c.db.Model(c.obj).Where(&Book{
		Model: godb.Model{
			ID: keyId,
		},
	}).First(&m)
	if r.Error != nil {
		return nil, errors.WithStack(r.Error)
	}
	if m.ID != keyId {
		return nil, errors.Errorf("[Book] get fail. param: %v, get: %v", keyId, m.ID)
	}

	return &m, nil
}

// GetUid 根据uid获取Book
func (c *beanBook) GetUid(uid string) (*Book, error) {
	if c.db == nil {
		return nil, errors.New("[Book] db is nil")
	}

	if len(uid) == 0 {
		return nil, errors.New("[Book] error param uid")
	}

	var m Book
	r := c.db.Model(c.obj).Where(&Book{
		Model: godb.Model{
			Uid: godb.String(uid),
		},
	}).First(&m)
	if r.Error != nil {
		return nil, errors.WithStack(r.Error)
	}
	if godb.PointString(m.Uid) != uid {
		return nil, errors.Errorf("[Book] get beanBook fail. param: %v, get: %v", uid, m.Uid)
	}

	return &m, nil
}

// UpdateByKeyId 根据keyId修改Book
func (c *beanBook) UpdateByKeyId(obj *Book) error {
	if c.db == nil {
		return errors.New("[Book] db is nil")
	}
	if obj == nil || obj.ID == 0 {
		return errors.New("[Book] param error")
	}

	obj.UpdatedAt = godb.Time(time.Now().UTC())

	r := c.db.Model(c.obj).Where(&Book{
		Model: godb.Model{
			ID: obj.ID,
		},
	}).Updates(obj)
	if r.Error != nil {
		return errors.WithStack(r.Error)
	}

	return nil
}

// UpdateByUId 根据uid修改Book
func (c *beanBook) UpdateByUId(obj *Book) error {
	if c.db == nil {
		return errors.New("[Book] db is nil")
	}
	if obj == nil || len(godb.PointString(obj.Uid)) == 0 {
		return errors.New("[Book] param error")
	}

	obj.UpdatedAt = godb.Time(time.Now().UTC())

	r := c.db.Model(c.obj).Where(&Book{
		Model: godb.Model{
			Uid: obj.Uid,
		},
	}).Updates(obj)
	if r.Error != nil {
		return errors.WithStack(r.Error)
	}

	return nil
}

// Update 根据自定义条件更新Book
func (c *beanBook) Update(m *Book, customFilters ...BeanBookDBOption) (int64, error) {
	if c.db == nil {
		return -1, errors.New("[Book] db is nil")
	}

	var (
		err error
		r   = c.db.Model(c.obj)
	)

	var field = newGoBook()
	for _, filter := range customFilters {
		if r, err = filter(r, field); err != nil {
			return -1, err
		}
	}

	r = r.Updates(m)
	if r.Error != nil {
		return -1, errors.WithStack(r.Error)
	}

	return r.RowsAffected, nil
}

// DeleteByKeyId 根据keyId删除Book
func (c *beanBook) DeleteByKeyId(keyId uint) error {
	if c.db == nil {
		return errors.New("[Book] db is nil")
	}

	if keyId == 0 {
		return errors.New("[Book] key id is 0")
	}

	var m Book
	r := c.db.Model(c.obj).Where(&Book{
		Model: godb.Model{
			ID: keyId,
		},
	}).Delete(&m)
	if r.Error != nil {
		return errors.WithStack(r.Error)
	}

	if r.RowsAffected != 1 {
		return errors.Errorf("[Book] delete row affected %v", r.RowsAffected)
	}

	return nil
}

// DeleteByUId 根据uid删除Book
func (c *beanBook) DeleteByUId(uid string) error {
	if c.db == nil {
		return errors.New("[Book] db is nil")
	}

	if len(uid) == 0 {
		return errors.New("[Book] uid error")
	}

	var m Book
	r := c.db.Model(c.obj).Where(&Book{
		Model: godb.Model{
			Uid: godb.String(uid),
		},
	}).Delete(&m)
	if r.Error != nil {
		return errors.WithStack(r.Error)
	}

	if r.RowsAffected != 1 {
		return errors.Errorf("[Book] delete row affected %v", r.RowsAffected)
	}

	return nil
}

// Delete 根据自定义条件删除Book
func (c *beanBook) Delete(filter *BookFilter, customFilters ...BeanBookDBOption) (int64, error) {
	if c.db == nil {
		return -1, errors.New("[Book] db is nil")
	}

	var (
		err error
		m   Book
		r   = c.db.Model(c.obj)
	)

	if filter != nil {
		r, err = c.funFilter(filter, customFilters...)
		if err != nil {
			return -1, err
		}
	}

	var field = newGoBook()
	for _, filterFn := range customFilters {
		if r, err = filterFn(r, field); err != nil {
			return -1, err
		}
	}

	r = r.Delete(&m)
	if r.Error != nil {
		return -1, errors.WithStack(r.Error)
	}

	return r.RowsAffected, nil
}

type BeanBookDBOption func(r *gorm.DB, field *GoBook) (*gorm.DB, error)

type BookFilter struct {
	Limit       int
	Offset      int
	Object      *Book            // 完全匹配
	ObjectNeq   *Book            // Neq
	Match       *Book            // 模糊匹配, 仅string类型字段有效, 推荐使用 Cond.Like
	MatchOr     *Book            // 模糊匹配, 仅string类型字段有效, 推荐使用 Cond.Like
	MatchString *string          // 模糊匹配所有string类型字段
	CondList    []BookFilterCond //
}

// Filter 筛选Book
func (c *beanBook) Filter(filter *BookFilter, customFilters ...BeanBookDBOption) ([]Book, int64, error) {
	if c.db == nil {
		return nil, -1, errors.New("[Book] db is nil")
	}

	var limit = 0
	var offset = 0
	if filter != nil {
		limit = filter.Limit
		offset = filter.Offset
	}
	if limit < 0 {
		return nil, -1, errors.New("[Book] limit error")
	}
	if offset < 0 {
		return nil, -1, errors.New("[Book] offset error")
	}

	var m []Book
	r, err := c.funFilter(filter, customFilters...)
	if err != nil {
		return nil, -1, err
	}
	r = r.Offset(offset)
	if limit > 0 {
		r = r.Limit(limit)
	}
	if r.Find(&m); r.Error != nil {
		return nil, -1, errors.WithStack(r.Error)
	}

	var totalCount int64
	{
		r, err := c.funFilter(filter, customFilters...)
		if err != nil {
			return nil, -1, err
		}
		if r := r.Count(&totalCount); r.Error != nil {
			return nil, -1, errors.WithStack(r.Error)
		}
	}

	return m, totalCount, nil
}

// Find 查找Book, matchObject: 仅string类型字段有效
func (c *beanBook) Find(filter *BookFilter, customFilters ...BeanBookDBOption) ([]Book, error) {
	if c.db == nil {
		return nil, errors.New("[Book] db is nil")
	}

	var limit = 0
	var offset = 0
	if filter != nil {
		limit = filter.Limit
		offset = filter.Offset
	}
	if limit < 0 {
		return nil, errors.New("[Book] limit error")
	}
	if offset < 0 {
		return nil, errors.New("[Book] offset error")
	}

	var m []Book
	r, err := c.funFilter(filter, customFilters...)
	if err != nil {
		return nil, err
	}
	r = r.Offset(offset)
	if limit > 0 {
		r = r.Limit(limit)
	}
	if r.Find(&m); r.Error != nil {
		return nil, errors.WithStack(r.Error)
	}

	return m, nil
}

// FindToDB 查找Book, matchObject: 仅string类型字段有效
//
//	subQuery, _ := FindToDB(***, func(***) {
//	    r = r.Select(field.FieldName***())
//	})
//
// r = r.Where(field.FieldNameXXXUid() + " IN (?)", subQuery)
func (c *beanBook) FindToDB(filter *BookFilter, customFilters ...BeanBookDBOption) (*gorm.DB, error) {
	if c.db == nil {
		return nil, errors.New("[Book] db is nil")
	}

	var limit = 0
	var offset = 0
	if filter != nil {
		limit = filter.Limit
		offset = filter.Offset
	}
	if limit < 0 {
		return nil, errors.New("[Book] limit error")
	}
	if offset < 0 {
		return nil, errors.New("[Book] offset error")
	}

	r, err := c.funFilter(filter, customFilters...)
	if err != nil {
		return nil, err
	}

	r = r.Offset(offset)
	if limit > 0 {
		r = r.Limit(limit)
	}

	return r, nil
}

// FindAndRange 查找Book, matchObject: 仅string类型字段有效
func (c *beanBook) FindAndRange(fun func(*Book) bool, filter *BookFilter, customFilters ...BeanBookDBOption) error {
	if c.db == nil {
		return errors.New("[Book] db is nil")
	}
	if fun == nil {
		return errors.New("[Book] param error")
	}

	var limit = 0
	var offset = 0
	if filter != nil {
		limit = filter.Limit
		offset = filter.Offset
	}
	if limit < 0 {
		return errors.New("[Book] limit error")
	}
	if offset < 0 {
		return errors.New("[Book] offset error")
	}

	r, err := c.funFilter(filter, customFilters...)
	if err != nil {
		return err
	}
	r = r.Offset(offset)
	if limit > 0 {
		r = r.Limit(limit)
	}

	rows, err := r.Rows()
	if err != nil {
		return errors.WithStack(err)
	}
	defer rows.Close()

	for rows.Next() {
		var object Book
		if err := r.ScanRows(rows, &object); err != nil {
			return errors.WithStack(err)
		}

		if fun(&object) {
			continue
		} else {
			break
		}
	}

	return nil
}

// First Book, matchObject: 仅string类型字段有效
func (c *beanBook) First(filter *BookFilter, customFilters ...BeanBookDBOption) (*Book, error) {
	if c.db == nil {
		return nil, errors.New("[Book] db is nil")
	}

	var offset = 0
	if filter != nil {
		offset = filter.Offset
	}
	if offset < 0 {
		return nil, errors.New("[Book] offset error")
	}

	var m Book
	r, err := c.funFilter(filter, customFilters...)
	if err != nil {
		return nil, err
	}
	if r := r.Offset(offset).First(&m); r.Error != nil {
		return nil, errors.WithStack(r.Error)
	}

	return &m, nil
}

// Last Book, matchObject: 仅string类型字段有效
func (c *beanBook) Last(filter *BookFilter, customFilters ...BeanBookDBOption) (*Book, error) {
	if c.db == nil {
		return nil, errors.New("[Book] db is nil")
	}

	var offset = 0
	if filter != nil {
		offset = filter.Offset
	}
	if offset < 0 {
		return nil, errors.New("[Book] offset error")
	}

	var m Book
	r, err := c.funFilter(filter, customFilters...)
	if err != nil {
		return nil, err
	}
	if r := r.Offset(offset).Last(&m); r.Error != nil {
		return nil, errors.WithStack(r.Error)
	}

	return &m, nil
}

// Has 判断是否存在 Book
func (c *beanBook) Has(filter *BookFilter, customFilters ...BeanBookDBOption) (bool, error) {
	if c.db == nil {
		return false, errors.New("[Book] db is nil")
	}

	var totalCount int64
	r, err := c.funFilter(filter, customFilters...)
	if err != nil {
		return false, err
	}
	if r := r.Limit(1).Count(&totalCount); r.Error != nil {
		return false, errors.WithStack(r.Error)
	}

	if totalCount > 0 {
		return true, nil
	} else {
		return false, nil
	}
}

// Count 筛选数量 Book
func (c *beanBook) Count(filter *BookFilter, customFilters ...BeanBookDBOption) (int64, error) {
	if c.db == nil {
		return -1, errors.New("[Book] db is nil")
	}

	var totalCount int64
	r, err := c.funFilter(filter, customFilters...)
	if err != nil {
		return 0, err
	}
	if r := r.Count(&totalCount); r.Error != nil {
		return -1, errors.WithStack(r.Error)
	}

	return totalCount, nil
}

func (c *beanBook) funFilter(filter *BookFilter, customFilters ...BeanBookDBOption) (*gorm.DB, error) {

	var (
		err   error
		r     = c.db.Model(c.obj)
		field = newGoBook()
	)

	if filter != nil {
		if filter.Object != nil {
			r = r.Where(filter.Object)
		}

		var objectNeq = filter.ObjectNeq
		if objectNeq != nil {

			if objectNeq.ID != 0 {
				r = r.Clauses(clause.Neq{
					Column: field.FieldNameID(),
					Value:  objectNeq.ID,
				})
			}

			if objectNeq.Uid != nil {
				r = r.Clauses(clause.Neq{
					Column: field.FieldNameUid(),
					Value:  objectNeq.Uid,
				})
			}

			if objectNeq.CreatedAt != nil {
				r = r.Clauses(clause.Neq{
					Column: field.FieldNameCreatedAt(),
					Value:  objectNeq.CreatedAt,
				})
			}

			if objectNeq.UpdatedAt != nil {
				r = r.Clauses(clause.Neq{
					Column: field.FieldNameUpdatedAt(),
					Value:  objectNeq.UpdatedAt,
				})
			}

			if objectNeq.Author != nil {
				r = r.Clauses(clause.Neq{
					Column: field.FieldNameAuthor(),
					Value:  objectNeq.Author,
				})
			}

			if objectNeq.Name != nil {
				r = r.Clauses(clause.Neq{
					Column: field.FieldNameName(),
					Value:  objectNeq.Name,
				})
			}

		}

		var match = filter.Match
		if match != nil {

			if match.Uid != nil {
				r = r.Clauses(clause.Like{
					Column: field.FieldNameUid(),
					Value:  "%" + godb.PointString(match.Uid) + "%",
				})
			}

			if match.Author != nil {
				r = r.Clauses(clause.Like{
					Column: field.FieldNameAuthor(),
					Value:  "%" + godb.PointString(match.Author) + "%",
				})
			}

			if match.Name != nil {
				r = r.Clauses(clause.Like{
					Column: field.FieldNameName(),
					Value:  "%" + godb.PointString(match.Name) + "%",
				})
			}

		}

		var matchOr = filter.MatchOr
		if matchOr != nil {
			var conds []clause.Expression

			if matchOr.Uid != nil {
				conds = append(conds, clause.Like{
					Column: field.FieldNameUid(),
					Value:  "%" + godb.PointString(matchOr.Uid) + "%",
				})
			}

			if matchOr.Author != nil {
				conds = append(conds, clause.Like{
					Column: field.FieldNameAuthor(),
					Value:  "%" + godb.PointString(matchOr.Author) + "%",
				})
			}

			if matchOr.Name != nil {
				conds = append(conds, clause.Like{
					Column: field.FieldNameName(),
					Value:  "%" + godb.PointString(matchOr.Name) + "%",
				})
			}

			if len(conds) > 0 {
				r = r.Clauses(clause.Or(conds...))
			}
		}

		var matchStr = godb.PointString(filter.MatchString)
		if len(matchStr) > 0 {
			var exps []clause.Expression

			exps = append(exps, clause.Like{
				Column: field.FieldNameUid(),
				Value:  "%" + matchStr + "%",
			})

			exps = append(exps, clause.Like{
				Column: field.FieldNameAuthor(),
				Value:  "%" + matchStr + "%",
			})

			exps = append(exps, clause.Like{
				Column: field.FieldNameName(),
				Value:  "%" + matchStr + "%",
			})

			if len(exps) > 0 {
				e := clause.Or(exps...)
				r.Clauses(e)
			}
		}

		for _, cond := range filter.CondList {
			switch cond.Type {
			case CondTypeGt:
				if cond.Number != nil {

					if cond.Number.ID != nil {
						r = r.Clauses(clause.Gt{
							Column: field.FieldNameID(),
							Value:  cond.Number.ID,
						})
					}

				}
			case CondTypeGte: // 大于等于, int
				if cond.Number != nil {

					if cond.Number.ID != nil {
						r = r.Clauses(clause.Gte{
							Column: field.FieldNameID(),
							Value:  cond.Number.ID,
						})
					}

				}
			case CondTypeLt: // 小于, int
				if cond.Number != nil {

					if cond.Number.ID != nil {
						r = r.Clauses(clause.Lt{
							Column: field.FieldNameID(),
							Value:  cond.Number.ID,
						})
					}

				}
			case CondTypeLte: // 小于等于, int
				if cond.Number != nil {

					if cond.Number.ID != nil {
						r = r.Clauses(clause.Lte{
							Column: field.FieldNameID(),
							Value:  cond.Number.ID,
						})
					}

				}
			case CondTypeBetween: // 区间, int
				if cond.Between != nil {

					if cond.Between.IDLessVal != nil {
						r = r.Clauses(clause.Gte{
							Column: field.FieldNameID(),
							Value:  cond.Between.IDLessVal,
						})
					}

					if cond.Between.IDMoreVal != nil {
						r = r.Clauses(clause.Lte{
							Column: field.FieldNameID(),
							Value:  cond.Between.IDMoreVal,
						})
					}

				}
			case CondTypeIn: // in, 集合
				if cond.Container != nil {

					if len(cond.Container.IDList) > 0 {
						var values = make([]interface{}, 0, len(cond.Container.IDList))
						for _, value := range cond.Container.IDList {
							values = append(values, value)
						}

						r = r.Clauses(clause.IN{
							Column: field.FieldNameID(),
							Values: values,
						})
					}

					if len(cond.Container.UidList) > 0 {
						var values = make([]interface{}, 0, len(cond.Container.UidList))
						for _, value := range cond.Container.UidList {
							values = append(values, value)
						}

						r = r.Clauses(clause.IN{
							Column: field.FieldNameUid(),
							Values: values,
						})
					}

					if len(cond.Container.CreatedAtList) > 0 {
						var values = make([]interface{}, 0, len(cond.Container.CreatedAtList))
						for _, value := range cond.Container.CreatedAtList {
							values = append(values, value)
						}

						r = r.Clauses(clause.IN{
							Column: field.FieldNameCreatedAt(),
							Values: values,
						})
					}

					if len(cond.Container.UpdatedAtList) > 0 {
						var values = make([]interface{}, 0, len(cond.Container.UpdatedAtList))
						for _, value := range cond.Container.UpdatedAtList {
							values = append(values, value)
						}

						r = r.Clauses(clause.IN{
							Column: field.FieldNameUpdatedAt(),
							Values: values,
						})
					}

					if len(cond.Container.AuthorList) > 0 {
						var values = make([]interface{}, 0, len(cond.Container.AuthorList))
						for _, value := range cond.Container.AuthorList {
							values = append(values, value)
						}

						r = r.Clauses(clause.IN{
							Column: field.FieldNameAuthor(),
							Values: values,
						})
					}

					if len(cond.Container.NameList) > 0 {
						var values = make([]interface{}, 0, len(cond.Container.NameList))
						for _, value := range cond.Container.NameList {
							values = append(values, value)
						}

						r = r.Clauses(clause.IN{
							Column: field.FieldNameName(),
							Values: values,
						})
					}

				}
			case CondTypeNotIn: // not in, 集合
				if cond.Container != nil {

					if len(cond.Container.IDList) > 0 {
						var values = make([]interface{}, 0, len(cond.Container.IDList))
						for _, value := range cond.Container.IDList {
							values = append(values, value)
						}

						r = r.Clauses(clause.Not(
							clause.IN{
								Column: field.FieldNameID(),
								Values: values,
							},
						))
					}

					if len(cond.Container.UidList) > 0 {
						var values = make([]interface{}, 0, len(cond.Container.UidList))
						for _, value := range cond.Container.UidList {
							values = append(values, value)
						}

						r = r.Clauses(clause.Not(
							clause.IN{
								Column: field.FieldNameUid(),
								Values: values,
							},
						))
					}

					if len(cond.Container.CreatedAtList) > 0 {
						var values = make([]interface{}, 0, len(cond.Container.CreatedAtList))
						for _, value := range cond.Container.CreatedAtList {
							values = append(values, value)
						}

						r = r.Clauses(clause.Not(
							clause.IN{
								Column: field.FieldNameCreatedAt(),
								Values: values,
							},
						))
					}

					if len(cond.Container.UpdatedAtList) > 0 {
						var values = make([]interface{}, 0, len(cond.Container.UpdatedAtList))
						for _, value := range cond.Container.UpdatedAtList {
							values = append(values, value)
						}

						r = r.Clauses(clause.Not(
							clause.IN{
								Column: field.FieldNameUpdatedAt(),
								Values: values,
							},
						))
					}

					if len(cond.Container.AuthorList) > 0 {
						var values = make([]interface{}, 0, len(cond.Container.AuthorList))
						for _, value := range cond.Container.AuthorList {
							values = append(values, value)
						}

						r = r.Clauses(clause.Not(
							clause.IN{
								Column: field.FieldNameAuthor(),
								Values: values,
							},
						))
					}

					if len(cond.Container.NameList) > 0 {
						var values = make([]interface{}, 0, len(cond.Container.NameList))
						for _, value := range cond.Container.NameList {
							values = append(values, value)
						}

						r = r.Clauses(clause.Not(
							clause.IN{
								Column: field.FieldNameName(),
								Values: values,
							},
						))
					}

				}
			case CondTypeLike: // like, string
				if cond.Like != nil {

					if cond.Like.Uid != nil {
						r = r.Clauses(clause.Like{
							Column: field.FieldNameUid(),
							Value:  "%" + godb.PointString(cond.Like.Uid) + "%",
						})
					}

					if cond.Like.Author != nil {
						r = r.Clauses(clause.Like{
							Column: field.FieldNameAuthor(),
							Value:  "%" + godb.PointString(cond.Like.Author) + "%",
						})
					}

					if cond.Like.Name != nil {
						r = r.Clauses(clause.Like{
							Column: field.FieldNameName(),
							Value:  "%" + godb.PointString(cond.Like.Name) + "%",
						})
					}

				}
			}
		}

	}

	if len(customFilters) > 0 {
		for _, filter := range customFilters {
			if filter == nil {
				continue
			}

			r, err = filter(r, field)
			if err != nil {
				return r, err
			}
		}
	}

	return r, err
}

// BookFilterCond cond
type BookFilterCond struct {
	Type      CondType
	Number    *BookCondNumber
	Between   *BookCondBetween
	Container *BookCondContainer
	Like      *BookLike
}

type BookCondNumber struct {
	ID *uint // ID,

}

type BookCondBetween struct {
	IDLessVal *uint // ID less value,
	IDMoreVal *uint // ID more value,

}

type BookCondContainer struct {
	IDList []uint // ID slice,

	UidList []string // Uid slice,

	CreatedAtList []time.Time // CreatedAt slice,

	UpdatedAtList []time.Time // UpdatedAt slice,

	AuthorList []string // Author slice,

	NameList []string // Name slice,

}

type BookLike struct {
	Uid *string // Uid,

	Author *string // Author,

	Name *string // Name,

}

type beanUser struct {
	obj *User
	db  *gorm.DB
}

type BeanUserOption func(*beanUser)

// WithUserDB 设置DB
func WithUserDB(db *gorm.DB) BeanUserOption {
	return func(c *beanUser) {
		c.db = db
	}
}

// NewBeanUser 创建User辅助
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

// Add 添加一条User记录
func (c *beanUser) Add(m *User) (*User, error) {
	if c.db == nil {
		return nil, errors.New("[User] db is nil")
	}

	if m == nil {
		return nil, errors.New("[User] param is nil")
	}

	m.CreatedAt = godb.Time(time.Now().UTC())

	if r := c.db.Model(c.obj).Create(m); r.Error != nil {
		return nil, errors.WithStack(r.Error)
	} else if r.RowsAffected != 1 {
		return nil, errors.Errorf("[User] add affected row %v", r.RowsAffected)
	}

	return m, nil
}

// GetByKeyId 根据keyId获取User
func (c *beanUser) GetByKeyId(keyId uint) (*User, error) {
	if c.db == nil {
		return nil, errors.New("[User] db is nil")
	}

	var m User
	r := c.db.Model(c.obj).Where(&User{
		Model: godb.Model{
			ID: keyId,
		},
	}).First(&m)
	if r.Error != nil {
		return nil, errors.WithStack(r.Error)
	}
	if m.ID != keyId {
		return nil, errors.Errorf("[User] get fail. param: %v, get: %v", keyId, m.ID)
	}

	return &m, nil
}

// GetUid 根据uid获取User
func (c *beanUser) GetUid(uid string) (*User, error) {
	if c.db == nil {
		return nil, errors.New("[User] db is nil")
	}

	if len(uid) == 0 {
		return nil, errors.New("[User] error param uid")
	}

	var m User
	r := c.db.Model(c.obj).Where(&User{
		Model: godb.Model{
			Uid: godb.String(uid),
		},
	}).First(&m)
	if r.Error != nil {
		return nil, errors.WithStack(r.Error)
	}
	if godb.PointString(m.Uid) != uid {
		return nil, errors.Errorf("[User] get beanBook fail. param: %v, get: %v", uid, m.Uid)
	}

	return &m, nil
}

// UpdateByKeyId 根据keyId修改User
func (c *beanUser) UpdateByKeyId(obj *User) error {
	if c.db == nil {
		return errors.New("[User] db is nil")
	}
	if obj == nil || obj.ID == 0 {
		return errors.New("[User] param error")
	}

	obj.UpdatedAt = godb.Time(time.Now().UTC())

	r := c.db.Model(c.obj).Where(&User{
		Model: godb.Model{
			ID: obj.ID,
		},
	}).Updates(obj)
	if r.Error != nil {
		return errors.WithStack(r.Error)
	}

	return nil
}

// UpdateByUId 根据uid修改User
func (c *beanUser) UpdateByUId(obj *User) error {
	if c.db == nil {
		return errors.New("[User] db is nil")
	}
	if obj == nil || len(godb.PointString(obj.Uid)) == 0 {
		return errors.New("[User] param error")
	}

	obj.UpdatedAt = godb.Time(time.Now().UTC())

	r := c.db.Model(c.obj).Where(&User{
		Model: godb.Model{
			Uid: obj.Uid,
		},
	}).Updates(obj)
	if r.Error != nil {
		return errors.WithStack(r.Error)
	}

	return nil
}

// Update 根据自定义条件更新User
func (c *beanUser) Update(m *User, customFilters ...BeanUserDBOption) (int64, error) {
	if c.db == nil {
		return -1, errors.New("[User] db is nil")
	}

	var (
		err error
		r   = c.db.Model(c.obj)
	)

	var field = newGoUser()
	for _, filter := range customFilters {
		if r, err = filter(r, field); err != nil {
			return -1, err
		}
	}

	r = r.Updates(m)
	if r.Error != nil {
		return -1, errors.WithStack(r.Error)
	}

	return r.RowsAffected, nil
}

// DeleteByKeyId 根据keyId删除User
func (c *beanUser) DeleteByKeyId(keyId uint) error {
	if c.db == nil {
		return errors.New("[User] db is nil")
	}

	if keyId == 0 {
		return errors.New("[User] key id is 0")
	}

	var m User
	r := c.db.Model(c.obj).Where(&User{
		Model: godb.Model{
			ID: keyId,
		},
	}).Delete(&m)
	if r.Error != nil {
		return errors.WithStack(r.Error)
	}

	if r.RowsAffected != 1 {
		return errors.Errorf("[User] delete row affected %v", r.RowsAffected)
	}

	return nil
}

// DeleteByUId 根据uid删除User
func (c *beanUser) DeleteByUId(uid string) error {
	if c.db == nil {
		return errors.New("[User] db is nil")
	}

	if len(uid) == 0 {
		return errors.New("[User] uid error")
	}

	var m User
	r := c.db.Model(c.obj).Where(&User{
		Model: godb.Model{
			Uid: godb.String(uid),
		},
	}).Delete(&m)
	if r.Error != nil {
		return errors.WithStack(r.Error)
	}

	if r.RowsAffected != 1 {
		return errors.Errorf("[User] delete row affected %v", r.RowsAffected)
	}

	return nil
}

// Delete 根据自定义条件删除User
func (c *beanUser) Delete(filter *UserFilter, customFilters ...BeanUserDBOption) (int64, error) {
	if c.db == nil {
		return -1, errors.New("[User] db is nil")
	}

	var (
		err error
		m   User
		r   = c.db.Model(c.obj)
	)

	if filter != nil {
		r, err = c.funFilter(filter, customFilters...)
		if err != nil {
			return -1, err
		}
	}

	var field = newGoUser()
	for _, filterFn := range customFilters {
		if r, err = filterFn(r, field); err != nil {
			return -1, err
		}
	}

	r = r.Delete(&m)
	if r.Error != nil {
		return -1, errors.WithStack(r.Error)
	}

	return r.RowsAffected, nil
}

type BeanUserDBOption func(r *gorm.DB, field *GoUser) (*gorm.DB, error)

type UserFilter struct {
	Limit       int
	Offset      int
	Object      *User            // 完全匹配
	ObjectNeq   *User            // Neq
	Match       *User            // 模糊匹配, 仅string类型字段有效, 推荐使用 Cond.Like
	MatchOr     *User            // 模糊匹配, 仅string类型字段有效, 推荐使用 Cond.Like
	MatchString *string          // 模糊匹配所有string类型字段
	CondList    []UserFilterCond //
}

// Filter 筛选User
func (c *beanUser) Filter(filter *UserFilter, customFilters ...BeanUserDBOption) ([]User, int64, error) {
	if c.db == nil {
		return nil, -1, errors.New("[User] db is nil")
	}

	var limit = 0
	var offset = 0
	if filter != nil {
		limit = filter.Limit
		offset = filter.Offset
	}
	if limit < 0 {
		return nil, -1, errors.New("[User] limit error")
	}
	if offset < 0 {
		return nil, -1, errors.New("[User] offset error")
	}

	var m []User
	r, err := c.funFilter(filter, customFilters...)
	if err != nil {
		return nil, -1, err
	}
	r = r.Offset(offset)
	if limit > 0 {
		r = r.Limit(limit)
	}
	if r.Find(&m); r.Error != nil {
		return nil, -1, errors.WithStack(r.Error)
	}

	var totalCount int64
	{
		r, err := c.funFilter(filter, customFilters...)
		if err != nil {
			return nil, -1, err
		}
		if r := r.Count(&totalCount); r.Error != nil {
			return nil, -1, errors.WithStack(r.Error)
		}
	}

	return m, totalCount, nil
}

// Find 查找User, matchObject: 仅string类型字段有效
func (c *beanUser) Find(filter *UserFilter, customFilters ...BeanUserDBOption) ([]User, error) {
	if c.db == nil {
		return nil, errors.New("[User] db is nil")
	}

	var limit = 0
	var offset = 0
	if filter != nil {
		limit = filter.Limit
		offset = filter.Offset
	}
	if limit < 0 {
		return nil, errors.New("[User] limit error")
	}
	if offset < 0 {
		return nil, errors.New("[User] offset error")
	}

	var m []User
	r, err := c.funFilter(filter, customFilters...)
	if err != nil {
		return nil, err
	}
	r = r.Offset(offset)
	if limit > 0 {
		r = r.Limit(limit)
	}
	if r.Find(&m); r.Error != nil {
		return nil, errors.WithStack(r.Error)
	}

	return m, nil
}

// FindToDB 查找User, matchObject: 仅string类型字段有效
//
//	subQuery, _ := FindToDB(***, func(***) {
//	    r = r.Select(field.FieldName***())
//	})
//
// r = r.Where(field.FieldNameXXXUid() + " IN (?)", subQuery)
func (c *beanUser) FindToDB(filter *UserFilter, customFilters ...BeanUserDBOption) (*gorm.DB, error) {
	if c.db == nil {
		return nil, errors.New("[User] db is nil")
	}

	var limit = 0
	var offset = 0
	if filter != nil {
		limit = filter.Limit
		offset = filter.Offset
	}
	if limit < 0 {
		return nil, errors.New("[User] limit error")
	}
	if offset < 0 {
		return nil, errors.New("[User] offset error")
	}

	r, err := c.funFilter(filter, customFilters...)
	if err != nil {
		return nil, err
	}

	r = r.Offset(offset)
	if limit > 0 {
		r = r.Limit(limit)
	}

	return r, nil
}

// FindAndRange 查找User, matchObject: 仅string类型字段有效
func (c *beanUser) FindAndRange(fun func(*User) bool, filter *UserFilter, customFilters ...BeanUserDBOption) error {
	if c.db == nil {
		return errors.New("[User] db is nil")
	}
	if fun == nil {
		return errors.New("[User] param error")
	}

	var limit = 0
	var offset = 0
	if filter != nil {
		limit = filter.Limit
		offset = filter.Offset
	}
	if limit < 0 {
		return errors.New("[User] limit error")
	}
	if offset < 0 {
		return errors.New("[User] offset error")
	}

	r, err := c.funFilter(filter, customFilters...)
	if err != nil {
		return err
	}
	r = r.Offset(offset)
	if limit > 0 {
		r = r.Limit(limit)
	}

	rows, err := r.Rows()
	if err != nil {
		return errors.WithStack(err)
	}
	defer rows.Close()

	for rows.Next() {
		var object User
		if err := r.ScanRows(rows, &object); err != nil {
			return errors.WithStack(err)
		}

		if fun(&object) {
			continue
		} else {
			break
		}
	}

	return nil
}

// First User, matchObject: 仅string类型字段有效
func (c *beanUser) First(filter *UserFilter, customFilters ...BeanUserDBOption) (*User, error) {
	if c.db == nil {
		return nil, errors.New("[User] db is nil")
	}

	var offset = 0
	if filter != nil {
		offset = filter.Offset
	}
	if offset < 0 {
		return nil, errors.New("[User] offset error")
	}

	var m User
	r, err := c.funFilter(filter, customFilters...)
	if err != nil {
		return nil, err
	}
	if r := r.Offset(offset).First(&m); r.Error != nil {
		return nil, errors.WithStack(r.Error)
	}

	return &m, nil
}

// Last User, matchObject: 仅string类型字段有效
func (c *beanUser) Last(filter *UserFilter, customFilters ...BeanUserDBOption) (*User, error) {
	if c.db == nil {
		return nil, errors.New("[User] db is nil")
	}

	var offset = 0
	if filter != nil {
		offset = filter.Offset
	}
	if offset < 0 {
		return nil, errors.New("[User] offset error")
	}

	var m User
	r, err := c.funFilter(filter, customFilters...)
	if err != nil {
		return nil, err
	}
	if r := r.Offset(offset).Last(&m); r.Error != nil {
		return nil, errors.WithStack(r.Error)
	}

	return &m, nil
}

// Has 判断是否存在 User
func (c *beanUser) Has(filter *UserFilter, customFilters ...BeanUserDBOption) (bool, error) {
	if c.db == nil {
		return false, errors.New("[User] db is nil")
	}

	var totalCount int64
	r, err := c.funFilter(filter, customFilters...)
	if err != nil {
		return false, err
	}
	if r := r.Limit(1).Count(&totalCount); r.Error != nil {
		return false, errors.WithStack(r.Error)
	}

	if totalCount > 0 {
		return true, nil
	} else {
		return false, nil
	}
}

// Count 筛选数量 User
func (c *beanUser) Count(filter *UserFilter, customFilters ...BeanUserDBOption) (int64, error) {
	if c.db == nil {
		return -1, errors.New("[User] db is nil")
	}

	var totalCount int64
	r, err := c.funFilter(filter, customFilters...)
	if err != nil {
		return 0, err
	}
	if r := r.Count(&totalCount); r.Error != nil {
		return -1, errors.WithStack(r.Error)
	}

	return totalCount, nil
}

func (c *beanUser) funFilter(filter *UserFilter, customFilters ...BeanUserDBOption) (*gorm.DB, error) {

	var (
		err   error
		r     = c.db.Model(c.obj)
		field = newGoUser()
	)

	if filter != nil {
		if filter.Object != nil {
			r = r.Where(filter.Object)
		}

		var objectNeq = filter.ObjectNeq
		if objectNeq != nil {

			if objectNeq.ID != 0 {
				r = r.Clauses(clause.Neq{
					Column: field.FieldNameID(),
					Value:  objectNeq.ID,
				})
			}

			if objectNeq.Uid != nil {
				r = r.Clauses(clause.Neq{
					Column: field.FieldNameUid(),
					Value:  objectNeq.Uid,
				})
			}

			if objectNeq.CreatedAt != nil {
				r = r.Clauses(clause.Neq{
					Column: field.FieldNameCreatedAt(),
					Value:  objectNeq.CreatedAt,
				})
			}

			if objectNeq.UpdatedAt != nil {
				r = r.Clauses(clause.Neq{
					Column: field.FieldNameUpdatedAt(),
					Value:  objectNeq.UpdatedAt,
				})
			}

			if objectNeq.Name != nil {
				r = r.Clauses(clause.Neq{
					Column: field.FieldNameName(),
					Value:  objectNeq.Name,
				})
			}

			if objectNeq.Age != nil {
				r = r.Clauses(clause.Neq{
					Column: field.FieldNameAge(),
					Value:  objectNeq.Age,
				})
			}

			if objectNeq.Birthday != nil {
				r = r.Clauses(clause.Neq{
					Column: field.FieldNameBirthday(),
					Value:  objectNeq.Birthday,
				})
			}

			if objectNeq.CompanyID != nil {
				r = r.Clauses(clause.Neq{
					Column: field.FieldNameCompanyID(),
					Value:  objectNeq.CompanyID,
				})
			}

			if objectNeq.ManagerID != nil {
				r = r.Clauses(clause.Neq{
					Column: field.FieldNameManagerID(),
					Value:  objectNeq.ManagerID,
				})
			}

			if objectNeq.A != nil {
				r = r.Clauses(clause.Neq{
					Column: field.FieldNameA(),
					Value:  objectNeq.A,
				})
			}

			if objectNeq.B != nil {
				r = r.Clauses(clause.Neq{
					Column: field.FieldNameB(),
					Value:  objectNeq.B,
				})
			}

		}

		var match = filter.Match
		if match != nil {

			if match.Uid != nil {
				r = r.Clauses(clause.Like{
					Column: field.FieldNameUid(),
					Value:  "%" + godb.PointString(match.Uid) + "%",
				})
			}

			if match.Name != nil {
				r = r.Clauses(clause.Like{
					Column: field.FieldNameName(),
					Value:  "%" + godb.PointString(match.Name) + "%",
				})
			}

		}

		var matchOr = filter.MatchOr
		if matchOr != nil {
			var conds []clause.Expression

			if matchOr.Uid != nil {
				conds = append(conds, clause.Like{
					Column: field.FieldNameUid(),
					Value:  "%" + godb.PointString(matchOr.Uid) + "%",
				})
			}

			if matchOr.Name != nil {
				conds = append(conds, clause.Like{
					Column: field.FieldNameName(),
					Value:  "%" + godb.PointString(matchOr.Name) + "%",
				})
			}

			if len(conds) > 0 {
				r = r.Clauses(clause.Or(conds...))
			}
		}

		var matchStr = godb.PointString(filter.MatchString)
		if len(matchStr) > 0 {
			var exps []clause.Expression

			exps = append(exps, clause.Like{
				Column: field.FieldNameUid(),
				Value:  "%" + matchStr + "%",
			})

			exps = append(exps, clause.Like{
				Column: field.FieldNameName(),
				Value:  "%" + matchStr + "%",
			})

			if len(exps) > 0 {
				e := clause.Or(exps...)
				r.Clauses(e)
			}
		}

		for _, cond := range filter.CondList {
			switch cond.Type {
			case CondTypeGt:
				if cond.Number != nil {

					if cond.Number.ID != nil {
						r = r.Clauses(clause.Gt{
							Column: field.FieldNameID(),
							Value:  cond.Number.ID,
						})
					}

					if cond.Number.Age != nil {
						r = r.Clauses(clause.Gt{
							Column: field.FieldNameAge(),
							Value:  cond.Number.Age,
						})
					}

					if cond.Number.CompanyID != nil {
						r = r.Clauses(clause.Gt{
							Column: field.FieldNameCompanyID(),
							Value:  cond.Number.CompanyID,
						})
					}

					if cond.Number.ManagerID != nil {
						r = r.Clauses(clause.Gt{
							Column: field.FieldNameManagerID(),
							Value:  cond.Number.ManagerID,
						})
					}

					if cond.Number.A != nil {
						r = r.Clauses(clause.Gt{
							Column: field.FieldNameA(),
							Value:  cond.Number.A,
						})
					}

					if cond.Number.B != nil {
						r = r.Clauses(clause.Gt{
							Column: field.FieldNameB(),
							Value:  cond.Number.B,
						})
					}

				}
			case CondTypeGte: // 大于等于, int
				if cond.Number != nil {

					if cond.Number.ID != nil {
						r = r.Clauses(clause.Gte{
							Column: field.FieldNameID(),
							Value:  cond.Number.ID,
						})
					}

					if cond.Number.Age != nil {
						r = r.Clauses(clause.Gte{
							Column: field.FieldNameAge(),
							Value:  cond.Number.Age,
						})
					}

					if cond.Number.CompanyID != nil {
						r = r.Clauses(clause.Gte{
							Column: field.FieldNameCompanyID(),
							Value:  cond.Number.CompanyID,
						})
					}

					if cond.Number.ManagerID != nil {
						r = r.Clauses(clause.Gte{
							Column: field.FieldNameManagerID(),
							Value:  cond.Number.ManagerID,
						})
					}

					if cond.Number.A != nil {
						r = r.Clauses(clause.Gte{
							Column: field.FieldNameA(),
							Value:  cond.Number.A,
						})
					}

					if cond.Number.B != nil {
						r = r.Clauses(clause.Gte{
							Column: field.FieldNameB(),
							Value:  cond.Number.B,
						})
					}

				}
			case CondTypeLt: // 小于, int
				if cond.Number != nil {

					if cond.Number.ID != nil {
						r = r.Clauses(clause.Lt{
							Column: field.FieldNameID(),
							Value:  cond.Number.ID,
						})
					}

					if cond.Number.Age != nil {
						r = r.Clauses(clause.Lt{
							Column: field.FieldNameAge(),
							Value:  cond.Number.Age,
						})
					}

					if cond.Number.CompanyID != nil {
						r = r.Clauses(clause.Lt{
							Column: field.FieldNameCompanyID(),
							Value:  cond.Number.CompanyID,
						})
					}

					if cond.Number.ManagerID != nil {
						r = r.Clauses(clause.Lt{
							Column: field.FieldNameManagerID(),
							Value:  cond.Number.ManagerID,
						})
					}

					if cond.Number.A != nil {
						r = r.Clauses(clause.Lt{
							Column: field.FieldNameA(),
							Value:  cond.Number.A,
						})
					}

					if cond.Number.B != nil {
						r = r.Clauses(clause.Lt{
							Column: field.FieldNameB(),
							Value:  cond.Number.B,
						})
					}

				}
			case CondTypeLte: // 小于等于, int
				if cond.Number != nil {

					if cond.Number.ID != nil {
						r = r.Clauses(clause.Lte{
							Column: field.FieldNameID(),
							Value:  cond.Number.ID,
						})
					}

					if cond.Number.Age != nil {
						r = r.Clauses(clause.Lte{
							Column: field.FieldNameAge(),
							Value:  cond.Number.Age,
						})
					}

					if cond.Number.CompanyID != nil {
						r = r.Clauses(clause.Lte{
							Column: field.FieldNameCompanyID(),
							Value:  cond.Number.CompanyID,
						})
					}

					if cond.Number.ManagerID != nil {
						r = r.Clauses(clause.Lte{
							Column: field.FieldNameManagerID(),
							Value:  cond.Number.ManagerID,
						})
					}

					if cond.Number.A != nil {
						r = r.Clauses(clause.Lte{
							Column: field.FieldNameA(),
							Value:  cond.Number.A,
						})
					}

					if cond.Number.B != nil {
						r = r.Clauses(clause.Lte{
							Column: field.FieldNameB(),
							Value:  cond.Number.B,
						})
					}

				}
			case CondTypeBetween: // 区间, int
				if cond.Between != nil {

					if cond.Between.IDLessVal != nil {
						r = r.Clauses(clause.Gte{
							Column: field.FieldNameID(),
							Value:  cond.Between.IDLessVal,
						})
					}

					if cond.Between.IDMoreVal != nil {
						r = r.Clauses(clause.Lte{
							Column: field.FieldNameID(),
							Value:  cond.Between.IDMoreVal,
						})
					}

					if cond.Between.AgeLessVal != nil {
						r = r.Clauses(clause.Gte{
							Column: field.FieldNameAge(),
							Value:  cond.Between.AgeLessVal,
						})
					}

					if cond.Between.AgeMoreVal != nil {
						r = r.Clauses(clause.Lte{
							Column: field.FieldNameAge(),
							Value:  cond.Between.AgeMoreVal,
						})
					}

					if cond.Between.CompanyIDLessVal != nil {
						r = r.Clauses(clause.Gte{
							Column: field.FieldNameCompanyID(),
							Value:  cond.Between.CompanyIDLessVal,
						})
					}

					if cond.Between.CompanyIDMoreVal != nil {
						r = r.Clauses(clause.Lte{
							Column: field.FieldNameCompanyID(),
							Value:  cond.Between.CompanyIDMoreVal,
						})
					}

					if cond.Between.ManagerIDLessVal != nil {
						r = r.Clauses(clause.Gte{
							Column: field.FieldNameManagerID(),
							Value:  cond.Between.ManagerIDLessVal,
						})
					}

					if cond.Between.ManagerIDMoreVal != nil {
						r = r.Clauses(clause.Lte{
							Column: field.FieldNameManagerID(),
							Value:  cond.Between.ManagerIDMoreVal,
						})
					}

					if cond.Between.ALessVal != nil {
						r = r.Clauses(clause.Gte{
							Column: field.FieldNameA(),
							Value:  cond.Between.ALessVal,
						})
					}

					if cond.Between.AMoreVal != nil {
						r = r.Clauses(clause.Lte{
							Column: field.FieldNameA(),
							Value:  cond.Between.AMoreVal,
						})
					}

					if cond.Between.BLessVal != nil {
						r = r.Clauses(clause.Gte{
							Column: field.FieldNameB(),
							Value:  cond.Between.BLessVal,
						})
					}

					if cond.Between.BMoreVal != nil {
						r = r.Clauses(clause.Lte{
							Column: field.FieldNameB(),
							Value:  cond.Between.BMoreVal,
						})
					}

				}
			case CondTypeIn: // in, 集合
				if cond.Container != nil {

					if len(cond.Container.IDList) > 0 {
						var values = make([]interface{}, 0, len(cond.Container.IDList))
						for _, value := range cond.Container.IDList {
							values = append(values, value)
						}

						r = r.Clauses(clause.IN{
							Column: field.FieldNameID(),
							Values: values,
						})
					}

					if len(cond.Container.UidList) > 0 {
						var values = make([]interface{}, 0, len(cond.Container.UidList))
						for _, value := range cond.Container.UidList {
							values = append(values, value)
						}

						r = r.Clauses(clause.IN{
							Column: field.FieldNameUid(),
							Values: values,
						})
					}

					if len(cond.Container.CreatedAtList) > 0 {
						var values = make([]interface{}, 0, len(cond.Container.CreatedAtList))
						for _, value := range cond.Container.CreatedAtList {
							values = append(values, value)
						}

						r = r.Clauses(clause.IN{
							Column: field.FieldNameCreatedAt(),
							Values: values,
						})
					}

					if len(cond.Container.UpdatedAtList) > 0 {
						var values = make([]interface{}, 0, len(cond.Container.UpdatedAtList))
						for _, value := range cond.Container.UpdatedAtList {
							values = append(values, value)
						}

						r = r.Clauses(clause.IN{
							Column: field.FieldNameUpdatedAt(),
							Values: values,
						})
					}

					if len(cond.Container.NameList) > 0 {
						var values = make([]interface{}, 0, len(cond.Container.NameList))
						for _, value := range cond.Container.NameList {
							values = append(values, value)
						}

						r = r.Clauses(clause.IN{
							Column: field.FieldNameName(),
							Values: values,
						})
					}

					if len(cond.Container.AgeList) > 0 {
						var values = make([]interface{}, 0, len(cond.Container.AgeList))
						for _, value := range cond.Container.AgeList {
							values = append(values, value)
						}

						r = r.Clauses(clause.IN{
							Column: field.FieldNameAge(),
							Values: values,
						})
					}

					if len(cond.Container.BirthdayList) > 0 {
						var values = make([]interface{}, 0, len(cond.Container.BirthdayList))
						for _, value := range cond.Container.BirthdayList {
							values = append(values, value)
						}

						r = r.Clauses(clause.IN{
							Column: field.FieldNameBirthday(),
							Values: values,
						})
					}

					if len(cond.Container.CompanyIDList) > 0 {
						var values = make([]interface{}, 0, len(cond.Container.CompanyIDList))
						for _, value := range cond.Container.CompanyIDList {
							values = append(values, value)
						}

						r = r.Clauses(clause.IN{
							Column: field.FieldNameCompanyID(),
							Values: values,
						})
					}

					if len(cond.Container.ManagerIDList) > 0 {
						var values = make([]interface{}, 0, len(cond.Container.ManagerIDList))
						for _, value := range cond.Container.ManagerIDList {
							values = append(values, value)
						}

						r = r.Clauses(clause.IN{
							Column: field.FieldNameManagerID(),
							Values: values,
						})
					}

					if len(cond.Container.AList) > 0 {
						var values = make([]interface{}, 0, len(cond.Container.AList))
						for _, value := range cond.Container.AList {
							values = append(values, value)
						}

						r = r.Clauses(clause.IN{
							Column: field.FieldNameA(),
							Values: values,
						})
					}

					if len(cond.Container.BList) > 0 {
						var values = make([]interface{}, 0, len(cond.Container.BList))
						for _, value := range cond.Container.BList {
							values = append(values, value)
						}

						r = r.Clauses(clause.IN{
							Column: field.FieldNameB(),
							Values: values,
						})
					}

				}
			case CondTypeNotIn: // not in, 集合
				if cond.Container != nil {

					if len(cond.Container.IDList) > 0 {
						var values = make([]interface{}, 0, len(cond.Container.IDList))
						for _, value := range cond.Container.IDList {
							values = append(values, value)
						}

						r = r.Clauses(clause.Not(
							clause.IN{
								Column: field.FieldNameID(),
								Values: values,
							},
						))
					}

					if len(cond.Container.UidList) > 0 {
						var values = make([]interface{}, 0, len(cond.Container.UidList))
						for _, value := range cond.Container.UidList {
							values = append(values, value)
						}

						r = r.Clauses(clause.Not(
							clause.IN{
								Column: field.FieldNameUid(),
								Values: values,
							},
						))
					}

					if len(cond.Container.CreatedAtList) > 0 {
						var values = make([]interface{}, 0, len(cond.Container.CreatedAtList))
						for _, value := range cond.Container.CreatedAtList {
							values = append(values, value)
						}

						r = r.Clauses(clause.Not(
							clause.IN{
								Column: field.FieldNameCreatedAt(),
								Values: values,
							},
						))
					}

					if len(cond.Container.UpdatedAtList) > 0 {
						var values = make([]interface{}, 0, len(cond.Container.UpdatedAtList))
						for _, value := range cond.Container.UpdatedAtList {
							values = append(values, value)
						}

						r = r.Clauses(clause.Not(
							clause.IN{
								Column: field.FieldNameUpdatedAt(),
								Values: values,
							},
						))
					}

					if len(cond.Container.NameList) > 0 {
						var values = make([]interface{}, 0, len(cond.Container.NameList))
						for _, value := range cond.Container.NameList {
							values = append(values, value)
						}

						r = r.Clauses(clause.Not(
							clause.IN{
								Column: field.FieldNameName(),
								Values: values,
							},
						))
					}

					if len(cond.Container.AgeList) > 0 {
						var values = make([]interface{}, 0, len(cond.Container.AgeList))
						for _, value := range cond.Container.AgeList {
							values = append(values, value)
						}

						r = r.Clauses(clause.Not(
							clause.IN{
								Column: field.FieldNameAge(),
								Values: values,
							},
						))
					}

					if len(cond.Container.BirthdayList) > 0 {
						var values = make([]interface{}, 0, len(cond.Container.BirthdayList))
						for _, value := range cond.Container.BirthdayList {
							values = append(values, value)
						}

						r = r.Clauses(clause.Not(
							clause.IN{
								Column: field.FieldNameBirthday(),
								Values: values,
							},
						))
					}

					if len(cond.Container.CompanyIDList) > 0 {
						var values = make([]interface{}, 0, len(cond.Container.CompanyIDList))
						for _, value := range cond.Container.CompanyIDList {
							values = append(values, value)
						}

						r = r.Clauses(clause.Not(
							clause.IN{
								Column: field.FieldNameCompanyID(),
								Values: values,
							},
						))
					}

					if len(cond.Container.ManagerIDList) > 0 {
						var values = make([]interface{}, 0, len(cond.Container.ManagerIDList))
						for _, value := range cond.Container.ManagerIDList {
							values = append(values, value)
						}

						r = r.Clauses(clause.Not(
							clause.IN{
								Column: field.FieldNameManagerID(),
								Values: values,
							},
						))
					}

					if len(cond.Container.AList) > 0 {
						var values = make([]interface{}, 0, len(cond.Container.AList))
						for _, value := range cond.Container.AList {
							values = append(values, value)
						}

						r = r.Clauses(clause.Not(
							clause.IN{
								Column: field.FieldNameA(),
								Values: values,
							},
						))
					}

					if len(cond.Container.BList) > 0 {
						var values = make([]interface{}, 0, len(cond.Container.BList))
						for _, value := range cond.Container.BList {
							values = append(values, value)
						}

						r = r.Clauses(clause.Not(
							clause.IN{
								Column: field.FieldNameB(),
								Values: values,
							},
						))
					}

				}
			case CondTypeLike: // like, string
				if cond.Like != nil {

					if cond.Like.Uid != nil {
						r = r.Clauses(clause.Like{
							Column: field.FieldNameUid(),
							Value:  "%" + godb.PointString(cond.Like.Uid) + "%",
						})
					}

					if cond.Like.Name != nil {
						r = r.Clauses(clause.Like{
							Column: field.FieldNameName(),
							Value:  "%" + godb.PointString(cond.Like.Name) + "%",
						})
					}

				}
			}
		}

	}

	if len(customFilters) > 0 {
		for _, filter := range customFilters {
			if filter == nil {
				continue
			}

			r, err = filter(r, field)
			if err != nil {
				return r, err
			}
		}
	}

	return r, err
}

// UserFilterCond cond
type UserFilterCond struct {
	Type      CondType
	Number    *UserCondNumber
	Between   *UserCondBetween
	Container *UserCondContainer
	Like      *UserLike
}

type UserCondNumber struct {
	ID *uint // ID,

	Age *int // Age,

	CompanyID *uint // CompanyID,

	ManagerID *uint // ManagerID,

	A *float64 // A,

	B *float64 // B,

}

type UserCondBetween struct {
	IDLessVal *uint // ID less value,
	IDMoreVal *uint // ID more value,

	AgeLessVal *int // Age less value,
	AgeMoreVal *int // Age more value,

	CompanyIDLessVal *uint // CompanyID less value,
	CompanyIDMoreVal *uint // CompanyID more value,

	ManagerIDLessVal *uint // ManagerID less value,
	ManagerIDMoreVal *uint // ManagerID more value,

	ALessVal *float64 // A less value,
	AMoreVal *float64 // A more value,

	BLessVal *float64 // B less value,
	BMoreVal *float64 // B more value,

}

type UserCondContainer struct {
	IDList []uint // ID slice,

	UidList []string // Uid slice,

	CreatedAtList []time.Time // CreatedAt slice,

	UpdatedAtList []time.Time // UpdatedAt slice,

	NameList []string // Name slice,

	AgeList []int // Age slice

	BirthdayList []time.Time // Birthday slice,

	CompanyIDList []uint // CompanyID slice,

	ManagerIDList []uint // ManagerID slice,

	AList []float64 // A slice,

	BList []float64 // B slice,

}

type UserLike struct {
	Uid *string // Uid,

	Name *string // Name,

}
