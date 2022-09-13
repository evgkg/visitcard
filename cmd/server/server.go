package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"text/template"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseGlob("./pages/index.html"))
	err := t.ExecuteTemplate(w, "index.html", nil)
	if err != nil {
		log.Fatal(err)
	}
}
func DonatePage(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseGlob("./pages/donate.html"))
	err := t.ExecuteTemplate(w, "donate.html", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]bool{"ok": true})
}

func RedirectTLS(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, os.Getenv("HOST")+os.Getenv("PORTS")+r.RequestURI, http.StatusMovedPermanently)
	fmt.Println("redirect to " + os.Getenv("HOST") + os.Getenv("PORTS") + r.RequestURI)
}
func StartServer() {
	fs := http.FileServer(http.Dir("source"))

	http.Handle("/source/", http.StripPrefix("/source", fs))

	http.HandleFunc("/", HomePage)
	http.HandleFunc("/donate", DonatePage)
	http.HandleFunc("/api/health", HealthCheck)
	go func() {
		fmt.Println("https started on port: " + os.Getenv("PORTS"))
		err := http.ListenAndServeTLS(":"+os.Getenv("PORTS"), "certs/fullchain.pem", "certs/privkey.pem", nil)
		if err != nil {
			log.Fatal(err)
		}
	}()
	fmt.Println("http started on port: " + os.Getenv("PORT"))
	err := http.ListenAndServe(":"+os.Getenv("PORT"), http.HandlerFunc(RedirectTLS))
	if err != nil {
		log.Fatal(err)
	}
}
