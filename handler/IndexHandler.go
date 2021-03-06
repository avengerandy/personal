package handler

import (
	"encoding/json"
	"html/template"
	"io/ioutil"
	"os"
)

type IndexHandler struct {
	templateData struct {
		SiteName      string
		Autobiography []template.HTML
	}
}

func (handler *IndexHandler) HandleTemplate(tmpl *template.Template) error {
	indexFile, err := os.Create("./docs/index.html")
	defer indexFile.Close()
	if err != nil {
		return err
	}

	data, err := ioutil.ReadFile("./setting/index.json")
	if err != nil {
		return err
	}
	json.Unmarshal(data, &handler.templateData)

	err = tmpl.ExecuteTemplate(indexFile, "index.html", handler.templateData)
	if err != nil {
		return err
	}
	return nil
}
