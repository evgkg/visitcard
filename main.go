package main

import (
	"log"
	"net/http"
	"text/template"
)

func main() {
	handleRequest()
}

func homePage(w http.ResponseWriter, r *http.Request) {

	t := template.Must(template.ParseGlob("./pages/index.html"))
	err := t.ExecuteTemplate(w, "index.html", nil)
	if err != nil {
		log.Fatal(err)
	}
}
func donatePage(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseGlob(".pages/donate.html"))
	err := t.ExecuteTemplate(w, "donate.html", nil)
	if err != nil {
		log.Fatal(err)
	}
}
func handleRequest() {
	fs := http.FileServer(http.Dir("styles"))

	http.Handle("/styles/", http.StripPrefix("/styles", fs))

	http.HandleFunc("/", homePage)
	http.HandleFunc("/donate", donatePage)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}

}
