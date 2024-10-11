package curriculum

import (
	"bytes"
	"errors"
	"fmt"
	"io/fs"
	"path"
	"slices"
	"strings"
	"text/template"

	"github.com/paulcalimache/go-curriculum/internal/templates"
	"gopkg.in/yaml.v3"
)

type CV struct {
	Firstname   string        `yaml:"firstname"`
	Lastname    string        `yaml:"lastname"`
	Job         string        `yaml:"job"`
	Description string        `yaml:"description"`
	Image       string        `yaml:"image"`
	Contact     Contact       `yaml:"contact"`
	Education   []Education   `yaml:"education"`
	Experiences []Experiences `yaml:"experiences"`
	Skills      []string      `yaml:"skills"`
	Hobbies     []string      `yaml:"hobbies"`
	Projects    []Projects    `yaml:"projects"`
}

type Education struct {
	Timerange   string `yaml:"timerange"`
	Title       string `yaml:"title"`
	Institution string `yaml:"institution"`
}

type Experiences struct {
	Timerange   string `yaml:"timerange"`
	Title       string `yaml:"title"`
	Institution string `yaml:"institution"`
	Description string `yaml:"description"`
}

type Projects struct {
	Title       string `yaml:"title"`
	Description string `yaml:"description"`
	Link        string `yaml:"link"`
}

type Contact struct {
	Mail     string `yaml:"mail"`
	Phone    string `yaml:"phone"`
	Linkedin string `yaml:"linkedin"`
	Website  string `yaml:"website"`
}

func Parse(fileSystem fs.FS, filePath string) (*CV, error) {
	if !strings.Contains(filePath, ".yaml") && !strings.Contains(filePath, ".yml") {
		return nil, errors.New(path.Ext(filePath) + " is not a valid file extension")
	}
	file, err := fs.ReadFile(fileSystem, filePath)
	if err != nil {
		return nil, err
	}
	var cv CV
	err = yaml.Unmarshal(file, &cv)
	return &cv, err
}

func (cv *CV) Templetize(tmplName string) (bytes.Buffer, error) {
	var file bytes.Buffer
	if !slices.Contains(templates.GetTemplatesList(), tmplName) {
		return file, fmt.Errorf("template %s doesn't exist. Here are the existing templates : %v", tmplName, templates.GetTemplatesList())
	}

	t, err := template.New(tmplName+".html").ParseFS(templates.TemplatesFiles, fmt.Sprintf("%s/*.html", tmplName))
	if err != nil {
		return file, err
	}

	err = t.ExecuteTemplate(&file, tmplName+".html", cv)
	return file, err
}
