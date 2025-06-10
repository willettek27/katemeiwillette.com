package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func main() {
	// Serve static files from the static directory
	http.Handle("/styles/", http.StripPrefix("/styles/", http.FileServer(http.Dir("styles/"))))

	http.HandleFunc("/", home)
	http.HandleFunc("/projects", projects)
	http.HandleFunc("/about", about)
	http.HandleFunc("/contact", contact)

	fmt.Println("🚀 Server is running on http://localhost:8080")

	// Start the web server on port 8080 and listen to incoming requests
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func home(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "home.html")
}

func projects(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "projects.html")
}

func about(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "about.html")
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
