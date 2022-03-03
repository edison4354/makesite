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
	inputFile := flag.String("file", "first-post.txt", "txt file to pass in")
	flag.Parse()

	fileName := strings.Split(*inputFile, ".")[0] + ".html"
	fileContents, _ := ioutil.ReadFile("data/" + *inputFile)

	page := Page{string(fileContents)}

	t := template.Must(template.New("template.tmpl").ParseFiles("template.tmpl"))

	newFile, _ := os.Create(fileName)

	err := t.Execute(newFile, page)
	checkError(err)
}
