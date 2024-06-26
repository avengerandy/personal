package handler

import (
	"encoding/json"
	"html/template"
	"io/ioutil"
	"os"
)

type IndexSideProjects struct {
	Title   string
	Image   string
	Link    string
	Content string
}

type IndexConference struct {
	Title     string
	Authors   string
	Published string
	Date      string
	Link      string
}

type IndexSkill struct {
	Title   string
	Image   string
	Content template.HTML
}

type IndexProject struct {
	Id      string
	Title   string
	Image   string
	Content template.HTML
}

type IndexHandler struct {
	templateData struct {
		SiteName      string
		BodyClass      string
		Autobiography []template.HTML
		Projects      []IndexProject
		SideProjects  []IndexSideProjects
		Skills        []IndexSkill
		Conferences   []IndexConference
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
