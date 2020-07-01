package handler

import (
	"html/template"
)

type Handler interface {
	HandleTemplate(*template.Template) error
}
