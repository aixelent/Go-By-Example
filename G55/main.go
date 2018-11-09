package main

import (
	"fmt"
	"time"
)

func main() {
	ti := time.Date(2018, 11, 29, 0, 0, 0, 0, time.Local)
	fmt.Printf("The %v is:\n", ti)

	monthDay := ti.Day()
	weekDay := ti.Weekday()
	month := ti.Month()

	fmt.Printf("%d of %v, %v.", monthDay, month, weekDay)
}
