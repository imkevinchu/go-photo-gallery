package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/go-chi/chi/v5"
)

func mainHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	// file path handling in case of non-unix directories
	templatePath := filepath.Join("templates", "home.gohtml")

	template, err := template.ParseFiles(templatePath)
	if err != nil {
		log.Printf("parsing template: %v", err)
		http.Error(w, "error parsing template", http.StatusInternalServerError)
		return
	}
	err = template.Execute(w, nil)
	if err != nil {
		log.Printf("executing template: %v", err)
		http.Error(w, "error executing template", http.StatusInternalServerError)
		return
	}
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>Contact</h1><p>Get in touch</p>")
}

func faqHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text-html; charset=utf-8")
	fmt.Fprint(w, `<h1>Questions</h1>
	<ul>
		<li>Who?</li>
		<li>What?</li>
		<li>Why?</li>
	</ul>`)
}

func userHandler(w http.ResponseWriter, r *http.Request) {
	userid := chi.URLParam(r, "userid")
	fmt.Fprint(w, userid)
}

func main() {
	r := chi.NewRouter()

	r.Get("/", mainHandler)
	r.Get("/contact", contactHandler)
	r.Get("/faq", faqHandler)
	r.Get("/user/{userid}", userHandler)
	r.NotFound(http.NotFound)

	fmt.Println("Starting server on :3000")
	http.ListenAndServe(":3000", r)

}
