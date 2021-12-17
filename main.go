package main

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"

	"github.com/go-chi/chi/v5"
)

func executeTemplate(w http.ResponseWriter, filepath string) {
	template, err := views.parseTemplate(filepath)
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

func mainHandler(w http.ResponseWriter, r *http.Request) {
	// file path handling in case of non-unix directories
	templatePath := filepath.Join("templates", "home.gohtml")
	executeTemplate(w, templatePath)
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	templatePath := filepath.Join("templates", "contact.gohtml")
	executeTemplate(w, templatePath)
}

func faqHandler(w http.ResponseWriter, r *http.Request) {
	templatePath := filepath.Join("templates", "faq.gohtml")
	executeTemplate(w, templatePath)
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
