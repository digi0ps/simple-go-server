package main

import "io/ioutil"

// File strcuture for a file
type File struct {
	Title string
	Body  []byte
}

func getFileName(title string) string {
	return "./files/" + title + ".gofl"
}

func loadFile(title string) (*File, error) {
	filename := getFileName(title)

	body, err := ioutil.ReadFile(filename)

	if err != nil {
		return nil, err
	}

	return &File{Title: title, Body: body}, nil
}

func (f *File) saveFile() error {
	filename := getFileName(f.Title)
	return ioutil.WriteFile(filename, f.Body, 0600)
}
