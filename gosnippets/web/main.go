package main

import (
	"TheGoProgrammingLanguage/gosnippets/web/bankaccounts"
	"html/template"
	"net/http"
)

func main() {
	tmpl, err := template.New("").ParseFiles("layout.html", "forms.html")
	err = tmpl.ExecuteTemplate(w, "layout", ba)
	//tmpl := template.Must(template.ParseFiles("templates/forms.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		if r.Method != http.MethodPost {
			tmpl.Execute(w, nil)
			return
		}

		ba := bankaccounts.Bankaccount{
			Name:          r.FormValue("name"),
			AccountNumber: r.FormValue("accountnumber"),
			Street:        r.FormValue("street"),
			HouseNumber:   r.FormValue("number"),
			City:          r.FormValue("city"),
		}

		newAccount := bankaccounts.NewBankAccount(
			ba.Name,
			ba.AccountNumber,
			ba.Street,
			ba.HouseNumber,
			ba.City)

		//tmpl.Execute(w, struct{ Success bool }{true})
		tmpl.Execute(w, ba)
	})

	http.ListenAndServe(":8080", nil)
}
