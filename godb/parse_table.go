package godb

import (
	"sync"

	"gorm.io/gorm/schema"
)

// Field 字段
type Field struct {
	GoFieldName  string            // go字段名
	DBColumnName string            // 表字段名
	GoType       string            // go字段类型
	DBType       string            //表字段类型
	Index        int               // 第n个字段(1,2,3...)
	Tag          map[string]string // gorm tag
	TagString    string            // gorm tag string
	Size         int               // 表字段长度
	Comment      string            // Comment
	Parent       *Table            `json:"-"` // 表
}

// Table 表
type Table struct {
	GoStructName string  // 结构体名称
	DBTableName  string  // 数据库表名
	Fields       []Field // 字段集合
	FieldsNoID   []Field // 字段集合
}

// DB 数据库
type DB struct {
	Tables []Table // 表集合
}

// ParseTables 根据bean获取table信息, 使用gorm schema解析
func ParseTables(modelBeans []interface{}) (*DB, error) {
	var db DB

	for _, objectBean := range modelBeans {
		objectBeanSchema, err := schema.Parse(objectBean, &sync.Map{}, schema.NamingStrategy{})
		if err != nil {
			return nil, err
		}

		var tab = Table{
			GoStructName: objectBeanSchema.Name,
			DBTableName:  objectBeanSchema.Table,
		}

		var index = 0
		for _, objectField := range objectBeanSchema.Fields {
			index++
			if objectField.Name != "ID" {
				tab.FieldsNoID = append(tab.FieldsNoID, Field{
					GoFieldName:  objectField.Name,
					DBColumnName: objectField.DBName,
					GoType:       string(objectField.GORMDataType),
					DBType:       string(objectField.DataType),
					Index:        index,
					Tag:          objectField.TagSettings,
					TagString:    string(objectField.Tag),
					Comment:      objectField.Comment,
					Size:         objectField.Size,
					Parent:       &tab,
				})
			}
			tab.Fields = append(tab.Fields, Field{
				GoFieldName:  objectField.Name,
				DBColumnName: objectField.DBName,
				GoType:       string(objectField.GORMDataType),
				DBType:       string(objectField.DataType),
				Index:        index,
				Tag:          objectField.TagSettings,
				TagString:    string(objectField.Tag),
				Comment:      objectField.Comment,
				Size:         objectField.Size,
				Parent:       &tab,
			})
		}

		db.Tables = append(db.Tables, tab)
	}

	return &db, nil
}
