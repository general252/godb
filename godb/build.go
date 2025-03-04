package godb

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
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

	structs, _ := parseStructComments(filepath.Join(dir, "model.go"))

	db, err := ParseTables(beans, structs)
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

type JsonStructInfo struct {
	Name   string                          // struct name
	Fields map[string]*JsonStructFieldInfo // map[fieldName]*JsonStructFieldInfo
}

type JsonStructFieldInfo struct {
	Name    string
	Type    string
	Tag     string
	Comment string
}

// 解析 Go 源文件，提取 struct 字段的行尾注释
func parseStructComments(filename string) (map[string]*JsonStructInfo, error) {
	// 打开源代码文件
	src, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	// 创建 Go 解析器
	fs := token.NewFileSet()
	node, err := parser.ParseFile(fs, filename, src, parser.AllErrors|parser.ParseComments)
	if err != nil {
		return nil, err
	}

	// 存储字段注释
	var results = map[string]*JsonStructInfo{}

	// 遍历 AST 语法树
	ast.Inspect(node, func(n ast.Node) bool {
		// 只处理 struct 类型
		ts, ok := n.(*ast.TypeSpec)
		if !ok {
			return true
		}

		// 确保是 struct 定义
		st, ok := ts.Type.(*ast.StructType)
		if !ok {
			return true
		}

		result := &JsonStructInfo{
			Name:   ts.Name.Name,
			Fields: map[string]*JsonStructFieldInfo{},
		}

		// 遍历 struct 字段
		for _, field := range st.Fields.List {
			// 获取字段名称
			var fieldName = "(匿名字段)"
			if len(field.Names) > 0 {
				fieldName = field.Names[0].Name
			}

			// 获取字段类型
			fieldType := getFieldType(field.Type)

			// 获取 struct tag
			var tag string
			if field.Tag != nil {
				tag = strings.Trim(field.Tag.Value, "`") // 去掉 ` 号
			}

			// 获取字段注释
			comment := ""
			if field.Comment != nil {
				comment = strings.TrimSpace(field.Comment.Text())
			}

			result.Fields[fieldName] = &JsonStructFieldInfo{
				Name:    fieldName,
				Type:    fieldType,
				Tag:     tag,
				Comment: comment,
			}
		}

		results[result.Name] = result
		return false
	})

	return results, nil
}

// 解析字段类型
func getFieldType(expr ast.Expr) string {
	switch t := expr.(type) {
	case *ast.Ident: // 普通类型 (int, string, 自定义类型)
		return t.Name
	case *ast.StarExpr: // 指针类型 (*User)
		return "*" + getFieldType(t.X)
	case *ast.SelectorExpr: // 选择器类型 (如 time.Time)
		return getFieldType(t.X) + "." + t.Sel.Name
	case *ast.ArrayType: // 数组类型 ([]int, []*User)
		return "[]" + getFieldType(t.Elt)
	case *ast.MapType: // map 类型 (map[string]int)
		return "map[" + getFieldType(t.Key) + "]" + getFieldType(t.Value)
	case *ast.StructType: // 匿名结构体 (struct {...})
		return "struct { ... }"
	default:
		return fmt.Sprintf("%T", t) // 未处理的类型
	}
}
