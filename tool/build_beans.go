package tool

import (
	"bytes"
	"io/ioutil"
	"os"
	"text/template"
	"time"
)

// CreateBean 创建beans
func CreateBean(modelFilename string, templateString string, outputPath string) error {
	structNames, err := getStructNames(modelFilename)
	if err != nil {
		return err
	}

	type Model struct {
		Name string
	}
	type Data struct {
		CreateTime string
		Models     []Model
	}

	var data = &Data{
		CreateTime: time.Now().Format(time.RFC3339),
		Models:     nil,
	}
	for _, structName := range structNames {
		data.Models = append(data.Models, Model{Name: structName})
	}

	if err := BuildFileFromTemplate(templateString, outputPath, data); err != nil {
		return err
	}

	return nil
}

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

	return nil
}
