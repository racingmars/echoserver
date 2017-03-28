package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
	"strings"
)

func main() {
	log.Printf("Starting echo")

	http.HandleFunc("/healthz", health)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)

	log.Printf("Ending echo")
}

func handler(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s [%s]", r.RemoteAddr, r.RequestURI)
	w.Header().Set("Content-type", "text/html")
	fmt.Fprintf(w, "<b>URI</b>: %s<p>", html.EscapeString(r.RequestURI))
	for header, value := range r.Header {
		fmt.Fprintf(w, "<b>%s</b>: %s<br>", html.EscapeString(header), html.EscapeString(strings.Join(value, ",")))
	}
}

func health(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s [%s]", r.RemoteAddr, r.RequestURI)
	fmt.Fprint(w, "OK")
}
