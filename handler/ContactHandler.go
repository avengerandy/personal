package handler

import (
	"encoding/json"
	"html/template"
	"os"
)

type ContactHandler struct {
	templateData struct {
		SiteName   string
		BodyClass  string
		Images     []string
		Content    []string
	}
}

func (handler *ContactHandler) HandleTemplate(tmpl *template.Template) error {
	indexFile, err := os.Create("./docs/contact.html")
	if err != nil {
		return err
	}
	defer indexFile.Close()

	data, err := os.ReadFile("./setting/contact.json")
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, &handler.templateData);
	if err != nil {
		return err
	}

	err = tmpl.ExecuteTemplate(indexFile, "contact.html", handler.templateData)
	if err != nil {
		return err
	}
	return nil
}
