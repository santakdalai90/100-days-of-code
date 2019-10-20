package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"text/template"
)

var tpl *template.Template

func filterVowel(r rune) rune {
	vowels := "aeiouAEIOU"
	if strings.IndexRune(vowels, r) < 0 {
		return r
	}
	return -1
}

func shortHand(input string) string {
	// removes vowels from input string
	return strings.Map(filterVowel, input)
}

func checkError(e error){
	if e != nil {
		fmt.Println(e.Error())
	}
}

func getWeather(city string) string {
	apiURL := "https://api.openweathermap.org/data/2.5/weather"
	apiKey := "b784d5b15067cea930d3e00cdc0596e6"

	req, _ := http.NewRequest(http.MethodGet, apiURL, nil)

	queryParams := req.URL.Query()
	queryParams.Add("q", city)
	queryParams.Add("APPID", apiKey)

	req.URL.RawQuery = queryParams.Encode()

	client := http.Client{}
	response, err := client.Do(req)
	checkError(err)
	body, _ := ioutil.ReadAll(response.Body)

	var weatherData map[string]interface{}
	err = json.Unmarshal(body, &weatherData)
	checkError(err)

	weather := weatherData["weather"].([]interface{})

	return weather[0].(map[string]interface{})["main"].(string)
}

var UtiltyFunctionMap = template.FuncMap {
	"shorthand":shortHand,
	"weather": getWeather,
}

func init() {
	tpl = template.Must(template.New("").Funcs(UtiltyFunctionMap).ParseGlob("templates/*"))
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
	Delhi := Capital {
		"Delhi",
		"India",
		"INR",
	}

	Moscow := Capital {
		"Moscow",
		"Russia",
		"RUB",
	}

	Beijing := Capital {
		"Beijing",
		"China",
		"CNY",
	}

	Tokyo := Capital {
		"Tokyo",
		"Japan",
		"JPY",
	}

	Paris := Capital {
		"Paris",
		"France",
		"EUR",
	}

	capitals := []Capital{Delhi, Moscow, Beijing, Tokyo, Paris}

	err := tpl.ExecuteTemplate(os.Stdout, "functions.gohtml", capitals)
	handleError(err)
}
