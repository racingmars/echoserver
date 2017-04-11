package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
	"os"
	"sort"
	"strconv"
	"strings"
)

var hostname string
var pid string
var appname string

func main() {
	log.Printf("Starting echo")

	host, err := os.Hostname()
	if err != nil {
		log.Printf("Error setting hostname: %s", err)
		host = "[unknown]"
	}
	hostname = host

	pid = strconv.Itoa(os.Getpid())

	appname = os.Getenv("APPNAME")

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
	fmt.Fprintf(w, "<b>My PID</b>: %s<br>", pid)
	if appname != "" {
		fmt.Fprintf(w, "<b>My appname</b>: %s", appname)
	}
	fmt.Fprintf(w, "<p><b>Client address</b>: %s<p>", r.RemoteAddr)
	fmt.Fprintf(w, "<b>URI</b>: %s<p>", html.EscapeString(r.RequestURI))

	// Sort the headers so they stay in a reasonable and repeatable order.
	keys := make([]string, len(r.Header))
	i := 0
	for header := range r.Header {
		keys[i] = header
		i++
	}
	sort.Strings(keys)
	for _, header := range keys {
		fmt.Fprintf(w, "<b>%s</b>: %s<br>", html.EscapeString(header), html.EscapeString(strings.Join(r.Header[header], ",")))
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
