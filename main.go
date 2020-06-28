package main

import (
	"html/template"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	mainTemplate, err := template.ParseFiles("./template/main.html")
	if err != nil {
		log.Print(err)
		return
	}
	indexTemplateFile, err := ioutil.ReadFile("./template/index.html")
	if err != nil {
		log.Print(err)
		return
	}
	indexFile, err := os.Create("./docs/index.html")
	defer indexFile.Close()
	if err != nil {
		log.Print(err)
		return
	}
	aboutTemplateFile, err := ioutil.ReadFile("./template/about.html")
	if err != nil {
		log.Print(err)
		return
	}
	aboutFile, err := os.Create("./docs/about.html")
	defer aboutFile.Close()
	if err != nil {
		log.Print(err)
		return
	}
	err = mainTemplate.Execute(indexFile, map[string]template.HTML{
		"content": template.HTML(string(indexTemplateFile)),
	})
	if err != nil {
		log.Print(err)
		return
	}
	err = mainTemplate.Execute(aboutFile, map[string]template.HTML{
		"content": template.HTML(string(aboutTemplateFile)),
	})
	if err != nil {
		log.Print(err)
		return
	}
}
