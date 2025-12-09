package handler

import (
	"encoding/json"
	"html/template"
	"os"
)

type Handler interface {
	HandleTemplate(*template.Template) error
}

func RenderTemplateFromJSON[T any](
	jsonPath,
	outputPath,
	tmplName string,
	tmpl *template.Template,
) error {
	var jsonStruct T
	jsonData, err := os.ReadFile(jsonPath)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(jsonData, &jsonStruct); err != nil {
		return err
	}

	outputFile, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer outputFile.Close()

	return tmpl.ExecuteTemplate(outputFile, tmplName, jsonStruct)
}
