package main

import (
	"fmt"
	"net/http"
)

func handlerFunc(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprint(writer, "<h1>Hello world</h1>")
}

func main() {
	http.HandleFunc("/", handlerFunc)
	fmt.Println("Starting server on :3000")
	http.ListenAndServe(":3000", nil)
}
