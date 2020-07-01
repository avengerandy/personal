package main

import (
	"github.com/avengerandy/personal/handler"
	"html/template"
	"log"
)

func main() {

	pages := map[string]handler.Handler{
		"Index": &handler.IndexHandler{},
		"About": &handler.AboutHandler{},
	}

	tmpl, err := template.ParseGlob("./template/*.html")
	if err != nil {
		log.Fatal("Error loading templates: " + err.Error())
	}

	for _, value := range pages {
		err := value.HandleTemplate(tmpl)
		if err != nil {
			log.Fatal(err.Error())
		}
	}
}
