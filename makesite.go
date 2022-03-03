package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"os"
)

// Page holds all the information we need to generate a new
// HTML page from a text file on the filesystem.
type Page struct {
	Paragraph string
}

func main() {
	fileContents, err := ioutil.ReadFile("data/first-post.txt")
	if err != nil {
		panic(err)
	}

	newFile, err := os.Create("new.html")
	if err != nil {
		panic(err)
	}

	page := Page{}

	page.Paragraph = string(fileContents)

	t := template.Must(template.New("template.tmpl").ParseFiles("template.tmpl"))

	t.Execute(newFile, page)

	fmt.Print(string(fileContents))
}
