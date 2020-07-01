package handler

import (
	"html/template"
	"os"
)

type IndexHandler struct{}

func (handler *IndexHandler) HandleTemplate(tmpl *template.Template) error {
	indexFile, err := os.Create("./docs/index.html")
	defer indexFile.Close()
	if err != nil {
		return err
	}
	err = tmpl.ExecuteTemplate(indexFile, "index.html", map[string]string{
		"siteName": "Index",
	})
	if err != nil {
		return err
	}
	return nil
}
