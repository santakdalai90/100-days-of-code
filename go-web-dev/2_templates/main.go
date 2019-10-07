package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func handleError(e error) {
	if e != nil {
		log.Fatalln(e.Error())
	}
}

type geo struct {
	Capital string
	Country string
	Currency string
}

func main() {
	s := geo{
		"Delhi",
		"India",
		"INR",
	}
	err := tpl.ExecuteTemplate(os.Stdout, "struct.gohtml", s)
	handleError(err)
}
