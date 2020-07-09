package handler

import (
	"encoding/json"
	"html/template"
	"io/ioutil"
	"os"
)

type AboutHandler struct {
	templateData struct {
		SiteName            string
		RecentAutobiography []template.HTML
		PastAutobiography   []template.HTML
	}
}

func (handler *AboutHandler) HandleTemplate(tmpl *template.Template) error {
	aboutFile, err := os.Create("./docs/about.html")
	defer aboutFile.Close()
	if err != nil {
		return err
	}

	data, err := ioutil.ReadFile("./setting/about.json")
	if err != nil {
		return err
	}
	json.Unmarshal(data, &handler.templateData)

	err = tmpl.ExecuteTemplate(aboutFile, "about.html", handler.templateData)
	if err != nil {
		return err
	}
	return nil
}
