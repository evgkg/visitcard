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

	t := template.Must(template.ParseGlob("./pages/index.html"))
	err := t.ExecuteTemplate(w, "index.html", nil)
	if err != nil {
		log.Fatal(err)
	}
}
func donatePage(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseGlob("./pages/donate.html"))
	err := t.ExecuteTemplate(w, "donate.html", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func redirectTLS(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, ":"+os.Getenv("PORTS")+r.RequestURI, http.StatusMovedPermanently)
}
func handleRequest() {
	fs := http.FileServer(http.Dir("source"))

	http.Handle("/source/", http.StripPrefix("/source", fs))

	http.HandleFunc("/", homePage)
	http.HandleFunc("/donate", donatePage)
	go func() {
		err := http.ListenAndServeTLS(":"+os.Getenv("PORTS"), "certs/fullchain.pem", "certs/privkey.pem", nil)
		if err != nil {
			log.Fatal(err)
		}
	}()
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), http.HandlerFunc(redirectTLS)))
}
