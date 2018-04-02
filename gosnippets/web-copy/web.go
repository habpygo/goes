package main

import (
	"TheGoProgrammingLanguage/gosnippets/web-copy/controllers"
	"fmt"
	"log"
	"net/http"
)

func btcHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Handler btcHandler is called")
	controllers.RenderTemplate(w, r, "index.html", "")

}

func main() {
	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	http.Redirect(w, r, "/templates/index.html", http.StatusTemporaryRedirect)
	// })
	http.HandleFunc("/", btcHandler)

	fmt.Println("Listening (http://localhost:3000/) ...")
	fmt.Println("=======================================================")
	log.Fatal(http.ListenAndServe("localhost:3000", nil))

}
