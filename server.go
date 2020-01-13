package main

import (
	"fmt"
	"log"
	"net/http"
)

const port int = 8888

func setupRoutes() {
	http.HandleFunc("/files/", logger(FileHandler))
	http.HandleFunc("/edit/", logger(CreateAndEditHandler))
	http.HandleFunc("/save/", logger(SaveHandler))
	http.HandleFunc("/", logger(HomePageHandler))
}

func startServer(port int) {
	setupRoutes()

	p := fmt.Sprintf(":%d", port)
	log.Fatal(http.ListenAndServe(p, nil))
}

func main() {
	fmt.Printf("Starting server at %d...\n", port)
	startServer(port)
}
