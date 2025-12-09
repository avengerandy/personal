package handler

import (
	"html/template"
)

type IndexSideProject struct {
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

type IndexData struct {
	SiteName      string
	BodyClass     string
	Autobiography []template.HTML
	Projects      []IndexProject
	SideProjects  []IndexSideProject
	Skills        []IndexSkill
	Conferences   []IndexConference
}

type IndexHandler struct{}

func (handler *IndexHandler) HandleTemplate(tmpl *template.Template) error {
	return RenderTemplateFromJSON[IndexData](
		"./setting/index.json",
		"./docs/index.html",
		"index.html",
		tmpl,
	)
}
