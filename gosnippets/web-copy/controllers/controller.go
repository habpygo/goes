package controllers

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"path/filepath"
)

type WebApplication struct {
}

var LayoutDef string
var TemplateDef string

var Layout string

func RenderTemplate(w http.ResponseWriter, r *http.Request, templateName string, data interface{}) {

	pwd, _ := os.Getwd()

	if templateName == "index.html" {
		LayoutDef = filepath.Join(pwd, "templates", "layout.html")
		Layout = "layout.html"
	}

	TemplateDef = filepath.Join(pwd, "templates", templateName)

	fmt.Println("TemplateDef is: ", TemplateDef)
	fmt.Println("LayoutDef is: ", LayoutDef)

	// Return a 404 if the template doesn't exist
	info, err := os.Stat(TemplateDef)
	if err != nil {
		fmt.Println("template not found")
		if os.IsNotExist(err) {
			http.NotFound(w, r)
			return
		}
	}

	// Return a 404 if the request is for a directory
	if info.IsDir() {
		fmt.Println("dir not found")
		http.NotFound(w, r)
		return
	}

	resultTemplate, err := template.ParseFiles(TemplateDef, LayoutDef)
	if err != nil {
		// Log the detailed error
		fmt.Println(err.Error)
		// Return a generic "Internal Server Error" message
		http.Error(w, http.StatusText(500), 500)
		return
	}

	if err := resultTemplate.ExecuteTemplate(w, Layout, data); err != nil {
		fmt.Println(err.Error())
		http.Error(w, http.StatusText(500), 500)
	}
}
