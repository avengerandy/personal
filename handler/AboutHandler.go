package handler

import (
	"html/template"
	"os"
)

type AboutHandler struct{}

func (handler *AboutHandler) HandleTemplate(tmpl *template.Template) error {
	aboutFile, err := os.Create("./docs/about.html")
	defer aboutFile.Close()
	if err != nil {
		return err
	}
	err = tmpl.ExecuteTemplate(aboutFile, "about.html", map[string]string{
		"siteName": "About",
	})
	if err != nil {
		return err
	}
	return nil
}
