package handler

import (
	"html/template"
)

type WorkEducation struct {
	Title   []string
	Content []string
	Time    string
}

type ExperienceData struct {
	SiteName   string
	BodyClass  string
	Works      []WorkEducation
	Educations []WorkEducation
}

type ExperienceHandler struct{}

func (handler *ExperienceHandler) HandleTemplate(tmpl *template.Template) error {
	return RenderTemplateFromJSON[ExperienceData](
		"./setting/experience.json",
		"./docs/experience.html",
		"experience.html",
		tmpl,
	)
}
