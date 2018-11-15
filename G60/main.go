package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

type Albums struct {
	XMLName xml.Name `xml:"albums"`
	Albums  []Album  `xml:"album"`
}

type Album struct {
	Title       string `xml:"title"`
	Author      string `xml:"author"`
	Genre       string `xml:"genre"`
	Style       string `xml:"style"`
	Year        string `xml:"year"`
	Description string `xml:"description"`
}

func readFle() {
	// f, err := os.OpenFile("file.xml", os.O_WRONLY, 0755)
	f, err := os.Open("file.xml")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	value, _ := ioutil.ReadAll(f)
	var albums Albums
	xml.Unmarshal(value, &albums)

	for i := 0; i < len(albums.Albums); i++ {
		fmt.Println("Album title", albums.Albums[i].Title)
		fmt.Println("Author:", albums.Albums[i].Author)
		fmt.Println("Genre:", albums.Albums[i].Genre)
		fmt.Println("Style:", albums.Albums[i].Style)
		fmt.Println("Year:", albums.Albums[i].Year)
		fmt.Println("Description:", albums.Albums[i].Description)
		fmt.Println()
	}
}

func main() {
	readFle()
}
