package godb

import (
	"fmt"
	"github.com/general252/godb/template"
	"github.com/general252/gout/ufile"
	"io/ioutil"
	"runtime"
)

// BuildDatabaseTable 根据bean构建数据访问接口
func BuildDatabaseTable(beans []interface{}) error {
	_, fullPath, _, ok := runtime.Caller(1)
	if !ok {
		return fmt.Errorf("get file fail")
	}

	dir, file, _ := ufile.Split(fullPath)
	modelFieldFilename := fmt.Sprintf("%v/%v_field.go", dir, file)
	typePointFilename := fmt.Sprintf("%v/type_point.go", dir)
	engineFilename := fmt.Sprintf("%v/engine.go", dir)

	db, err := ParseTables(beans)
	if err != nil {
		return err
	}

	// 生成结构体字段models_field.go
	if err := BuildFileFromTemplate(template.GoGoModelsField, modelFieldFilename, db); err != nil {
		return err
	}

	// 生成type_point.go
	if err := ioutil.WriteFile(typePointFilename, []byte(template.GoGoTypePoint), 0666); err != nil {
		return err
	}

	if err := ioutil.WriteFile(engineFilename, []byte(template.GoGoEngine), 0666); err != nil {
		return err
	}

	// 生成db help bean
	for _, table := range db.Tables {
		filename := fmt.Sprintf("%v/bean_%v.go", dir, table.GoStructName)
		if err := BuildFileFromTemplate(template.GoGoBean, filename, table); err != nil {
			return err
		}
	}

	return nil
}
