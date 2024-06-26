package godb

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/general252/godb/template"
)

// BuildDatabaseTable 根据bean构建数据访问接口
func BuildDatabaseTable(beans []interface{}) error {
	_, fullPath, _, ok := runtime.Caller(1)
	if !ok {
		return fmt.Errorf("get file fail")
	}

	dir, file, _ := Split(fullPath)
	modelFieldFilename := filepath.ToSlash(filepath.Clean(fmt.Sprintf("%v/%v_field.go", dir, file)))
	engineFilename := filepath.ToSlash(filepath.Clean(fmt.Sprintf("%v/engine.go", dir)))

	db, err := ParseTables(beans)
	if err != nil {
		return err
	}

	// 生成结构体字段models_field.go
	if err = BuildFileFromTemplate(template.GoGoModelsField(), modelFieldFilename, db); err != nil {
		return err
	}

	if err = os.WriteFile(engineFilename, []byte(template.GoGoEngine()), os.ModePerm); err != nil {
		return err
	} else {
		//if err = formatGoFile(engineFilename); err != nil {
		//	return err
		//}
	}

	// 生成db help bean
	if true {
		filename := filepath.ToSlash(filepath.Clean(fmt.Sprintf("%v/model_beans.go", dir)))
		if err = BuildFileFromTemplate(template.GoGoBean(), filename, db); err != nil {
			return err
		}
	} else {
		for _, table := range db.Tables {
			filename := filepath.ToSlash(filepath.Clean(fmt.Sprintf("%v/bean_%v.go", dir, table.GoStructName)))
			if err = BuildFileFromTemplate(template.GoGoBean(), filename, table); err != nil {
				return err
			}
		}
	}

	return nil
}

// Split 分割路径
func Split(fullPath string) (dir, file, ext string) {
	var tmpFile string

	dir, tmpFile = filepath.Split(fullPath)
	ext = filepath.Ext(tmpFile)

	index := strings.LastIndex(tmpFile, ext)
	if index >= 0 {
		file = tmpFile[:index]
	} else {
		file = tmpFile
	}

	return
}
