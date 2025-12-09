package handler

import (
	"html/template"
)

type AboutData struct {
	SiteName            string
	BodyClass           string
	RecentAutobiography []template.HTML
	PastAutobiography   []template.HTML
}

type AboutHandler struct{}

func (handler *AboutHandler) HandleTemplate(tmpl *template.Template) error {
	return RenderTemplateFromJSON[AboutData](
		"./setting/about.json",
		"./docs/about.html",
		"about.html",
		tmpl,
	)
}
