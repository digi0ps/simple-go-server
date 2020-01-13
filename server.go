package main

import (
	"fmt"
	"log"
	"net/http"
)

const port int = 8888

func setupRoutes() {
	http.HandleFunc("/", HomePageHandler)
	http.HandleFunc("/profile", ProfilePageHandler)
	http.HandleFunc("/files/", FileHandler)
	http.HandleFunc("/edit/", CreateAndEditHandler)
	http.HandleFunc("/save/", SaveHandler)
}

func startServer(port int) {
	setupRoutes()

	p := fmt.Sprintf(":%d", port)
	log.Fatal(http.ListenAndServe(p, nil))
}

func main() {
	fmt.Printf("Starting server at %d...", port)
	startServer(port)
}
