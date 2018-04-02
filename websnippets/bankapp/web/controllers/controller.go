package controllers

import (
	"TheGoProgrammingLanguage/websnippets/bankapp/web/bankaccounts"
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"html/template"
)

//"TheGoProgrammingLanguage/websnippets/bankapp/web/bankaccounts"

type Application struct {
	Bank *bankaccounts.BankAccount
}

//Lp layout definition
var Lp string

//Tp template definition
var Tp string

//Layout is layout
var Layout string

func renderTemplate(w http.ResponseWriter, r *http.Request, templateName string, data interface{}) {

	if templateName == "index.html" {
		Lp = filepath.Join("web", "templates", "layout.html")
		Layout = "layout"
	}

	Tp = filepath.Join("web", "templates", templateName)
	fmt.Println("templateName is: ", Tp)
	fmt.Println("layout is: ", Lp)

	// Return a 404 if the template doesn't exist
	info, err := os.Stat(Tp)
	if err != nil {
		if os.IsNotExist(err) {
			http.NotFound(w, r)
			return
		}
	}

	// Return a 404 if the request is for a directory
	if info.IsDir() {
		http.NotFound(w, r)
		return
	}

	resultTemplate, err := template.ParseFiles(Tp, Lp)
	if err != nil {
		// Log the detailed error
		fmt.Println(err.Error())
		// Return a generic "Internal Server Error" message
		http.Error(w, http.StatusText(500), 500)
		return
	}
	if err := resultTemplate.ExecuteTemplate(w, Layout, data); err != nil {
		fmt.Println(err.Error())
		http.Error(w, http.StatusText(500), 500)
	}
}
func (app *Application) IndexHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, r, "index.html", nil)
}
