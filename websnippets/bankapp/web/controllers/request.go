package controllers

import (
	"net/http"

	"TheGoProgrammingLanguage/websnippets/bankapp/web/bankaccounts"
)

func (app *Application) RequestHandler(w http.ResponseWriter, r *http.Request) {

	// Retrieve the values from the template request.html
	if r.FormValue("submitted") == "true" {

		bankStruct := bankaccounts.BankAccount{
			Name:          r.FormValue("name"),
			AccountNumber: r.FormValue("accountnumber"),
			Street:        r.FormValue("street"),
			HouseNumber:   r.FormValue("number"),
			City:          r.FormValue("city"),
		}
		renderTemplate(w, r, "request.html", bankStruct)
	}
}
