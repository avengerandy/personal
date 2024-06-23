package handler

import (
	"encoding/json"
	"html/template"
	"io/ioutil"
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
		Works      []Work_Education
		Educations []Work_Education
	}
}

func (handler *ExperienceHandler) HandleTemplate(tmpl *template.Template) error {
	indexFile, err := os.Create("./docs/experience.html")
	defer indexFile.Close()
	if err != nil {
		return err
	}

	data, err := ioutil.ReadFile("./setting/experience.json")
	if err != nil {
		return err
	}
	json.Unmarshal(data, &handler.templateData)

	err = tmpl.ExecuteTemplate(indexFile, "experience.html", handler.templateData)
	if err != nil {
		return err
	}
	return nil
}
