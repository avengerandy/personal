package handler

import (
	"encoding/json"
	"html/template"
	"os"
)

type Work_Education struct {
	Title   []string
	Content []string
	Time    string
}

type ExperienceHandler struct {
	templateData struct {
		SiteName   string
		BodyClass  string
		Works      []Work_Education
		Educations []Work_Education
	}
}

func (handler *ExperienceHandler) HandleTemplate(tmpl *template.Template) error {
	indexFile, err := os.Create("./docs/experience.html")
	if err != nil {
		return err
	}
	defer indexFile.Close()

	data, err := os.ReadFile("./setting/experience.json")
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, &handler.templateData)
	if err != nil {
		return err
	}

	err = tmpl.ExecuteTemplate(indexFile, "experience.html", handler.templateData)
	if err != nil {
		return err
	}
	return nil
}
