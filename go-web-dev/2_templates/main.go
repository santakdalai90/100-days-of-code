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
	m := map[string]string{
		"Delhi": "India",
		"Paris": "France",
		"Tokyo": "Japan",
		"Washington DC": "USA",
		"Beijing": "China",
	}
	err := tpl.ExecuteTemplate(os.Stdout, "map.gohtml", m)
	handleError(err)
}

