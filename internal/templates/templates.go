package templates

import (
	"embed"
)

//go:embed *
var TemplatesFiles embed.FS

var templatesList = []string{
	"classic",
}

func GetTemplatesList() []string {
	return templatesList
}
