package web

import (
	"TheGoProgrammingLanguage/websnippets/bankapp/web/controllers"
	"fmt"
	"net/http"
)

func Serve(bankwebpage *controllers.Application) {
	fs := http.FileServer(http.Dir("web/assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))

	// The Handlers
	http.HandleFunc("/index.html", bankwebpage.IndexHandler)

	fmt.Println("Listening (http://localhost:3000/) ...")
	fmt.Println("=======================================================")
	http.ListenAndServe(":3000", nil)

}
