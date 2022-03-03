package main

import (
	"flag"
	"html/template"
	"io/ioutil"
	"os"
	"strings"
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

// Page holds all the information we need to generate a new
// HTML page from a text file on the filesystem.
type Page struct {
	Data string
}

func main() {
	fileInput := flag.String("file", "first-post.txt", "txt file to pass in")
	flag.Parse()

	fileName := strings.Split(*fileInput, ".")[0] + ".html"
	fileContents, _ := ioutil.ReadFile("data/" + *fileInput)

	page := Page{string(fileContents)}

	t := template.Must(template.New("template.tmpl").ParseFiles("template.tmpl"))

	newFile, _ := os.Create(fileName)

	err := t.Execute(newFile, page)
	checkError(err)
}
