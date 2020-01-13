package main

import (
	"errors"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
)

var templates = template.Must(template.ParseFiles("templates/main.html", "templates/file.html", "templates/edit.html"))

func render(out http.ResponseWriter, templateName string, file *File) {
	tmpl := "templates/" + templateName
	templates.ExecuteTemplate(out, tmpl, file)
}

// TempFiles for sending the file list in Home
type TempFiles struct {
	Names []string
}

// HomePageHandler handles the / route (aka home)
func HomePageHandler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path[1:]

	fmt.Println("Path")

	if path == "" {
		files, _ := ioutil.ReadDir("files")
		var names []string

		for _, file := range files {
			name := file.Name()
			name = name[:len(name)-len(".gofl")]
			names = append(names, name)
		}

		fmt.Println(names)

		err := templates.ExecuteTemplate(w, "templates/main.html", &TempFiles{Names: names})
		fmt.Println("Err", err)
	} else {
		http.Error(w, errors.New("404 Route Not Found").Error(), http.StatusNotFound)
	}
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
		http.Redirect(w, r, "/edit/"+filename, http.StatusFound)
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

// SaveHandler is a controller for saving the file
func SaveHandler(w http.ResponseWriter, r *http.Request) {
	title := r.FormValue("title")
	body := r.FormValue("body")

	fmt.Println("POST --> ", title, body)

	file := File{Title: title, Body: []byte(body)}
	file.saveFile()

	http.Redirect(w, r, "/files/"+title, http.StatusFound)
}
