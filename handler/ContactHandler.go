package handler

import (
	"encoding/json"
	"html/template"
	"io/ioutil"
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
	defer indexFile.Close()
	if err != nil {
		return err
	}

	data, err := ioutil.ReadFile("./setting/contact.json")
	if err != nil {
		return err
	}
	json.Unmarshal(data, &handler.templateData)

	err = tmpl.ExecuteTemplate(indexFile, "contact.html", handler.templateData)
	if err != nil {
		return err
	}
	return nil
}
