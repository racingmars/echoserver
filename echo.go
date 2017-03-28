package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

var hostname string
var pid string

func main() {
	log.Printf("Starting echo")

	host, err := os.Hostname()
	if err != nil {
		log.Printf("Error setting hostname: %s", err)
		host = "[unknown]"
	}
	hostname = host

	pid = strconv.Itoa(os.Getpid())

	http.HandleFunc("/healthz", health)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)

	log.Printf("Ending echo")
}

func handler(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s [%s]", r.RemoteAddr, r.RequestURI)
	w.Header().Set("Content-type", "text/html")
	w.Header().Set("Connection", "close")
	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("Expires", "0")
	fmt.Fprintf(w, "<b>My hostname</b>: %s<br>", hostname)
	fmt.Fprintf(w, "<b>My PID</b>: %s<p>", pid)
	fmt.Fprintf(w, "<b>Client address</b>: %s<p>", r.RemoteAddr)
	fmt.Fprintf(w, "<b>URI</b>: %s<p>", html.EscapeString(r.RequestURI))
	for header, value := range r.Header {
		fmt.Fprintf(w, "<b>%s</b>: %s<br>", html.EscapeString(header), html.EscapeString(strings.Join(value, ",")))
	}
}

func health(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s [%s]", r.RemoteAddr, r.RequestURI)
	fmt.Fprint(w, "OK")
}

func getEnvOrDefault(env, defaultValue string) string {
	value := os.Getenv(env)
	if value == "" {
		value = defaultValue
	}
	return value
}
