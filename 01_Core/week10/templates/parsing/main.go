package main

import (
	"log"
	"os"
	"text/template"
)

//global variable (package is template, type is Template)
var tmpl *template.Template

func init() {
	// Must handles the error for us, takes as arguments a pointer
	// to a template and an error, which is exactly what parseglob returns
	tmpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	tmpl, err := template.ParseFiles("tpl.gohtml")
	if err != nil {
		log.Fatalln(err)
	}

	err = tmpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}

	tmpl, err = tmpl.ParseFiles("test1.literallyanything", "test2.test")
	if err != nil {
		log.Fatalln(err)
	}

	err = tmpl.ExecuteTemplate(os.Stdout, "test1.literallyanything", nil)
	if err != nil {
		log.Fatalln(err)
	}

	newFile, err := os.Create("index.html")
	if err != nil {
		log.Println("error creating file", err)
	}

	err = tmpl.ExecuteTemplate(newFile, "test2.test", nil)
	if err != nil {
		log.Fatalln(err)
	}

	defer newFile.Close()

	// Parse a glob, basically parses all files in a folder
	tmpl, err = template.ParseGlob("FilesGlob/*")
	if err != nil {
		log.Fatalln(err)
	}
	err = tmpl.Execute(os.Stdout, nil)
}
