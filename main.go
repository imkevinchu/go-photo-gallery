package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func mainHandler(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(writer, "<h1>Hello world</h1>")
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

func main() {
	router := chi.NewRouter()

	router.Get("/", mainHandler)
	router.Get("/contact", contactHandler)
	router.Get("/faq", faqHandler)
	router.NotFound(http.NotFound)

	fmt.Println("Starting server on :3000")
	http.ListenAndServe(":3000", router)

}
