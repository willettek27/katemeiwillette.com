package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static/"))))

	http.HandleFunc("/", index)
	// http.HandleFunc("/about", about)
	http.HandleFunc("/projects", projects)
	http.HandleFunc("/contact", contact)
	fmt.Println("Server is running on http://katemeiwillette.com")

	err := http.ListenAndServe(":8080", nil)
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


func renderTemplate(w http.ResponseWriter, tmpl string) {

	t, err := template.ParseFiles("templates/" + tmpl)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	t.Execute(w, nil)
}
