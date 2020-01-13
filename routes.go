package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func render(out http.ResponseWriter, templateName string, file *File) {
	t, _ := template.ParseFiles("templates/" + templateName)
	t.Execute(out, file)
}

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

// FileHandler handles the /files/<filename> and fetches the file from local
func FileHandler(w http.ResponseWriter, r *http.Request) {
	filename := r.URL.Path[len("/files/"):]
	fmt.Println("FILENAME --> ", filename)

	if filename == "" {
		fmt.Fprintf(w, "Enter the correct path")
		return
	}

	file, err := loadFile(filename)

	if err != nil {
		fmt.Fprintf(w, "Error happened: %v", err)
	} else {
		render(w, "file.html", file)
	}
}

// CreateAndEditHandler handles the editing/creating of old/new files.
func CreateAndEditHandler(w http.ResponseWriter, r *http.Request) {
	filename := r.URL.Path[len("/edit/"):]

	file, err := loadFile(filename)

	if err != nil {
		render(w, "edit.html", &File{Title: filename})
	} else {
		render(w, "edit.html", file)
	}
}
