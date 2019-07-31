package main

import (
	"bufio"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"strings"
)

type entry struct {
	webSite string
	webName string
	quality string
}

var filename string

func dynamicContent(w http.ResponseWriter, r *http.Request) {
	var data []entry
	var f *os.File

	if filename == "" {
		f = os.Stdin
	} else {
		fileHandler, err := os.Open(filename)
		if err != nil {
			fmt.Printf("Error opening %s: %s", filename, err)
			os.Exit(1)
		}

		f = fileHandler
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	webTemplate := template.Must(template.ParseGlob("template.gohtml"))

	for scanner.Scan() {
		parts := strings.Fields(scanner.Text())
		if len(parts) == 3 {
			temp := entry{
				webSite: parts[0],
				webName: parts[1],
				quality: parts[2],
			}

			data = append(data, temp)
		}
	}

	fmt.Printf("Serving %s for %s \n", r.Host, r.URL.Path)
	webTemplate.ExecuteTemplate(w, "template.gohtml", data)
}

func staticPage(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Serving %s for %s \n", r.Host, r.URL.Path)
	webTemplate := template.Must(template.ParseGlob("static.gohtml"))
	webTemplate.ExecuteTemplate(w, "static.gohtml", nil)
}


func main() {
	arguments := os.Args

	if len(arguments) == 1 {
		filename = ""
	} else {
		filename = arguments[1]
	}

	http.HandleFunc("/static", staticPage)
	http.HandleFunc("/dynamic", dynamicContent)

	http.ListenAndServe(":8001", nil)
}