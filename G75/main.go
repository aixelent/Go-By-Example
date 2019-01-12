package main

import (
	"io"
	"net/http"
)

func serverHandler(writer http.ResponseWriter, request *http.Request) {
	io.WriteString(writer, "HTTP server started.")
}

func main() {
	http.HandleFunc("/", serverHandler)
	http.ListenAndServe(":8001", nil)
}
