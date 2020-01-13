package main

import "io/ioutil"

// File strcuture for a file
type File struct {
	title string
	body  []byte
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

	return &File{title: title, body: body}, nil
}

func (f *File) saveFile() error {
	filename := getFileName(f.title)
	return ioutil.WriteFile(filename, f.body, 0600)
}
