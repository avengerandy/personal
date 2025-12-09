package handler

import (
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

type ProjectHandler struct{}

func (handler *ProjectHandler) HandleTemplate(tmpl *template.Template) error {
    files, err := os.ReadDir("./setting/projects")
    if err != nil {
        return err
    }

    for _, file := range files {
        jsonPath := "./setting/projects/" + file.Name()
        outputPath := "./docs/" + file.Name()[:len(file.Name())-5] + ".html" // .json

        if err := RenderTemplateFromJSON[Project](jsonPath, outputPath, "project.html", tmpl); err != nil {
            return err
        }
    }

    return nil
}
