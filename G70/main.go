package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	r, err := http.Get("https://github.com")
	if err != nil {
		log.Fatalln(err)
	}
	defer r.Body.Close()

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatalln(err)
	}

	res := string(b)
	fmt.Println(res)
}
