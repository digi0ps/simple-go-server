package main

import (
	"log"
	"net/http"
)

func logger(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.URL.Path)
		fn(w, r)
	}
}
