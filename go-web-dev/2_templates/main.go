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

type Capital struct {
	Name string
	Country string
	Currency string
}

type Company struct {
	Name string
	Product string
}

func main() {
	Delhi := Capital{
		"Delhi",
		"India",
		"INR",
	}

	Moscow := Capital{
		"Moscow",
		"Russia",
		"RUB",
	}

	Beijing := Capital{
		"Beijing",
		"China",
		"CNY",
	}

	Tokyo := Capital{
		"Tokyo",
		"Japan",
		"JPY",
	}

	Paris := Capital{
		"Paris",
		"France",
		"EUR",
	}

	capitals := []Capital{Delhi, Moscow, Beijing, Tokyo, Paris}

	AsianPaint := Company{
		Name:    "Asian Paints",
		Product: "Paints",
	}

	TataSteel := Company{
		Name:    "TATA Steel",
		Product: "Steel",
	}

	MRF := Company{
		Name:    "MRF",
		Product: "Tyres",
	}

	companies := []Company{AsianPaint, TataSteel, MRF}

	data := struct{
		Capitals []Capital
		Companies []Company
	}{
		capitals,
		companies,
	}
	err := tpl.ExecuteTemplate(os.Stdout, "struct_list_struct.gohtml", data)
	handleError(err)
}
