package handler

import (
	"encoding/json"
	"html/template"
	"os"
)

type Project struct {
	Id        string
	SiteName  string
	BodyClass string
	Title     string
	ImageSet  []string
	Content   []string
	Note      template.HTML
}

type ProjectHandler struct {
	templateData struct {
		Projects []Project
	}
}

func (handler *ProjectHandler) HandleTemplate(tmpl *template.Template) error {
	data, err := os.ReadFile("./setting/project.json")
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, &handler.templateData)
	if err != nil {
		return err
	}
	for _, project := range handler.templateData.Projects {
		projectFile, err := os.Create("./docs/project_" + project.Id + ".html")
		if err != nil {
			return err
		}

		err = tmpl.ExecuteTemplate(projectFile, "project.html", project)
		projectFile.Close()
		if err != nil {
			return err
		}
	}

	return nil
}
