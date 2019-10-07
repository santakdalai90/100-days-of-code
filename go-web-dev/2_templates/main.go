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
	err := tpl.ExecuteTemplate(os.Stdout, "one.gohtml", "work hard and be smart")
	handleError(err)
}

