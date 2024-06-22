package handler

import (
	"encoding/json"
	"html/template"
	"io/ioutil"
	"os"
)

type Project struct {
	Id       string
	SiteName string
	Title    string
	ImageSet []string
	Content  string
	Note     template.HTML
}

type ProjectHandler struct {
	templateData struct {
		Projects []Project
	}
}

func (handler *ProjectHandler) HandleTemplate(tmpl *template.Template) error {
	data, err := ioutil.ReadFile("./setting/project.json")
	if err != nil {
		return err
	}

	json.Unmarshal(data, &handler.templateData)
	for _, project := range handler.templateData.Projects {
		projectFile, err := os.Create("./docs/project_" + project.Id + ".html")
		defer projectFile.Close()
		if err != nil {
			return err
		}

		err = tmpl.ExecuteTemplate(projectFile, "project.html", project)
		if err != nil {
			return err
		}
	}

	return nil
}
