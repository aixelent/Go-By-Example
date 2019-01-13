package main

import (
	"html/template"
	"net/http"
)

func main() {
	tpl, err := template.ParseFiles("G76/tpl/page.tpl")
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		err := tpl.Execute(writer, "go template")
		if err != nil {
			panic(err)
		}
	})
	http.ListenAndServe(":8001", nil)
}
