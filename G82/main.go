package main

import (
	"fmt"
	"net/http"
	"time"
)

var Name = "asdfqwerty"

func main() {
	http.HandleFunc("/set-cookie", func(writer http.ResponseWriter, request *http.Request) {
		cookie := &http.Cookie{
			Name:    Name,
			Value:   "I'm a cookie!",
			Domain:  "localhost",
			Expires: time.Now().Add(time.Hour * 2),
		}
		http.SetCookie(writer, cookie)
		fmt.Fprintln(writer, "Cookie has been set.")
	})

	http.HandleFunc("/get-cookie", func(writer http.ResponseWriter, request *http.Request) {
		cookie, err := request.Cookie(Name)
		if err != nil {
			fmt.Fprintln(writer, "Cookie err: No cookies here, sorry!\n"+err.Error())
			return
		}
		fmt.Fprintf(writer, "Cookie %s \n", cookie.Value)
		fmt.Fprintf(writer, "other cookies")
		for _, v := range request.Cookies() {
			fmt.Fprintf(writer, "%s, %s \n", v.Name, v.Value)
		}
	})

	http.HandleFunc("/remove-cookie", func(writer http.ResponseWriter, request *http.Request) {
		cookie, err := request.Cookie(Name)
		if err != nil {
			fmt.Fprintln(writer, "Cookie error:"+err.Error())
			return
		}
		cookie.MaxAge = -1
		http.SetCookie(writer, cookie)
		http.SetCookie(writer, cookie)
		fmt.Fprintln(writer, "Cookie was removed.")
	})

	if err := http.ListenAndServe(":8001", nil); err != nil {
		panic(err)
	}
}
