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
	err := tpl.Execute(os.Stdout, nil)
	handleError(err)

	err = tpl.ExecuteTemplate(os.Stdout, "vespa.gotpl", nil)
	handleError(err)

	err = tpl.ExecuteTemplate(os.Stdout, "one.gotpl", nil)
	handleError(err)

	err = tpl.ExecuteTemplate(os.Stdout, "two.gotpl", nil)
	handleError(err)
}

