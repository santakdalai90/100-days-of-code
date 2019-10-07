package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init(){
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func handleError(e error) {
	if e != nil {
		log.Fatalln(e.Error())
	}
}

func main(){
	s := []string{
		"Delhi",
		"Paris",
		"Tokyo",
		"Washington DC",
		"Beijing",
	}
	err := tpl.ExecuteTemplate(os.Stdout, "slice.gohtml", s)
	handleError(err)
}

