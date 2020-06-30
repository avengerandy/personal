package main

import (
	"html/template"
	"log"
	"os"
)

func main() {

	tmpl, err := template.ParseGlob("./template/*.html")
	if err != nil {
		log.Fatal("Error loading templates: " + err.Error())
	}

	indexFile, err := os.Create("./docs/index.html")
	defer indexFile.Close()
	if err != nil {
		log.Fatal("Error create output index: " + err.Error())
	}
	aboutFile, err := os.Create("./docs/about.html")
	defer aboutFile.Close()
	if err != nil {
		log.Fatal("Error create output about: " + err.Error())
	}

	err = tmpl.ExecuteTemplate(indexFile, "index.html", map[string]string{
		"siteName": "Index",
	})
	if err != nil {
		log.Fatal("Error write output index: " + err.Error())
	}
	err = tmpl.ExecuteTemplate(aboutFile, "about.html", map[string]string{
		"siteName": "About",
	})
	if err != nil {
		log.Fatal("Error write output about: " + err.Error())
	}
}
