package main

import (
    "fmt"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Content-Type","text/html")
    fmt.Fprintf(w, "<h4>Guessing Game!<h4>")
}

func main() {
    http.HandleFunc("/", handler)
    http.ListenAndServe(":8080", nil)
}