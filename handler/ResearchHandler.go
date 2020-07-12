package handler

import (
	"encoding/json"
	"html/template"
	"io/ioutil"
	"os"
)

type ResearchHandler struct {
	templateData struct {
		SiteName string
	}
}

func (handler *ResearchHandler) HandleTemplate(tmpl *template.Template) error {
	researchFile, err := os.Create("./docs/research.html")
	defer researchFile.Close()
	if err != nil {
		return err
	}

	data, err := ioutil.ReadFile("./setting/research.json")
	if err != nil {
		return err
	}
	json.Unmarshal(data, &handler.templateData)

	err = tmpl.ExecuteTemplate(researchFile, "research.html", handler.templateData)
	if err != nil {
		return err
	}
	return nil
}
