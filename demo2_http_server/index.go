package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

func main() {

	// Define path "/"
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Home, %q", html.EscapeString(r.URL.Path))
	})

	// Define path "/profile"
	http.HandleFunc("/profile", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Profile, %q", html.EscapeString(r.URL.Path))
	})

	// Define path "/login" and get string query
	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Login, %s, %s", r.URL.Query().Get("username"), r.URL.Query().Get("password"))
	})

	// Run and show fatal log
	log.Fatal(http.ListenAndServe(":82", nil))
}
