package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

const addr = "localhost:8000"

func Server(addr string) http.Server {
	albums := []Album{
		{"0", "Strength In Solitude LP", "Bvdub", "Electronic", "Techno", "2007", "Digital content is not available for sale or trade on Discogs."},
		{"1", "We Were The Sun", "Bvdub", "Electronic", "Ambient", "2009", "Digital content is not available for sale or trade on Discogs."},
		{"2", "A Prayer To False Gods<", "Bvdub", "Electronic", "Ambient", "2009", "Digital content is not available for sale or trade on Discogs."},
		{"3", "The Art Of Dying Alone", "Bvdub", "Electronic", "Ambient", "2010", "Digital content is not available for sale or trade on Discogs."},
		{"4", "A Silent Reign", "Bvdub", "Electronic", "Ambient", "2010", "Digital content is not available for sale or trade on Discogs."},
		{"5", "Presents Deep Space Mix 21", "Bvdub", "Electronic", "Ambient", "2011", "Digital content is not available for sale or trade on Discogs."},
		{"6", "Resistance Is Beautiful", "Bvdub", "Electronic", "Techno, Ambient", "2011", "Digital content is not available for sale or trade on Discogs."},
	}
	mux := http.NewServeMux()
	mux.HandleFunc("/albums", func(w http.ResponseWriter, r *http.Request) {
		encoder := json.NewEncoder(w)
		if r.Method == http.MethodGet {
			encoder.Encode(albums)
		} else if r.Method == http.MethodPost {
			data, err := ioutil.ReadAll(r.Body)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			r.Body.Close()
			album := Album{}
			json.Unmarshal(data, &album)
			album.ID = strconv.Itoa(len(albums) + 1)
			albums = append(albums, album)
			encoder.Encode(album)
		}
	})
	return http.Server{
		Addr:    addr,
		Handler: mux,
	}
}

type Album struct {
	ID          string
	Title       string `json:"title"`
	Author      string `json:"author"`
	Genre       string `json:"genre"`
	Style       string `json:"style"`
	Year        string `json:"year"`
	Description string `json:"description"`
}

func (a Album) toJson() string {
	return fmt.Sprintf(`{"Title":"%s", "Author":"%s", "Genre":"%s", "Style":"%s", "Year":"%s", "Description":"%s"}`, a.Title, a.Author, a.Genre, a.Style, a.Year, a.Description)
}

func AddAlbum(album Album) (Album, error) {
	r, err := http.Post("http://"+addr+"/albums", "application/json", strings.NewReader(album.toJson()))
	if err != nil {
		return Album{}, err
	}
	defer r.Body.Close()
	return DecodeAlbum(r.Body)
}

func DecodeAlbum(r io.Reader) (Album, error) {
	album := Album{}
	decoder := json.NewDecoder(r)
	err := decoder.Decode(&album)
	return album, err
}

func DecodeAlbums(r io.Reader) ([]Album, error) {
	albums := []Album{}
	decoder := json.NewDecoder(r)
	err := decoder.Decode(&albums)
	return albums, err
}

func GetAlbums() ([]Album, error) {
	r, err := http.Get("http://" + addr + "/albums")
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()
	return DecodeAlbums(r.Body)
}

func main() {
	s := Server(addr)
	go s.ListenAndServe()

	albums, err := GetAlbums()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Retrieved albums: %v\n", albums)

	album, err := AddAlbum(Album{"", "One Last Look At The Sea", "Bdvub", "Electronic", "Techno, Minimal, Ambient", "2011", ""})
	if err != nil {
		panic(err)
	}
	fmt.Printf("Added album: %v\n", album)
}
