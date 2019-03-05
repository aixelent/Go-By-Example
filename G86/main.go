package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func xkcd() {
	client := &http.Client{}

	req, _ := http.NewRequest("GET", "https://xkcd.com/info.0.json", nil)
	resp, err := client.Do(req)

	if err != nil {
		log.Fatalln("Something went wrong...")
	}
	defer resp.Body.Close()

	respBody, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(resp.Status, string(respBody))
}

func main() {
	xkcd()
}
