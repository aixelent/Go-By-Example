package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

const addr = "http://localhost:8080"

type StringServer string

func (s StringServer) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	request.ParseForm()
	fmt.Printf("Data: %v\n", request.Form)
	writer.Write([]byte(string(s)))
}

func goServer(addr string) http.Server {
	return http.Server{
		Addr:    addr,
		Handler: StringServer("Test"),
	}
}

func simplePost() {
	res, err := http.Post(addr, "application//x-www-form-urlencoded", strings.NewReader("name=Adrian"))
	if err != nil {
		panic(err)
	}

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	res.Body.Close()
	fmt.Println("Server response:" + string(data))
}

func useRequest() {}

func main() {
	s := goServer(addr)
	go s.ListenAndServe()

	simplePost()
	useRequest()
}
