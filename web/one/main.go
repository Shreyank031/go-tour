package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

// Template cache
var templates = template.Must(template.ParseGlob("static/*.html"))

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/submit", submitHandler)

	fmt.Println("Server is running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	err := templates.ExecuteTemplate(w, "index.html", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	err := templates.ExecuteTemplate(w, "form.html", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func submitHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/form", http.StatusSeeOther)
		return
	}

	name := r.FormValue("name")
	email := r.FormValue("email")

	fmt.Fprintf(w, "Form submitted successfully!<br>Name: %s<br>Email: %s", name, email)
}
