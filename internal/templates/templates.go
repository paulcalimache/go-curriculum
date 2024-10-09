package templates

import (
	"bytes"
	"embed"
	"fmt"
	"html/template"
	"slices"
)

//go:embed *
var TemplatesFiles embed.FS

var templatesList = []string{
	"classic",
}

func GetTemplatesList() string {
	return fmt.Sprintf("Templates list : %v\n", templatesList)
}

func Templetize(templateName string, data any) (bytes.Buffer, error) {
	var file bytes.Buffer
	if !slices.Contains(templatesList, templateName) {
		return file, fmt.Errorf("template %s doesn't exist. Here are the existing templates : %v", templateName, templatesList)
	}

	t, err := template.New(templateName+".html").ParseFS(TemplatesFiles, fmt.Sprintf("%s/*.html", templateName))
	if err != nil {
		return file, err
	}

	err = t.ExecuteTemplate(&file, templateName+".html", data)
	if err != nil {
		return file, err
	}
	return file, nil
}
