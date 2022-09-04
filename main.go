package main

import (
	"log"
	"net/http"
	"os"
	"text/template"
)

func main() {
	handleRequest()
}

func homePage(w http.ResponseWriter, r *http.Request) {

	t := template.Must(template.ParseGlob("index.html"))
	err := t.ExecuteTemplate(w, "index.html", nil)
	if err != nil {
		log.Fatal(err)
	}
}
func donatePage(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseGlob("donate.html"))
	err := t.ExecuteTemplate(w, "donate.html", nil)
	if err != nil {
		log.Fatal(err)
	}
}
func handleRequest() {
	fs := http.FileServer(http.Dir("templates"))

	http.Handle("/templates/", http.StripPrefix("/templates", fs))

	http.HandleFunc("/", homePage)
	http.HandleFunc("/donate", donatePage)
	err := http.ListenAndServe(":"+os.Getenv("PORT"), nil)
	if err != nil {
		log.Fatal(err)
	}

}
