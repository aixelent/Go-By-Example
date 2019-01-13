package main

import "net/http"

func main() {
	fileServ := http.FileServer(http.Dir("html"))
	fileServ = http.StripPrefix("/html", fileServ)

	http.HandleFunc("/file", serveFile)
	http.Handle("/html/", fileServ)
	http.ListenAndServe(":8001", nil)
}

func serveFile(writer http.ResponseWriter, request *http.Request) {
	http.ServeFile(writer, request, "file.txt")
}
