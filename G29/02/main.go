package main

import (
	"fmt"
	"time"
)

func main() {
	loc, _ := time.LoadLocation("UTC")
	now := time.Now().In(loc)
	fmt.Println("\nToday: ", loc, "Time: ", now)

	ti1 := time.Date(2016, time.August, 11, 11, 22, 20, 100, loc)
	ti2 := time.Date(2018, time.August, 11, 11, 22, 20, 100, loc)

	duration := ti2.Sub(ti1)
	fmt.Println("Duration between two dates is:", duration)
}
