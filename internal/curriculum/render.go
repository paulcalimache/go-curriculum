package curriculum

import (
	"bytes"
	"os"

	"github.com/paulcalimache/go-curriculum/internal/pdf"
	"github.com/paulcalimache/go-curriculum/internal/templates"
)

func (cv *CV) Render(output string, tmplName string) error {
	file, err := templates.Templetize(tmplName, cv)
	if err != nil {
		return err
	}

	// Create output directory
	err = os.MkdirAll(output, os.ModePerm)
	if err != nil {
		return err
	}
	err = os.Chdir(output)
	if err != nil {
		return err
	}

	err = saveAsHTML(file)
	if err != nil {
		return err
	}

	err = pdf.ConvertHtmlToPdf(file)
	if err != nil {
		return err
	}
	return nil
}

func saveAsHTML(file bytes.Buffer) error {
	return os.WriteFile("curriculum.html", file.Bytes(), 0644)
}
