package main

import (
	"flag"
	"fmt"
	"html/template"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// Page holds all the information we need to generate a new
// HTML page from a text file on the filesystem.
type Page struct {
	Data string
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func new_html(fileName string) {
	fileContents, _ := ioutil.ReadFile(fileName)

	page := Page{string(fileContents)}

	t := template.Must(template.New("template.tmpl").ParseFiles("template.tmpl"))

	newFile, _ := os.Create(strings.Split(fileName, ".")[0] + ".html")

	err := t.Execute(newFile, page)
	checkError(err)
}

func main() {
	// fileInput := flag.String("file", "first-post.txt", "The txt file passed in to parse")
	dirInput := flag.String("dir", "", "All .txt files in directory")
	flag.Parse()

	if *dirInput != "" {
		println("Converting all txt files in working directory")
		files, err := ioutil.ReadDir(*dirInput)
		checkError(err)

		for _, file := range files {
			if filepath.Ext(file.Name()) == ".txt" {
				new_html(file.Name())
				fmt.Println(file.Name())
			}
		}
	}

	// fileName := strings.Split(*fileInput, ".")[0] + ".html"
	// fileContents, _ := ioutil.ReadFile("data/" + *fileInput)

	// page := Page{string(fileContents)}

	// t := template.Must(template.New("template.tmpl").ParseFiles("template.tmpl"))

	// newFile, _ := os.Create(fileName)

	// err := t.Execute(newFile, page)
	// checkError(err)
}
