package main

import (
	"fmt"
	"net/http"
)

// HomePageHandler handles the / route (aka home)
func HomePageHandler(w http.ResponseWriter, r *http.Request) {
	var path string = r.URL.Path[1:]
	fmt.Fprintf(w, "Welcome to the my go site.\n")

	if path == "" {
		fmt.Fprintf(w, "You're in homepage")
	} else {
		fmt.Fprintf(w, "You're in page %s", path)
	}

	fmt.Println("USER AGENT", r.UserAgent())
	fmt.Println("URL", path)
}

// ProfilePageHandler handles the /profile route
func ProfilePageHandler(w http.ResponseWriter, r *http.Request) {
	ua := r.UserAgent()

	fmt.Fprintf(w, "Welcome User\n")
	fmt.Fprintf(w, "Your user agent is \t%s", ua)
}
