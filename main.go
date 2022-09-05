package main

import (
	"fmt"
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
	http.Redirect(w, r, os.Getenv("HOST")+os.Getenv("PORTS")+r.RequestURI, http.StatusMovedPermanently)
	fmt.Println("redirect to " + os.Getenv("HOST") + os.Getenv("PORTS") + r.RequestURI)
}
func handleRequest() {
	fs := http.FileServer(http.Dir("source"))

	http.Handle("/source/", http.StripPrefix("/source", fs))

	http.HandleFunc("/", homePage)
	http.HandleFunc("/donate", donatePage)
	go func() {
		fmt.Println("https started on port: " + os.Getenv("PORTS"))
		err := http.ListenAndServeTLS(":"+os.Getenv("PORTS"), "certs/fullchain.pem", "certs/privkey.pem", nil)
		if err != nil {
			log.Fatal(err)
		}
	}()
	fmt.Println("http started on port: " + os.Getenv("PORT"))
	err := http.ListenAndServe(":"+os.Getenv("PORT"), http.HandlerFunc(redirectTLS))
	if err != nil {
		log.Fatal(err)
	}
}
