package tool

import (
	"go/ast"
	"go/parser"
	"go/token"
)

// getStructNames 获取文件中的结构体名称
func getStructNames(filename string) ([]string, error) {
	var structNameList = make([]string, 0, 50)
	fSet := token.NewFileSet()

	astFile, err := parser.ParseFile(fSet, filename, nil, parser.AllErrors)
	if err != nil {
		return nil, err
	}

	ast.Inspect(astFile, func(node ast.Node) bool {
		sp, ok := node.(*ast.TypeSpec)
		if ok {
			structNameList = append(structNameList, sp.Name.Name)
		}

		return true
	})

	return structNameList, nil
}
