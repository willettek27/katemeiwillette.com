package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func main() {
	port := ":8080"
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.HandleFunc("/", index)
	// http.HandleFunc("/about", about)
	http.HandleFunc("/projects", projects)
	http.HandleFunc("/contact", contact)
	fmt.Printf("Go Server running on port %s\n", port)

	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "index.html")
}

/* func about(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "about.html")
}
*/

func projects(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "projects.html")
}

func contact(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "contact.html")
}

/* func workInProgress(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "work-in-progress.html")
}
*/

func renderTemplate(w http.ResponseWriter, tmpl string) {

	t, err := template.ParseFiles("html/" + tmpl)
	if err != nil {
		log.Println("Template Error", err)
		http.Error(w, "Internal Issue", http.StatusInternalServerError)
		return
	}
	t.Execute(w, nil)
}
