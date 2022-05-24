package godb

import (
	"bytes"
	"github.com/fsgo/go_fmt/gofmtapi"
	"io/ioutil"
	"os"
	"text/template"
)

// BuildFileFromTemplate 根据模板生成文件
func BuildFileFromTemplate(templateString string, outFilePath string, data interface{}) error {

	tmpl, err := template.New("test").Parse(templateString)
	if err != nil {
		return err
	}

	var buffer bytes.Buffer
	if err := tmpl.Execute(&buffer, data); err != nil {
		return err
	}

	if err := ioutil.WriteFile(outFilePath, buffer.Bytes(), os.ModePerm); err != nil {
		return err
	}

	gf := gofmtapi.NewFormatter()
	opt := gofmtapi.NewOptions()
	opt.Files = []string{outFilePath}

	if err = gf.Execute(opt); err != nil {
		return err
	}

	return nil
}
