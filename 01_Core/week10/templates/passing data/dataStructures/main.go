package main

import (
	"html/template"
	"log"
	"os"
)

var tpl *template.Template

// Must use capital letters for struct to be visible outside of current package
type person struct {
	Fname string
	Lname string
}

func init() {
	tpl = template.Must(template.ParseFiles("structs.gohtml", "slices.gohtml"))
}

func main() {

	sliceExample := []string{"Invest", "in", "Bitcoin", "OR", "else"}
	structExample := person{
		Fname: "Kelton",
		Lname: "Williams",
	}

	err := tpl.ExecuteTemplate(os.Stdout, "slices.gohtml", sliceExample)
	if err != nil {
		log.Fatalln(err)
	}
	err = tpl.ExecuteTemplate(os.Stdout, "structs.gohtml", structExample)
	if err != nil {
		log.Fatalln(err)
	}

}
