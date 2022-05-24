package godb

import (
	"bytes"
	"io/ioutil"
	"os"
	"text/template"

	"github.com/fsgo/go_fmt/gofmtapi"
)

// BuildFileFromTemplate 根据模板生成文件
func BuildFileFromTemplate(templateString string, outFilePath string, data interface{}) error {

	tmpl, err := template.New("test").Parse(templateString)
	if err != nil {
		return err
	}

	var buffer bytes.Buffer
	if err = tmpl.Execute(&buffer, data); err != nil {
		return err
	}

	if err = ioutil.WriteFile(outFilePath, buffer.Bytes(), os.ModePerm); err != nil {
		return err
	}

	if err = formatGoFile(outFilePath); err != nil {
		return err
	}

	return nil
}

func formatGoFile(fileFullPath string) error {
	gf := gofmtapi.NewFormatter()
	opt := gofmtapi.NewOptions()
	opt.Files = []string{fileFullPath}

	if err := gf.Execute(opt); err != nil {
		return err
	}

	return nil
}
