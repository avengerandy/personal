package handler

import (
	"html/template"
)

type ContactData struct {
	SiteName  string
	BodyClass string
	Images    []string
	Content   []string
}

type ContactHandler struct{}

func (handler *ContactHandler) HandleTemplate(tmpl *template.Template) error {
	return RenderTemplateFromJSON[ContactData](
		"./setting/contact.json",
		"./docs/contact.html",
		"contact.html",
		tmpl,
	)
}
