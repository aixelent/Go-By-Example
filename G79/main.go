package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.Handle("/users/", http.RedirectHandler("/login", http.StatusTemporaryRedirect))
	http.HandleFunc("/redirect/handlefunc/", func(writer http.ResponseWriter, request *http.Request) {
		http.Redirect(writer, request, "/login", http.StatusTemporaryRedirect)
	})
	http.HandleFunc("/login", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "Please login.")
	})
	http.ListenAndServe(":8001", nil)
}
