package main

import (
	"fmt"
	"net/http"
)

func mainHandler(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(writer, "<h1>Hello world</h1>")
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>Contact</h1><p>Get in touch</p>")
}

func pathHandler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		mainHandler(w, r)
	case "/contact":
		contactHandler(w, r)
	default:
		http.NotFound(w, r)
	}
}

func main() {
	http.HandleFunc("/", pathHandler)
	fmt.Println("Starting server on :3000")
	http.ListenAndServe(":3000", nil)

}
