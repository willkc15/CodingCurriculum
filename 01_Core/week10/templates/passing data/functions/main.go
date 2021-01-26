package main

import (
	"log"
	"math"
	"os"
	"strings"
	"text/template"
)

var fm = template.FuncMap{
	"uc":   strings.ToUpper,
	"dbl":  double,
	"sq":   square,
	"sqrt": squareRoot,
}

var tpl *template.Template

func double(i int) int {
	return 2 * i
}

func square(i int) int {
	return i * i
}

func squareRoot(i int) float64 {
	return math.Sqrt(float64(i))
}

func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("functions.gohtml"))
}

func main() {
	err := tpl.ExecuteTemplate(os.Stdout, "functions.gohtml", 2)
	if err != nil {
		log.Fatalln(err)
	}
}
