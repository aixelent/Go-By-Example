package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/album", func(writer http.ResponseWriter, request *http.Request) {
		if request.Method == http.MethodGet {
			fmt.Fprintln(writer, "GET Album")
		}
		if request.Method == http.MethodPost {
			fmt.Fprintln(writer, "POST Album")
		}
	})

	elemMux := http.NewServeMux()
	elemMux.HandleFunc("/artists/bvdub", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintln(writer, "BVDUB")
	})
	mux.Handle("/artists/", elemMux)

	http.ListenAndServe(":8001", mux)
}
