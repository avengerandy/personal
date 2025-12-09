package handler

import (
	"encoding/json"
	"html/template"
	"os"
)

type AboutHandler struct {
	templateData struct {
		SiteName            string
		BodyClass            string
		RecentAutobiography []template.HTML
		PastAutobiography   []template.HTML
	}
}

func (handler *AboutHandler) HandleTemplate(tmpl *template.Template) error {
	aboutFile, err := os.Create("./docs/about.html")
	if err != nil {
		return err
	}
	defer aboutFile.Close()

	data, err := os.ReadFile("./setting/about.json")
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, &handler.templateData);
	if err != nil {
		return err
	}

	err = tmpl.ExecuteTemplate(aboutFile, "about.html", handler.templateData)
	if err != nil {
		return err
	}
	return nil
}
