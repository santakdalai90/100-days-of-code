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
	Delhi := geo{
		"Delhi",
		"India",
		"INR",
	}

	Moscow := geo{
		"Moscow",
		"Russia",
		"RUB",
	}

	Beijing := geo{
		"Beijing",
		"China",
		"CNY",
	}

	Tokyo := geo{
		"Tokyo",
		"Japan",
		"JPY",
	}

	Paris := geo{
		"Paris",
		"France",
		"EUR",
	}

	capitals := []geo{Delhi, Moscow, Beijing, Tokyo, Paris}
	err := tpl.ExecuteTemplate(os.Stdout, "list_of_structs.gohtml", capitals)
	handleError(err)
}
